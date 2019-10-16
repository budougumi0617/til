import code.ch02.embedded as embedded
import pytest


@pytest.mark.parametrize("n, want", [
    (20, '10100'),
    (23, '10111'),
])
def test_get_binary_number(n, want):
    assert embedded.get_binary_number(n) == want
