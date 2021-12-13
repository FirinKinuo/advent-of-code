from typing import Union

from python import SolvingBase


class Solving(SolvingBase):
    class DiveSubmarine:
        def __init__(self):
            self.depth = 0
            self.horizontal = 0

            self.move_commands = {
                'forward': lambda position: self._change_horizontal(self.horizontal + position),
                'up': lambda position: self._change_depth(self.depth - position),
                'down': lambda position: self._change_depth(self.depth + position),
            }

        def _change_depth(self, value: int):
            self.depth = value

        def _change_horizontal(self, value: int):
            self.horizontal = value

    class DiveSubmarineToTarget:
        def __init__(self):
            self.depth = 0
            self.horizontal = 0
            self.target = 0
            self.move_commands = {
                'forward': lambda position: self._change_position(position),
                'up': lambda position: self._change_target(self.target - position),
                'down': lambda position: self._change_target(self.target + position),
            }

        def _change_target(self, value: int):
            self.target = value

        def _change_position(self, value: int):
            self.horizontal += value
            self.depth += self.target * value

    def _drive_submarine(self, dive_submarine: Union[DiveSubmarine, DiveSubmarineToTarget]) -> int:
        with open(self.test_case, 'r', encoding='utf-8') as sonar_data:
            commands = map(str.split, sonar_data)
            for direction, position in commands:
                dive_submarine.move_commands[direction](position=int(position))
        return dive_submarine.depth * dive_submarine.horizontal

    def first_problem(self) -> int:
        return self._drive_submarine(dive_submarine=self.DiveSubmarine())

    def second_problem(self) -> int:
        return self._drive_submarine(dive_submarine=self.DiveSubmarineToTarget())


if __name__ == "__main__":
    solve = Solving(test_case=False)

    print(f"First Problem: {solve.first_problem()}\nSecond Problem: {solve.second_problem()}")
