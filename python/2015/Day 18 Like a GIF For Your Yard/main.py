from copy import deepcopy

from python import SolvingBase


class Solving(SolvingBase):
    STEPS = 100

    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        self.grid = []
        self.__init_new_light_grid()

    def __init_new_light_grid(self):
        with open(self.test_case, 'r', encoding="utf-8") as input_file:
            self.grid = [list(map(lambda state: state == "#", line.replace('\n', '')))
                         for line in input_file.readlines()]

    def first_problem(self, wrong_lights: bool = False) -> int:
        for _ in range(self.STEPS):
            tmp_grid = deepcopy(self.grid)
            grid_size = len(self.grid) - 1
            for line in range(grid_size + 1):
                for row in range(grid_size + 1):
                    current = self.grid[line][row]
                    active_count = sum([grid_line[(row or 1) - 1:row + 2].count(1)
                                        for grid_line in tmp_grid[(line or 1) - 1:line + 2]]) - current
                    if wrong_lights:
                        if (line, row) in ((0, 0), (0, grid_size), (grid_size, 0), (grid_size, grid_size)):
                            continue
                    self.grid[line][row] = current and active_count in (2, 3) or not current and active_count == 3

        return sum(list(map(sum, self.grid)))

    def second_problem(self):
        self.__init_new_light_grid()

        grid_size = len(self.grid) - 1
        for line in (0, grid_size):
            for row in (0, grid_size):
                self.grid[line][row] = True

        return self.first_problem(wrong_lights=True)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
