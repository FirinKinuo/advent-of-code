from collections import defaultdict
from pathlib import Path

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.dirs_sizes: defaultdict = self._get_dirs_sizes(input_file.read().split('\n'))

    @classmethod
    def _get_dirs_sizes(cls, input_lines: list[str]) -> defaultdict:
        dirs_sizes: defaultdict = defaultdict(int)
        current_dir: Path = Path("/")

        for line in input_lines:
            match line.split():
                case ("dir", _) | ("$", "ls"):
                    continue
                case "$", "cd", path:
                    current_dir = current_dir.joinpath(path).resolve()
                case size, _:
                    dirs_sizes[current_dir] += int(size)
                    for parent in current_dir.parents:
                        dirs_sizes[parent] += int(size)

        return dirs_sizes

    def first_problem(self):
        return sum([dir_ for dir_ in self.dirs_sizes.values() if dir_ < 100000])

    def second_problem(self):
        required: int = 30000000 - (70000000 - self.dirs_sizes[Path("/").resolve()])
        return min(dir_size for dir_size in self.dirs_sizes.values() if dir_size >= required)


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
