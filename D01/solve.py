import re

data = []

while True:
    new = input()
    if new == "":
        break
    data.append(new)

REGEX_P1 = re.compile(r"[0-9]") # Chall 1
REGEX = re.compile(r"(?=([0-9]|one|two|three|four|five|six|seven|eight|nine))") # Chall 2
MAPPING = {"one": 1, "two": 2, "three": 3, "four": 4, "five":5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

result = sum(10*MAPPING[REGEX_P1.findall(line)[0]] + MAPPING[REGEX_P1.findall(line)[-1]] for line in data) # Chall 1

print("Part 1:", result)

result = sum(10*MAPPING[REGEX.findall(line)[0]] + MAPPING[REGEX.findall(line)[-1]] for line in data) # Chall 1

print("Part 2:", result)