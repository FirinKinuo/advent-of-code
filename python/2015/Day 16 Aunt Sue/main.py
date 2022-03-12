import re

from python import SolvingBase


class Solving(SolvingBase):
    ATTRIBUTES_RE = r"(?P<attr>\w+): (?P<value>\d+)"

    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        with open('tape.txt', 'r', encoding="utf-8") as ticket_tape_file:
            self.ticket_tape = {attr[0]: attr[1] for attr in re.findall(self.ATTRIBUTES_RE, ticket_tape_file.read())}

    def first_problem(self) -> int:
        with open(self.test_case, 'r', encoding="utf-8") as file:
            for aunt in file.readlines():
                if all(self.ticket_tape[attr[0]] == attr[1] for attr in re.findall(self.ATTRIBUTES_RE, aunt)):
                    return int(re.search(r'(?P<id>\d+):', aunt).group('id'))

    def second_problem(self):
        with open(self.test_case, 'r', encoding="utf-8") as file:
            for aunt in file.readlines():
                if all(self.ticket_tape[attr[0]] < attr[1] if attr[0] in ('cats', 'trees') else
                       self.ticket_tape[attr[0]] > attr[1] if attr[0] in ('pomeranians', 'goldfish') else
                       self.ticket_tape[attr[0]] == attr[1]
                       for attr in re.findall(self.ATTRIBUTES_RE, aunt)):
                    return int(re.search(r'(?P<id>\d+):', aunt).group('id'))


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
