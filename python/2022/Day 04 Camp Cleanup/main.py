from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.pairs: tuple[tuple[set[int]]] = tuple(map(self._parse_pairs, input_file.read().split('\n')))

    @classmethod
    def _get_section_set(cls, section_assignment: str) -> set[int]:
        section_range: tuple[int] = tuple(map(int, section_assignment.split('-')))
        return set(i for i in range(section_range[0], section_range[1] + 1))

    @classmethod
    def _parse_pairs(cls, pair_string: str) -> tuple[set[int]]:
        return tuple(map(cls._get_section_set, pair_string.split(',')))

    def first_problem(self):
        return len([pair for pair in self.pairs if set.issubset(*pair) or set.issubset(*pair[::-1])])

    def second_problem(self):
        return len([pair for pair in self.pairs if set.intersection(*pair)])


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
