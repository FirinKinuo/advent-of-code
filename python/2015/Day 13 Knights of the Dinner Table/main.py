import re

from itertools import permutations

from python import SolvingBase


class Solving(SolvingBase):
    STATE = {
        'gain': lambda state: state,
        'lose': lambda state: -state
    }

    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        with open(self.test_case, 'r', encoding='utf-8') as file:
            self.guests_state = self.init_guests_state(guests_list=file.readlines())

    @classmethod
    def init_guests_state(cls, guests_list: list) -> dict:
        guests_state = {}
        for guest_state in guests_list:
            guest_state = re.search(
                r'(?P<name>\w+).*(?P<state>gain|lose)\s(?P<unit>\d+).*to\s(?P<neighbor>\w+)',
                guest_state).groupdict()
            if guest_state['name'] not in guests_state.keys():
                guests_state[guest_state['name']] = {}

            guests_state[guest_state['name']] |= {
                guest_state['neighbor']: cls.STATE[guest_state['state']](int(guest_state['unit']))}

        return guests_state

    @classmethod
    def _calculate_happiness(cls, order: list | tuple, guests_state: dict) -> int:
        happiness = 0
        left_guest = order[-1]
        for right_guest in order:
            happiness += guests_state[left_guest][right_guest] + guests_state[right_guest][left_guest]
            left_guest = right_guest
        return happiness

    def first_problem(self) -> int:
        max_happiness = 0
        for order in permutations(self.guests_state.keys()):
            happiness = self._calculate_happiness(order=order, guests_state=self.guests_state)
            max_happiness = max(max_happiness, happiness)

        return max_happiness

    def second_problem(self):
        self.guests_state['Me'] = {}
        for guest in self.guests_state.keys():
            self.guests_state[guest]['Me'] = 0
            self.guests_state['Me'][guest] = 0

        return self.first_problem()


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
