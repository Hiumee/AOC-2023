import re

data = []

while True:
    new = input()
    if new == "":
        break
    data.append(new)

REGEX = re.compile(r"([0-9]+)")

result = sum(
        [
            int(number)
            for i, line in enumerate(data)
            for j, char in enumerate(line)
            if char in "1234567890" and (j == 0 or data[i][j-1] not in "1234567890") and (number := REGEX.match(line[j:]).group(0)) and  (
                (
                    ((start := max(j-1, 0)),
                    (end := min(j+len(number), len(line)-1))) or
                    True
                ) and
                (
                    (
                        i != 0 and
                        any(
                            [
                                data[i-1][x] not in "1234567890."
                                for x in range(start, end+1)
                            ]
                        )
                    ) or
                    any(
                        [
                            data[i][x] not in "1234567890."
                            for x in range(start, end+1)
                        ]
                    ) or
                    (
                        i != len(data)-1 and
                        any(
                            [
                                data[i+1][x] not in "1234567890."
                                for x in range(start, end+1)
                            ]
                        )
                    )
                )
            )
        ]
    )


print("Part 1:", result)

# Giving up on part 2