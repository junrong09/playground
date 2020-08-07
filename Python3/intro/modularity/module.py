##########################
# Modules
##########################
def fib(n: int) -> list:
    result = []
    a, b = 0, 1
    while a < n:
        result.append(a)
        a, b = b, a + b
    return result


# If run as main (not imported), run fib fn as default.
# If imported, __name__ will resolved "modularity.module"
if __name__ == "__main__":
    import sys

    fib(int(sys.argv[1]))
