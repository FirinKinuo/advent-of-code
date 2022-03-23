from math import inf

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        with open(file=self.test_case, mode='r', encoding='utf-8') as input_file:
            self.total_presents = int(input_file.read())

    def first_problem(self, max_visits: float | int = inf, presents_per_elf: int = 10) -> int:
        total_presents = self.total_presents // presents_per_elf
        houses = [0] * total_presents
        search_house = total_presents

        for elf in range(1, total_presents):
            visits = 0
            for house in range(elf, total_presents, elf):
                visits += 1
                houses[house] += elf
                if houses[house] >= total_presents and house < search_house:
                    search_house = house

                if visits >= max_visits != inf:
                    break

        return search_house

    def second_problem(self):
        return self.first_problem(max_visits=50, presents_per_elf=11)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
