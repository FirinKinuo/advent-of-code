from string import ascii_letters

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.rucksacks: list[str] = input_file.read().split('\n')

    @classmethod
    def _split_rucksack(cls, rucksack: str) -> tuple[str, str]:
        split: int = len(rucksack) // 2
        return rucksack[:split], rucksack[split:]

    @classmethod
    def _find_appears(cls, rucksack: str) -> set[str]:
        compartments: tuple[str, str] = cls._split_rucksack(rucksack=rucksack)
        return set(compartments[0]) & set(compartments[1])

    @classmethod
    def _find_badge(cls, rucksacks: list[str]) -> set[str]:
        return set(rucksacks[0]) & set(rucksacks[1]) & set(rucksacks[2])

    def first_problem(self):
        return sum(ascii_letters.index(*self._find_appears(rucksack=rucksack)) + 1
                   for rucksack in self.rucksacks)

    def second_problem(self):
        return sum(ascii_letters.index(*self._find_badge(rucksacks=self.rucksacks[i - 3:i])) + 1
                   for i in range(3, len(self.rucksacks) + 1, 3))


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
