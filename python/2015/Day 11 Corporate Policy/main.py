import re

from python import SolvingBase


class Solving(SolvingBase):
    STRAIGHT = r"abc|bcd|cde|def|efg|fgh|ghi|hij|ijk|jkl|klm|lmn|mno|nop|opq|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz"
    PAIRS_CHARS = r"(.)\1.*(.)\2"
    BAD_CHARS = r"[iol]"

    @classmethod
    def increment(cls, password: str) -> str:
        end_char_position = len(password) - 1
        updated_char = password[end_char_position]
        pre_password = password[:end_char_position]

        updated_char = chr(ord(updated_char) + 1) if updated_char != 'z' else 'a'

        if updated_char == 'a':
            pre_password = cls.increment(pre_password)

        return pre_password + updated_char

    @classmethod
    def check_password(cls, password: str) -> bool:
        if not re.search(cls.STRAIGHT, password) or re.search(cls.BAD_CHARS, password):
            return False

        return pairs.groups()[0] != pairs.groups()[1] if (pairs := re.search(cls.PAIRS_CHARS, password)) else False

    @classmethod
    def update_password(cls, password):
        while True:
            try:
                password = cls.increment(password)
                if cls.check_password(password):
                    return password
            except IndexError:  # The password has the letter Z in all positions
                break

    def first_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            current_password = file.read()
            return self.update_password(current_password)

    def second_problem(self):
        current_password = self.first_problem()
        return self.update_password(current_password)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
