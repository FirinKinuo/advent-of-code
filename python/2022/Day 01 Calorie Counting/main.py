from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.elfs: list[int] = [sum(map(int, calories.split('\n')))
                                    for calories in "".join(input_file.readlines())[:-1].split("\n\n")]

    def first_problem(self):
        return max(self.elfs)

    def second_problem(self):
        return sum(sorted(self.elfs, reverse=True)[:3])


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
   