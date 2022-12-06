import re
from copy import deepcopy

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.stream: str = input_file.read()

    def find_start_marker(self, unique_count: int) -> int:
        for i in range(len(self.stream)):
            marker = self.stream[i:i + unique_count]
            if len(marker) == len(set(marker)):
                return unique_count + i

    def first_problem(self):
        return self.find_start_marker(unique_count=4)

    def second_problem(self):
        return self.find_start_marker(unique_count=14)


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
