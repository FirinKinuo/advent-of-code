from os import path


class SolvingBase:
    def __init__(self, test_case: bool = True):
        self.test_case: str = path.join("test.txt") if test_case else path.join("input.txt")

    def first_problem(self):
        """Solving first problem"""
        ...

    def second_problem(self):
        """Solving first problem"""
        ...
