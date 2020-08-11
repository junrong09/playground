##########################
# Scope & Namespace
##########################
def scope_test():
    def do_local():
        spam = "local spam"  # Garbage collected

    def do_nonlocal():
        nonlocal spam  # Ref to variable with value "test spam"
        spam = "nonlocal spam"

    def do_global():
        global spam  # Not declared in global scope yet
        spam = "global spam"

    spam = "test spam"
    do_local()
    print("After local assignment:", spam)  # After local assignment: test spam
    do_nonlocal()
    print("After nonlocal assignment:", spam)  # After nonlocal assignment: nonlocal spam
    do_global()
    print("After global assignment:", spam)  # After global assignment: nonlocal spam


scope_test()
print("In global scope:", spam)  # In global scope: global spam


##########################
# Classes
##########################
class Person(object):
    type = "human"

    def __init__(self):
        pass

    def print_details(self):
        print("Type: " + self.type)

    # Private by naming convention, __variableName
    def __get_type(self):
        return {"type": self.type}


class Student(Person):
    """
    Student object
    """
    school = "ABC"  # Class attribute

    def __init__(self, name):
        self.__name = name  # instance attribute

    # Similar to Java's toString(), so that object instance can be printed
    def __repr__(self):
        return f"<Student name={self.__name}?"

    # For len(obj) to work
    def __len__(self):
        return 10

    # To implement membership validation, (For the "in" syntax to work, e.g. "james" in studentObj)
    def __contains__(self, name):
        return self.__name == name

    def get_details(self):
        # To access private attri, use _classname__attriName
        return {"name": self.__name, **(super()._Person__get_type())}

    def print_details(self):
        super().print_details()
        print("Name: " + self.__name)

    # Python getter and setter (fn name of getter and setter has to align)
    @property
    def name(self):
        return self.__name

    # Python setter
    @name.setter
    def name(self, name):
        self.__name = name

    # # Alternatively, for backwards compatibility, if the fn name has already been declared
    # def get_name(self):
    #     return self.__name
    #
    # def set_name(self, name):
    #     self.__name = name
    #
    # name = property(get_name, set_name)  # This forces the obj.name to call the following fns when get,set,del.

    @classmethod  # Can access class attri/methods, static method can't
    def print_school(cls):
        print(cls.school)

    @classmethod
    def set_school(cls, school):
        cls.school = school

    @staticmethod  # Static method, cannot pass self/cls
    def go_school():
        print("Go School")




std1 = Student("James")
std2 = Student("Anny")

std1.name = "Change NAME"
print(std1.name)

print(std1.print_details())
print(std1.get_details())

Student.go_school()
std1.go_school()
