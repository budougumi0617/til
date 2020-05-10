import logging
from contextvars import ContextVar
from aiohttp import web
from aiohttp.abc import BaseRequest, AbstractAccessLogger
from aiohttp.web import middleware
from aiohttp.web_app import Application
from aiohttp.web_request import Request
from aiohttp.web_response import StreamResponse

# https://docs.aiohttp.org/en/stable/web_advanced.html#contextvars-support
USER_ID_CONTEXT = ContextVar('VAR', default='default')


class CustomAccessLogger(AbstractAccessLogger):
    def log(self,
            request: BaseRequest,
            response: StreamResponse,
            time: float) -> None:
        self.logger.info(f'access_log: {request.remote} '
                         f'"{request.method} {request.path} '
                         f'done in {time}s: {response.status}')


@middleware
async def user_id_middleware(request: Request, handler) -> StreamResponse:
    if request.headers is not None and request.headers.get('x-user-id', '') != '':
        USER_ID_CONTEXT.set(request.headers.get('x-user-id'))
    return await handler(request)


async def handler(request: Request) -> StreamResponse:
    user_id: str = USER_ID_CONTEXT.get()
    return web.Response(text='previous user_id:' + user_id + '\n')


access_logger = logging.getLogger('aiohttp.access')
ch = logging.StreamHandler()
access_logger.setLevel(logging.INFO)
access_logger.addHandler(ch)

app: Application = web.Application()
app.middlewares.append(user_id_middleware)
app.router.add_get('/', handler)

web.run_app(app, port=33000, access_log_class=CustomAccessLogger)
