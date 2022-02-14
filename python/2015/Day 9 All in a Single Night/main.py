import re

from itertools import permutations
from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        self.routes = {}

        with open(self.test_case, 'r', encoding="utf-8") as file:
            for distance in file.readlines():
                route = re.search(r'(?P<from>\w*)\sto\s(?P<to>\w*)\s=\s(?P<distance>\d*)', distance).groupdict()
                self.routes[route['from']] = self.routes.get(route['from']) or {}
                self.routes[route['to']] = self.routes.get(route['to']) or {}
                self.routes[route['from']] |= {route['to']: int(route['distance'])}
                self.routes[route['to']] |= {route['from']: int(route['distance'])}

    def find_route(self, method: callable, start_distance: int) -> int:
        cites = self.routes.keys()
        search_distance = start_distance
        for city_order in permutations(cites):
            distance = 0
            current_city = city_order[0]
            for city in city_order[1:]:
                try:
                    distance += self.routes[current_city].get(city)
                    current_city = city
                except TypeError:
                    pass

            search_distance = method(search_distance, distance)

        return search_distance

    def first_problem(self):
        return self.find_route(min, max([max(x.values()) for x in self.routes.values()]))

    def second_problem(self):
        return self.find_route(max, 0)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
