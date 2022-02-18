import ast
import typing

from python import SolvingBase


class Solving(SolvingBase):
    @classmethod
    def sum_storage(cls, storage: typing.Any, exclude: typing.Any = None) -> int:
        result = 0
        if isinstance(storage, int):
            return storage

        if isinstance(storage, dict):
            if exclude in storage.values():
                return 0
            result += sum(cls.sum_storage(value, exclude) for value in storage.values())

        if isinstance(storage, (list, tuple)):
            result += sum(cls.sum_storage(value, exclude) for value in storage)

        return result

    def first_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            storage = ast.literal_eval(file.read())
            return self.sum_storage(storage=storage)

    def second_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            storage = ast.literal_eval(file.read())
            return self.sum_storage(storage=storage, exclude='red')


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
