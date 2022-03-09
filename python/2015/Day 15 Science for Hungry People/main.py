import re
import math

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        self.best_cookie, self.best_cookie_for_food_lovers = 0, 0

    def first_problem(self, check_calories: bool = False) -> int:
        with open(self.test_case, 'r', encoding="utf-8") as file:
            cookies_ingredient = [list(map(int, re.findall(r'-?\d+', line))) for line in file.readlines()]

        # I tried to use permutations, but with its use, the execution time becomes inappropriately long
        for i in range(101):
            for j in range(101 - i):
                for k in range(101 - i - j):
                    m = 100 - i - j - k
                    cookie_score = math.prod(max(0, sum(y * z for y, z in zip(x, (i, j, k, m))))
                                             for x in list(zip(*cookies_ingredient))[:-1])

                    self.best_cookie = max(self.best_cookie, cookie_score)
                    if sum(x[-1] * y for x, y in zip(cookies_ingredient, (i, j, k, m))) == 500:
                        self.best_cookie_for_food_lovers = max(self.best_cookie_for_food_lovers, cookie_score)

        return self.best_cookie

    def second_problem(self):
        return self.best_cookie_for_food_lovers


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
