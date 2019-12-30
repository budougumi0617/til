from typing import TypedDict, NoReturn


class UserDict(TypedDict):
    """Typed User definition."""

    user_name: str
    email: str


def show_user(user_dict: UserDict) -> NoReturn:
    print(user_dict)


user = UserDict(
    user_name=20,
)  # Parameter 'email' unfilled

user2 = UserDict(
    user_name=20,  # Expected type 'str', got 'int' instead
    email='user@exmaple.com',
)

show_user('user')  # Expected type 'UserDict', got 'str' instead
show_user(user2)
