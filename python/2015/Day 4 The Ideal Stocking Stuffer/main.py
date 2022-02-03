from hashlib import md5
from itertools import count

from python import SolvingBase


class Solving(SolvingBase):
    @classmethod
    def find_stuffer(cls, secret: str, count_zeros: int):
        for number in count(1):
            if md5(f'{secret}{number}'.encode('utf-8')).hexdigest()[:count_zeros] == '0' * count_zeros:
                return number

    def first_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            secret = file.read()
        return self.find_stuffer(secret=secret, count_zeros=5)

    def second_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            secret = file.read()
        return self.find_stuffer(secret=secret, count_zeros=6)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
