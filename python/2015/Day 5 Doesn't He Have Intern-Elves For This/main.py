from collections import Counter
from python import SolvingBase


class Solving(SolvingBase):
    def first_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            return len([word for word in file.readlines()
                        if all(map(lambda prohibited: prohibited not in word, ('ab', 'cd', 'pq', 'xy'))) and
                        sum((Counter(word)[vowel] for vowel in ('a', 'e', 'i', 'o', 'u'))) >= 3 and
                        [word for letter in range(1, len(word)) if word[letter] == word[letter - 1]]])

    def second_problem(self):
        with open(self.test_case, 'r', encoding='utf-8') as file:
            return len([word for word in file.readlines()
                        if [word for letter in range(1, len(word) - 1) if word[letter - 1] == word[letter + 1]] and
                        [word for letter in range(len(word) - 1) if word[letter: letter + 2] in word[letter + 2:]]])


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
