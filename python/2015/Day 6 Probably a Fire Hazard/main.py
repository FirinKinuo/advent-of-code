import re

from python import SolvingBase


class Solving(SolvingBase):
    def first_problem(self):
        light_map = [False for _ in range(1000 * 1000)]
        actions = {
            'on': lambda _: True,
            'off': lambda _: False,
            'toggle': lambda state: not state
        }
        with open(self.test_case, 'r', encoding='utf-8') as file:
            for command in file.readlines():
                instruction = re.search(r'(?P<action>off|on|toggle)\s'
                                        r'(?P<fromX>\d+),(?P<fromY>\d+)\sthrough\s'
                                        r'(?P<toX>\d+),(?P<toY>\d+)', command).groupdict()

                for light_x in range(int(instruction['fromY']), int(instruction['toY']) + 1):
                    for light_y in range(int(instruction['fromX']), int(instruction['toX']) + 1):
                        light = 1000 * light_x + light_y
                        light_map[light] = actions[instruction['action']](light_map[light])
            return len([light for light in light_map if light])

    def second_problem(self):
        light_map = [0 for _ in range(1000 * 1000)]
        actions = {
            'on': lambda brightness: brightness + 1,
            'off': lambda brightness: brightness - 1,
            'toggle': lambda brightness: brightness + 2
        }
        with open(self.test_case, 'r', encoding='utf-8') as file:
            for command in file.readlines():
                instruction = re.search(r'(?P<action>off|on|toggle)\s'
                                        r'(?P<fromX>\d+),(?P<fromY>\d+)\sthrough\s'
                                        r'(?P<toX>\d+),(?P<toY>\d+)', command).groupdict()

                for light_x in range(int(instruction['fromY']), int(instruction['toY']) + 1):
                    for light_y in range(int(instruction['fromX']), int(instruction['toX']) + 1):
                        light = 1000 * light_x + light_y
                        light_map[light] = max(0, actions[instruction['action']](light_map[light]))
            return sum(light_map)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
