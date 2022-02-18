from os import path
from datetime import datetime


class SolvingBase:
    def __init__(self, test_case: bool = True):
        self.test_case: str = path.join("test.txt") if test_case else path.join("input.txt")

    def first_problem(self):
        """Solving first problem"""
        ...

    def second_problem(self):
        """Solving first problem"""
        ...

    def print_solutions(self):
        print("===Troubleshooting Begins===")
        start_first_solution = datetime.now()
        first_problem = {
            'result': self.first_problem(),
            'time': datetime.now() - start_first_solution
        }
        start_second_solution = datetime.now()
        second_problem = {
            'result': self.second_problem(),
            'time': datetime.now() - start_second_solution
        }
        print(f"""~~~First problem~~~
        Time: {first_problem['time']}
        Result: {first_problem['result']}""")

        print(f"""~~~Second problem~~~
        Time: {second_problem['time']}
        Result: {second_problem['result']}""")

        print("===Problems solved===")
