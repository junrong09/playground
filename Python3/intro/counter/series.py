#  A module named "series" in a package "counter"

def fib(n: int) -> list:
    result = []
    a, b = 0, 1
    while a < n:
        result.append(a)
        a, b = b, a + b
    return result
