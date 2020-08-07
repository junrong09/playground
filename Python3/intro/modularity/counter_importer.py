##########################
# Package
##########################
from counter.series import *
result = fib(10)  # fib declared in __init__.py
add(10)  # Error, despite being declared in __init__.py

from counter import *
result = fib(10)
add(10)