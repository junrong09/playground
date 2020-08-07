##########################
# List (Mutable, ordered)
##########################
nums = [1, 2]
nums.insert(1, 10)
print(nums)

a = 10
del a

nums = [1, 2, 3]
del nums[: -2]
print(nums)  # [2, 3]

print(2 in nums)  # Check membership

for i, v in enumerate(nums):  # Loop with index and value
    print(i, v)

# More functions, look at docs

##########################
# Tuple (Immutable, ordered)
##########################
t2 = (1, 2, 3)
print(t2[1])
# t2[1] = 20 # Error, tuple immutable

t1 = 1, 2, 3  # Tuple packing
x, y, z = t1  # Tuple unpacking (Must have matching len)

t3 = (t1, t2, [1, 2])  # Can nest objects
t3[2][0] = 20  # List is mutable
print(t3)

tEmpty = ()  # Empty tuple
tOneElement = 1,  # Single element, Cannot use (1)
print(tOneElement)
tOneElement = (1,)
print(tOneElement)

##########################
# Sets
##########################
setObj = {1, 2, 3, 3}  # duplicates will be removed
a = set('abc')  # Constructor takes in a iterable
print(a)  # {'a', 'b', 'c'}

emptySet = set()  # Cannot use {}, this creates empty dict
print(setObj)

print(1 in setObj)  # Test membership, True

# Set operator
a = set('abcd')
b = set('cde')

print(a - b)  # Minus
print(a | b)  # Union
print(a & b)  # Intercept
print(a ^ b)  # Sym diff (letters in a or b but not both)

##########################
# Dictionaries (Key, value pairs)
##########################
# keys must be immutable (e.g. strings, numbers, tuples with immutable elements)
nums = {'one': 1, 'two': 2, 4: 4}

del nums[4]

print("two" in nums)  # Check membership

list(nums)  # Convert keys to list

for k, v in nums.items():  # Loop with key and value (instead of just key)
    print(k, v)
