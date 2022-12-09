import math

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.trees: tuple[tuple[int]] = tuple(map(lambda tree_line: tuple(map(int, tree_line[:-1])),
                                                      input_file.readlines()))

    @classmethod
    def _find_visible_trees_for_side(cls, current_tree: int, side: list[int]) -> int:
        tree_count: int = 0
        for tree in side:
            if tree >= current_tree:
                tree_count += 1
                break
            tree_count += 1

        return tree_count

    def _find_visible_trees_from_point(self, pos_x: int, pos_y: int) -> int:
        current_tree = self.trees[pos_y][pos_x]
        row, column = self.trees[pos_y], [row[pos_x] for row in self.trees]
        return math.prod(self._find_visible_trees_for_side(current_tree=current_tree, side=side)
                         for side in (row[:pos_x][::-1], row[pos_x + 1:], column[:pos_y][::-1], column[pos_y + 1:]))

    def first_problem(self):
        height, width = len(self.trees), len(self.trees[0])
        visible_trees = 0
        for x in range(width):
            for y in range(height):
                if y in (0, height - 1) or x in (0, width - 1):
                    visible_trees += 1
                    continue

                tree = self.trees[y][x]
                row, column = self.trees[y], [row[x] for row in self.trees]

                visible_trees += max(row[:x]) < tree or max(row[x + 1:]) < tree or max(column[:y]) < tree or max(
                    column[y + 1:]) < tree

        return visible_trees

    def second_problem(self):
        height, width = len(self.trees), len(self.trees[0])
        return max(self._find_visible_trees_from_point(pos_x=x, pos_y=y) for y in range(height) for x in range(width))


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
