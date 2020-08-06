##########################
# Functions
##########################
# Scope
c = 20


def add(a, b):
    # c = 10 # creates a new local variable
    return a + b + c


print(add(1, 2))  # 23


# Default args
def fn1(a, b=10, c=1):  # Default fields must come after non-default
    return a + b * c


print(fn1(1))  # 11
print(fn1(1, c=2))  # 21, keyword args


# Positional args as tuple
def fn21(a, *tup, mustBeKeyword):  # Args after *variableName must be called with keyword
    print("Args as tuple")
    print(a)
    # Rest of the positional args stored in *variableName as tuple
    for t in tup:
        print(t)
    print(mustBeKeyword)


fn21(1, 2, 3, mustBeKeyword=20)


# Args as tuple (positional)/dict (keyword)
def fn22(a, *tup, **dict):
    print("Args as tuple/dict")
    print(a)
    # Rest of the positional args stored in *variableName as tuple
    for t in tup:
        print(t)
    # Keyword args stored in ##variableName as dict
    for d in dict:
        print(d, " ", dict[d])


fn22(20, 30, 40, b=50, c=60)

# Focus fn args to positional/keyword
# Args before "/" must be positional-only (Support in PY > 3.8)
# Args after "*" must be keyword-only
# def f(pos1, pos2, /, pos_or_kwd, *, kwd1, kwd2):
#     pass


# Unpacking (Spread operator)
listNum = [1, 2]
print(listNum)  # [1, 2]
print(*listNum)  # 1 2 # Unpack list

dictNum = {"a": 1, "b": 2}
print(dictNum)
print(add(**dictNum))  # Error because it is passed in as keyword args "one=1, two=2"


# Lambda Expression
# Come back

# Annotation & Documentation
def add(a: int, b: int = 0) -> int:  # Annotation
    """ Sum 2 numbers.

    Other details, if any.
    :param a:
    :param b:
    :return:
    """
    return a + b


print(add.__doc__)
print(add.__annotations__)  # {'a': <class 'int'>, 'b': <class 'int'>, 'return': <class 'int'>}

# Coding Style
# Standard: PEP 8
# Classname: UpperCamelCase, function/methods: lowercase_with_underscores
