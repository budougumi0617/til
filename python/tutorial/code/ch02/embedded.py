import math


# 10進数の2進数表記を得る関数
def get_binary_number(n):
    result = ''
    while n != 0:
        n, r = divmod(n, 2)
        result = str(r) + result
    return result
