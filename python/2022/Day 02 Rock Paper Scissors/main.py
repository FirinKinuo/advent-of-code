from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.strategy: list[tuple[int, int]] = [self._decrypt_strategy(strategy=strategy.replace('\n', ''))
                                                    for strategy in input_file.readlines()]

    @classmethod
    def _decrypt_strategy(cls, strategy: str) -> tuple[int, int]:
        opponent, player = strategy.split()
        return "ABC".index(opponent) + 1, "XYZ".index(player) + 1

    def first_problem(self):
        def calculate_round(current_round: tuple[int, int]) -> int:
            win, draw, lose = 6, 3, 0
            result: int = (current_round[1] - current_round[0]) % 3
            return current_round[1] + (win if result == 1 else draw if result == 0 else lose)

        return sum(map(calculate_round, self.strategy))

    def second_problem(self):
        def calculate_round(current_round: tuple[int, int]) -> int:
            result: tuple = (
                ((current_round[0] - 2) % 3) + 1,
                current_round[0] + 3,
                ((current_round[0]) % 3) + 6 + 1
            )
            return result[current_round[1] - 1]

        return sum(map(calculate_round, self.strategy))


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
