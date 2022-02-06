import re

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        self.identifiers = {}
        self.gates = {}

        self.operators = {
            'AND': lambda op_left, op_right: self.get_gate(op_left) & self.get_gate(op_right),
            'OR': lambda op_left, op_right: self.get_gate(op_left) | self.get_gate(op_right),
            'RSHIFT': lambda op_left, op_right: self.get_gate(op_left) >> self.get_gate(op_right),
            'LSHIFT': lambda op_left, op_right: self.get_gate(op_left) << self.get_gate(op_right),
            'NOT': lambda operand, _: 0xFFFF & ~(self.get_gate(operand)),
            '': lambda operand, _: self.get_gate(operand)
        }
        self.instruction_regexps = [
            r'(?P<operand_left>.*)\s(?P<operator>AND|OR|LSHIFT|RSHIFT)\s(?P<operand_right>.*)\s->\s(?P<dest>.*)',
            r'^(?P<operator>NOT|) *(?P<operand_left>[a-z0-9]*)\s->\s(?P<dest>.*)(?P<operand_right>)'
        ]

    def get_gate(self, key) -> int:
        if key.isnumeric():
            return int(key)

        if key not in self.gates.keys():
            gate = self.operators[self.identifiers[key]['operator']](
                self.identifiers[key]['operand_left'],
                self.identifiers[key]['operand_right']) if key.isalpha() else int(key)
            self.gates.update({key: gate})
        return self.gates[key]

    def first_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            for instruction_line in file.readlines():
                for regex in self.instruction_regexps:
                    instruction = re.search(regex, instruction_line)
                    if instruction:
                        instruction = instruction.groupdict()
                        self.identifiers |= {instruction['dest']: {
                            'operator': instruction['operator'],
                            'operand_left': instruction['operand_left'],
                            'operand_right': instruction['operand_right']}
                        }
        return self.get_gate('a')

    def second_problem(self):
        gate_a = self.first_problem()
        self.gates.clear()
        self.gates |= {'b': gate_a}
        return self.get_gate('a')


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
