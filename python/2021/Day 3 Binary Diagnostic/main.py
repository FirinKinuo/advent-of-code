from collections import Counter

from python import SolvingBase


class Solving(SolvingBase):
    def first_problem(self) -> int:
        with open(self.test_case, 'r', encoding='utf-8') as diagnosis:
            bites = ''.join(Counter(el).most_common(1)[0][0] for el in zip(*diagnosis.readlines()[::-1]))
            gamma = int(bites, 2)
            epsilon = gamma ^ int('1' * len(bites), 2)
            return gamma * epsilon

    @classmethod
    def _compute_rating_gases(cls, diagnosis_data: list, fewer_bit: bool = False) -> int:
        bit = 0
        while len(diagnosis_data) > 1:
            bit_criteria = list(Counter(el).most_common(fewer_bit + 1) for el in zip(*diagnosis_data[::-1]))
            bit_criteria = bit_criteria[bit][-fewer_bit][0] if bit_criteria[bit][-fewer_bit][1] != len(
                diagnosis_data) / 2 else str(int(not fewer_bit))
            diagnosis_data = list(filter(lambda bit_line: bit_line[bit] == bit_criteria, diagnosis_data))
            bit += 1

        return int(diagnosis_data[0], 2)

    def second_problem(self) -> int:
        with open(self.test_case, 'r', encoding='utf-8') as diagnosis:
            diagnosis_lines = diagnosis.readlines()

            oxygen_rating = self._compute_rating_gases(diagnosis_data=diagnosis_lines.copy(), fewer_bit=False)
            co2_rating = self._compute_rating_gases(diagnosis_data=diagnosis_lines.copy(), fewer_bit=True)
            return oxygen_rating * co2_rating


if __name__ == "__main__":
    solve = Solving(test_case=False)
    solve.print_solutions()
