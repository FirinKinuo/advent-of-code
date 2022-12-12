from dataclasses import dataclass

from python import SolvingBase


@dataclass
class Point:
    x: int
    y: int

    def move(self, direction: "Point"):
        self.x += direction.x
        self.y += direction.y

    def move_new_point(self, direction: "Point") -> "Point":
        new_point = Point(x=self.x, y=self.y)
        new_point.move(direction=direction)

        return new_point

    def horizontal_touching(self, other: "Point") -> bool:
        return (self.x == other.x and abs(self.y - other.y) == 2) or (self.y == other.y and abs(self.x - other.x) == 2)

    def distance(self, other: "Point") -> int:
        return abs(self.x - other.x) + abs(self.y - other.y)

    def __hash__(self) -> int:
        return hash((self.x, self.y))


class Solving(SolvingBase):
    def __init__(self, *args, **kwargs):
        super(Solving, self).__init__(*args, **kwargs)

        self.directions: dict[str, Point] = {
            "U": Point(x=0, y=-1),
            "D": Point(x=0, y=1),
            "L": Point(x=-1, y=0),
            "R": Point(x=1, y=0)
        }

        self.diagonal_directions: tuple[Point, Point, Point, Point] = (
            Point(x=1, y=1),
            Point(x=-1, y=1),
            Point(x=1, y=-1),
            Point(x=-1, y=-1)
        )

        with open(self.test_case, "r", encoding="utf-8") as input_file:
            self.motions: list = list(map(str.split, input_file.read().split("\n")))

    def first_problem(self):
        head: Point = Point(x=0, y=0)
        tail: Point = Point(x=0, y=0)
        seen: set = set()

        for move in self.motions:
            direction = self.directions[move[0]]
            for _ in range(int(move[1])):
                head.move(direction=direction)

                if head.horizontal_touching(tail):
                    tail = tail.move_new_point(direction=direction)
                elif head.distance(tail) == 3:
                    tail = min(map(tail.move_new_point, self.diagonal_directions),
                               key=lambda point: head.distance(point))

                seen.add(tail)

        return len(seen)

    def second_problem(self):
        head: Point = Point(x=0, y=0)
        tails: list[Point] = [Point(x=0, y=0)] * 9
        seen: set = set()

        for move in self.motions:
            direction = self.directions[move[0]]
            for _ in range(int(move[1])):
                head.move(direction=direction)

                for i in range(len(tails)):
                    next_tail = head if i == 0 else tails[i - 1]
                    tail = tails[i]
                    if next_tail.horizontal_touching(tail):
                        tail_direction = self.directions.values()
                    elif next_tail.distance(tail) >= 3:
                        tail_direction = self.diagonal_directions
                    else:
                        continue

                    tails[i] = min(map(tail.move_new_point, tail_direction),
                                   key=lambda point: next_tail.distance(point))

                seen.add(tails[-1])

        return len(seen)


if __name__ == "__main__":
    solve: Solving = Solving(test_case=False)

    solve.print_solutions()
