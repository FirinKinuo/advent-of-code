import re
import math

from python import SolvingBase


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)
        self.second = 2503

    def first_problem(self):
        with open(self.test_case, 'r', encoding="utf-8") as file:
            """
            Looks scary, right? I'll explain right now!
            
            S - Speed
            D - Duration
            R - Rest
            X - Second
            
            If second % (duration + rest) > duration:
                Calculate the distance that the deer flying before the start of the rest
                D + S * math.ceil(X / (D + R))
            Else:
                Calculate the distance that the deer flew
                S * (X % (D + R)) + S * D * math.floor(X / (D + R))
            """
            return max(map(lambda deer_: deer_['during'] * deer_['speed'] * math.ceil(
                self.second / (deer_['during'] + deer_['rest'])) if self.second % (
                    deer_['during'] + deer_['rest']) > deer_['during'] else deer_['speed'] * (
                    self.second % (deer_['during'] + deer_['rest'])) + deer_['speed'] * deer_['during'] * math.floor(
                self.second / (deer_['during'] + deer_['rest'])), [{k: int(v) for k, v in re.search(
                    pattern=r'(?P<speed>\d+).* (?P<during>\d+).* (?P<rest>\d+)',
                    string=deer).groupdict().items()} for deer in file.readlines()]))

    def second_problem(self):
        with open(self.test_case, 'r', encoding="utf-8") as file:
            deer_list = [{key: int(value) for key, value in re.search(
                pattern=r'(?P<speed>\d+).* (?P<during>\d+).* (?P<rest>\d+)',
                string=deer).groupdict().items()}.values() for deer in file.readlines()]

            points = [0 for _ in range(len(deer_list))]
            distance = points.copy()
            flying = [fly - 1 for _, fly, _ in deer_list]

            for _ in range(self.second):
                for deer, (speed, fly, rest) in enumerate(deer_list):
                    if flying[deer] >= 0:
                        distance[deer] += speed
                    elif flying[deer] <= -rest:
                        flying[deer] = fly

                    flying[deer] -= 1

                max_distance = max(distance)
                for deer in range(len(distance)):
                    if distance[deer] == max_distance:
                        points[deer] += 1

        return max(points)


if __name__ == "__main__":
    solve = Solving(test_case=False)

    solve.print_solutions()
