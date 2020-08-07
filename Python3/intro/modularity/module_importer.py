# Run each code block in isolation to test
# from modularity import module
# result = module.fib(4)
# print(result)

# from modularity import module as mod  # Renaming imported modules using "as" is optional
# result = mod.fib(10)
# print(result)

# Best practice!!
# from modularity.module import fib  # Import class/function
# result = fib(10)
# print(result)

# import modularity.module  # Unable to import class/function e.g. cannot "import modularity.module.fib"
# result = modularity.module.fib(4)  # Must ref full name
# print(result)

# from modularity.module import *
# result = fib(10)  # Error (when comment other import), need to declare in __init__.py of package for * import
# print(result)

# print(dir(module))  # Prints all the definition of a module (e.g. variables, functions)
