import re

from itertools import count

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        self.replacements = []
        self.molecules = ''

        with open(file=self.test_case, mode='r', encoding='utf-8') as input_file:
            for line in input_file:
                if replace := re.search(r'(?P<from>\w+) => (?P<to>\w+)', line):
                    self.replacements.append(replace.groups())
                elif molecules_string := re.search(r'^(?P<molecules>\w+)$', line):
                    self.molecules = molecules_string.group('molecules')

    def first_problem(self) -> int:
        replaced_molecules = set()
        for replace, find in self.replacements:
            for molecule_index, molecule in enumerate(self.molecules):
                if molecule == replace:
                    replaced_molecules.add(
                        self.molecules[:molecule_index] + find + self.molecules[molecule_index + 1:])
                    continue

                if self.molecules[molecule_index: molecule_index + 2] == replace:
                    replaced_molecules.add(
                        self.molecules[:molecule_index] + find + self.molecules[molecule_index + 2:]
                    )

        return len(replaced_molecules)

    def second_problem(self):
        for steps in count(0, 1):
            if self.molecules == 'e':
                return steps

            for replace, find in self.replacements:
                if find in self.molecules:
                    self.molecules = self.molecules.replace(find, replace, 1)
                    break


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
