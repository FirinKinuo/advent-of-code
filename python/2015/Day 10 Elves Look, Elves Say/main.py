from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        self.sequences = ''
        with open(self.test_case, 'r', encoding='utf-8') as file:
            self.sequences = file.read()

    @classmethod
    def look_and_say(cls, sequences: str) -> str:
        output = ''
        count = 1
        for i in range(1, len(sequences)):
            if sequences[i - 1] == sequences[i]:
                count += 1
            else:
                output += f'{count}{sequences[i - 1]}'
                count = 1

        return output + f'{count}{sequences[-1]}'

    def first_problem(self):
        for _ in range(40):
            self.sequences = self.look_and_say(sequences=self.sequences)
        return len(self.sequences)

    def second_problem(self):
        for _ in range(10):
            self.sequences = self.look_and_say(sequences=self.sequences)
        return len(self.sequences)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
