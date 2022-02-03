from python import SolvingBase


class Solving(SolvingBase):
    @classmethod
    def move_santa(cls, command: str, homes_coords: set, last_coords: list):
        if command == '^':
            last_coords[0] += 1
        elif command == 'v':
            last_coords[0] -= 1
        elif command == '>':
            last_coords[1] += 1
        elif command == '<':
            last_coords[1] -= 1

        homes_coords.add(tuple(last_coords))

    def first_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            homes_coords = {(0, 0)}  # X, Y
            last_coords = [0, 0]

            for command in file.read():
                self.move_santa(command=command, homes_coords=homes_coords, last_coords=last_coords)

            return len(homes_coords)

    def second_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            real_santa = {
                'homes': {(0, 0)},
                'last_coords': [0, 0]
            }

            robo_santa = real_santa.copy()

            for index_command, command in enumerate(file.read()):
                self.move_santa(
                    command=command,
                    homes_coords=robo_santa['homes'],
                    last_coords=robo_santa['last_coords']
                ) if index_command % 2 else self.move_santa(
                    command=command,
                    homes_coords=real_santa['homes'],
                    last_coords=real_santa['last_coords']
                )

            return len(real_santa['homes'] | robo_santa['homes'])


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
