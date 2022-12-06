from dataclasses import dataclass

from python import SolvingBase


@dataclass()
class Field:
    number: int
    marked: bool = False


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        with open(self.test_case, 'r', encoding='utf-8') as file:
            input_lines: list[str] = file.readlines()

            self.numbers: list[int] = self._read_numbers(input_string=input_lines[0])
            self.boards: list[list[Field]] = self._read_boards(input_strings=input_lines[2:])

    @classmethod
    def _read_numbers(cls, input_string: str) -> list[int]:
        return [int(num) for num in input_string.split(',')]

    @classmethod
    def _read_boards(cls, input_strings: list[str]) -> list[list[Field]]:
        split_boards = "".join(input_strings).replace('\n\n', '\t').replace('\n', ' ').replace('  ', ' ').split('\t')
        boards = [[Field(number=int(num)) for num in line.split(' ')] for line in split_boards]

        return boards

    def _separate_blocks(self, input_string: str) -> dict:
        ...

    @classmethod
    def _find_block_separator_index(cls, input_string: str) -> int:
        return input_string.find("\n\n")

    def first_problem(self) -> int:
        ...

    def second_problem(self) -> int:
        ...


if __name__ == "__main__":
    solve = Solving(test_case=True)
    solve.print_solutions()
