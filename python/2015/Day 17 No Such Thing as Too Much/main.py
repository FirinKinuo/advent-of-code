from itertools import combinations, chain, zip_longest

from python import SolvingBase


class Solving(SolvingBase):
    EGGNOG_LITRES = 150

    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, 'r', encoding="utf-8") as input_file:
            self.containers = list(map(int, input_file))

        self.combinations_chain = list(chain(*(combinations(self.containers, container_count)
                                               for container_count in range(1, len(self.containers) + 1))))

    def first_problem(self) -> int:
        return list(map(sum, self.combinations_chain)).count(self.EGGNOG_LITRES)

    def second_problem(self):
        containers_ways = list(len(combination)
                               for combination in self.combinations_chain
                               if sum(combination) == self.EGGNOG_LITRES)

        return containers_ways.count(min(containers_ways))


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
