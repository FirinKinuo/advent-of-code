from python import SolvingBase


class Solving(SolvingBase):
    @classmethod
    def _get_sonar_response(cls, case: str) -> list[int]:
        with open(case, 'r', encoding='utf-8') as sonar_data:
            return [int(sonar_echo) for sonar_echo in sonar_data]

    def first_problem(self):
        sonar_response = self._get_sonar_response(case=self.test_case)
        return len(
            [sonar_echo for index, sonar_echo in enumerate(sonar_response) if sonar_echo > sonar_response[index - 1]])

    def second_problem(self):
        sonar_response = self._get_sonar_response(case=self.test_case)
        return len(
            [sonar_echo for index, sonar_echo in enumerate(sonar_response) if sonar_echo > sonar_response[index - 3]])


if __name__ == "__main__":
    solve = Solving(test_case=False)

    print(f"First Problem: {solve.first_problem()}\nSecond Problem: {solve.second_problem()}")
