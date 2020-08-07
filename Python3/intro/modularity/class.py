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

    # Private by naming convention
    def __get_type(self):
        return {"type": self.type}


class Student(Person):
    """
    Student object
    """
    def __init__(self, name):
        self.name = name

    def get_details(self):
        # To access private attri, use _classname__attriName
        return {"name": self.name, **(super()._Person__get_type())}

    def print_details(self):
        super().print_details()
        print("Name: " + self.name)


std1 = Student("James")
std2 = Student("Anny")

print(std1.print_details())
print(std1.get_details())