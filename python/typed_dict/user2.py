from datetime import datetime
from typing import TypedDict, NoReturn


class UserDict(TypedDict):
    """Typed User definition."""

    user_name: str
    email: str
    created: datetime
    updated: datetime


def show_user(user_dict: UserDict) -> NoReturn:
    print(user_dict)


user = UserDict(
    user_name='John Due',
    email='john@foo.com',
    created=datetime.now(),
    updated=datetime.now()
)

print(user['phone'])  # TypedDict "UserDict" has no key 'phone'


