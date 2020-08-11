def get_number():
    number = float(input("Enter a int number: "))
    return number


try:
    number = get_number()
    if number % 1 > 0:
        raise ValueError("Non-int", "number")
except (ValueError, TypeError) as err:
    print(f"Invalid number {err}")
    # p1, p2 = err.args  # Access exception params
else:  # Exec when no exception
    print(f"Valid number: {number}")
finally:  # Exec regardless
    print("End of program")
