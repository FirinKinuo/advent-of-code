from ast import literal_eval

from python import SolvingBase


class Solving(SolvingBase):
    def first_problem(self):
        with open(self.test_case, 'r', encoding="utf-8") as file:
            return sum(len(s[:-1]) - len(literal_eval(s[:-1])) for s in file.readlines())

    def second_problem(self):
        with open(self.test_case, 'r', encoding="utf-8") as file:
            return sum(len(s[:-1].translate(str.maketrans({"\\": "\\\\", '"': '\\"'}))) + 2 - len(s[:-1])
                       for s in file.readlines())


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
