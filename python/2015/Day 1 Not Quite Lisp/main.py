from python import SolvingBase


class Solving(SolvingBase):
    def first_problem(self):
        floor = 0

        with open(self.test_case, 'r', encoding='utf-8') as file:
            instructions = file.read()

        for command in instructions:
            floor += 1 if command == '(' else -1

        return floor

    def second_problem(self):
        floor = 0

        with open(self.test_case, 'r', encoding='utf-8') as file:
            instructions = file.read()

        for command_index, command in enumerate(instructions):
            floor += 1 if command == '(' else -1

            if floor == -1:
                return command_index + 1


if __name__ == "__main__":
    solve = Solving(test_case=False)

    print(f"First Problem: {solve.first_problem()}\nSecond Problem: {solve.second_problem()}")
