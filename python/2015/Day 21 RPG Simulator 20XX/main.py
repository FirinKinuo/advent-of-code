import re
from typing import TypedDict

from python import SolvingBase


class Item(TypedDict):
    cost: int
    damage: int
    armor: int


class Solving(SolvingBase):
    PLAYER_HIT = 100

    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        self.weapons = [
            Item(cost=8, damage=4, armor=0),
            Item(cost=10, damage=5, armor=0),
            Item(cost=25, damage=6, armor=0),
            Item(cost=40, damage=7, armor=0),
            Item(cost=74, damage=8, armor=0)
        ]
        self.armors = [
            Item(cost=13, damage=0, armor=1),
            Item(cost=31, damage=0, armor=2),
            Item(cost=53, damage=0, armor=3),
            Item(cost=75, damage=0, armor=4),
            Item(cost=102, damage=0, armor=5)
        ]
        self.rings = [
            Item(cost=0, damage=0, armor=0),
            Item(cost=25, damage=1, armor=0),
            Item(cost=50, damage=2, armor=0),
            Item(cost=100, damage=3, armor=0),
            Item(cost=20, damage=0, armor=1),
            Item(cost=40, damage=0, armor=2),
            Item(cost=80, damage=0, armor=3)
        ]

        with open(file=self.test_case, mode='r', encoding='utf-8') as input_file:
            boss_stats = re.search(
                pattern=r'(?P<hit>\d+).*\s(?P<damage>\d+).*\s(?P<armor>\d+)',
                string=input_file.read().replace('\n', '')).groupdict()

            self.boss = {key: int(value) for key, value in boss_stats.items()}

    def process_battle(self, player_set: tuple) -> bool:
        return (self.boss['damage'] - sum(item['armor'] for item in player_set)) / self.PLAYER_HIT <= \
               (sum(item['damage'] for item in player_set) - self.boss['armor']) / self.boss['hit']

    def first_problem(self) -> int:
        return min([sum([ring['cost'], ring_['cost'], armor['cost'], weapon['cost']])
                    for ring_ in self.rings for ring in self.rings for armor in self.armors for weapon in self.weapons
                    if ring_ != ring and self.process_battle((armor, weapon, ring, ring_))])

    def second_problem(self) -> int:
        return max([sum([ring['cost'], ring_['cost'], armor['cost'], weapon['cost']])
                    for ring_ in self.rings for ring in self.rings for armor in self.armors for weapon in self.weapons
                    if ring_ != ring and not self.process_battle((armor, weapon, ring, ring_))])


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
