from contextvars import ContextVar
from aiohttp import web
from aiohttp.web import middleware
from aiohttp.web_app import Application
from aiohttp.web_request import Request
from aiohttp.web_response import StreamResponse

# https://docs.aiohttp.org/en/stable/web_advanced.html#contextvars-support
USER_ID_CONTEXT = ContextVar('VAR', default='default')


@middleware
async def user_id_middleware(request: Request, handler) -> StreamResponse:
    if request.headers is not None and request.headers.get('x-user-id', '') != '':
        USER_ID_CONTEXT.set(request.headers.get('x-user-id'))
    return await handler(request)


async def handler(request: Request) -> StreamResponse:
    user_id: str = USER_ID_CONTEXT.get()
    return web.Response(text='previous user_id:' + user_id + '\n')

app: Application = web.Application()
app.middlewares.append(user_id_middleware)
app.router.add_get('/', handler)

web.run_app(app, port=33000)
