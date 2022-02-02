import math

from python import SolvingBase


class Solving(SolvingBase):
    def first_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            return sum(map(lambda perimeters: sum([perimeter * 2 for perimeter in perimeters]) + sorted(perimeters)[0],
                           list(map(lambda dimensions: [
                               dimensions[0] * dimensions[1],
                               dimensions[1] * dimensions[2],
                               dimensions[2] * dimensions[0]
                           ], (list(map(int, box.split('x'))) for box in file.readlines())))))

    def second_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            return sum(map(lambda dimensions: (dimensions[0] * 2 + dimensions[1] * 2) + math.prod(dimensions),
                           (sorted(list(map(int, box.split('x')))) for box in file.readlines())))


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
