import re
from copy import deepcopy

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            input_sections: tuple[list[str]] = tuple(line.split('\n')
                                                     for line in input_file.read().split("\n\n"))
            self.starting_stacks: list[list[str]] = [list(item for item in column if item != "")
                                                     for column in zip(*[line.replace("    ", " ").split(" ")
                                                                         for line in input_sections[0][:-1]][::-1])]
            self.instructions: tuple[tuple[int]] = tuple(tuple(map(int, re.findall(r'\d+', line)))
                                                         for line in input_sections[1])

    def first_problem(self):
        result: list[list[str]] = deepcopy(self.starting_stacks)
        for instruction in self.instructions:
            count, from_stack, to_stack = instruction[0], instruction[1] - 1, instruction[2] - 1
            for count in range(count):
                result[to_stack].append(result[from_stack].pop(-1))

        return "".join(stack[-1] for stack in result).replace('[', '').replace(']', '')

    def second_problem(self):
        result: list[list[str]] = deepcopy(self.starting_stacks)
        for instruction in self.instructions:
            count, from_stack, to_stack = instruction[0], instruction[1] - 1, instruction[2] - 1

            result[to_stack].extend(result[from_stack][-count:])
            del result[from_stack][-count:]

        return "".join(stack[-1] for stack in result).replace('[', '').replace(']', '')


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
