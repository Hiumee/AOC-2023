import re

data = []

while True:
    new = input()
    if new == "":
        break
    data.append(new)

R = 12
G = 13
B = 14

RREGEX = re.compile(r"([0-9]+) red")
GREGEX = re.compile(r"([0-9]+) green")
BREGEX = re.compile(r"([0-9]+) blue")

result = sum(
    int(line.split(":")[0][5:])
    for line in data
    if max(int(x) for x in RREGEX.findall(line.split(":")[1])) <= R and max(int(x) for x in GREGEX.findall(line.split(":")[1])) <= G and max(int(x) for x in BREGEX.findall(line.split(":")[1])) <= B
)

print("Part 1:", result)

result = sum(
    max(int(x) for x in RREGEX.findall(line.split(":")[1])) * max(int(x) for x in GREGEX.findall(line.split(":")[1])) * max(int(x) for x in BREGEX.findall(line.split(":")[1]))
    for line in data
)

print("Part 2:", result)
