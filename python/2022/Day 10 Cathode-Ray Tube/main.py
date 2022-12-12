from dataclasses import dataclass

from python import SolvingBase


@dataclass
class Instruction:
    cost: int
    value: int


@dataclass(init=False)
class CPU:
    register: int
    tact: int

    def __init__(self, register: int):
        self.register = register
        self.tact = 0

    def tick(self):
        self.tact += 1

    def update_register(self, value: int):
        self.register += value


@dataclass(init=False)
class CRT:
    width: int
    height: int
    pixels: list[list[str]]
    cpu: CPU

    def __init__(self, width: int, height: int, cpu: CPU):
        self.width = width
        self.height = height
        self.pixels = [["."] * self.width for _ in range(self.height)]
        self.cpu = cpu

    def draw(self):
        if self.cpu.tact % self.width in (self.cpu.register - 1, self.cpu.register, self.cpu.register + 1):
            self.lit_pixel(row=self.cpu.tact // self.width, column=self.cpu.tact % self.width)

    def lit_pixel(self, row: int, column: int):
        self.pixels[row][column] = "#"


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.instructions: list[Instruction] = [
                Instruction(cost=(2 if line[:4] == "addx" else 1), value=int(line[4:] or 0))
                for line in input_file.read().split("\n")
            ]

    @classmethod
    def get_signal_strength(cls, cpu: CPU) -> int:
        return cpu.tact * cpu.register if cpu.tact % 40 == 20 else 0

    def first_problem(self):
        cpu: CPU = CPU(register=1)
        signal_strength: int = 0
        for instruction in self.instructions:
            for _ in range(instruction.cost):
                cpu.tick()
                signal_strength += self.get_signal_strength(cpu=cpu)

            cpu.update_register(value=instruction.value)

        return signal_strength

    def second_problem(self):
        crt: CRT = CRT(width=40, height=6, cpu=CPU(register=1))

        for instruction in self.instructions:
            for tact in range(instruction.cost):
                crt.cpu.tick()
                if tact + 1 == instruction.cost:
                    crt.cpu.update_register(instruction.value)
                crt.draw()

        return crt.pixels  # To parse the letters, you need to print the list line by line


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
