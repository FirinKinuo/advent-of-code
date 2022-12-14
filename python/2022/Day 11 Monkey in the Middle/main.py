from dataclasses import dataclass
from functools import total_ordering
from copy import deepcopy
import re
import math

from python import SolvingBase


@dataclass
@total_ordering
class Monkey:
    items: list[int]
    operator: callable
    operand: int | str
    test_value: int
    test_true: int
    test_false: int
    inspect_count: int = 0

    def test(self, item: int) -> int:
        return self.test_true if item % self.test_value == 0 else self.test_false

    def inspect_item(self, item: int) -> int:
        self.inspect_count += 1
        return self.operator(item, self.operand if isinstance(self.operand, int) else item)

    def __eq__(self, other: "Monkey") -> bool:
        return self.inspect_count == other.inspect_count

    def __lt__(self, other: "Monkey") -> bool:
        return self.inspect_count < other.inspect_count


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            pattern: re.Pattern = re.compile(r'.*items:\s(?P<items>.*)\\.*old\s(?P<operator>.*)\s'
                                             r'(?P<operand>\d*\w*)\\.*\s(?P<test>\d+).*\s'
                                             r'(?P<test_true>\d+).*\s(?P<test_false>\d+).*')
            self.monkeys: list[Monkey] = [self._parse_monkey(input_string=monkey.replace("\n", "\\n"), pattern=pattern)
                                          for monkey in input_file.read().split("\n\n")]

    @classmethod
    def _parse_monkey(cls, input_string: str, pattern: re.Pattern) -> Monkey:
        raw_monkey: dict = pattern.search(input_string).groupdict()
        monkey = Monkey(
            items=list(map(int, raw_monkey['items'].split(', '))),
            operator=int.__mul__ if raw_monkey['operator'] == "*" else int.__add__,
            operand=int(raw_monkey['operand']) if raw_monkey['operand'].isdigit() else raw_monkey['operand'],
            test_value=int(raw_monkey['test']),
            test_true=int(raw_monkey['test_true']),
            test_false=int(raw_monkey['test_false'])
        )

        return monkey

    def first_problem(self):
        monkeys = deepcopy(self.monkeys)
        for _ in range(20):
            for monkey in monkeys:
                while monkey.items:
                    new_item = monkey.inspect_item(item=monkey.items[-1])
                    new_item //= 3
                    monkeys[monkey.test(new_item)].items.append(new_item)
                    monkey.items.pop(-1)

        return math.prod(monkey.inspect_count for monkey in sorted(monkeys)[-2:])

    def second_problem(self):
        monkeys = deepcopy(self.monkeys)
        mod = math.prod(monkey.test_value for monkey in monkeys)
        for _ in range(10000):
            for monkey in monkeys:
                while monkey.items:
                    new_item = monkey.inspect_item(item=monkey.items[-1])
                    new_item %= mod
                    monkeys[monkey.test(new_item)].items.append(new_item)
                    monkey.items.pop(-1)

        return math.prod(monkey.inspect_count for monkey in sorted(monkeys)[-2:])


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
