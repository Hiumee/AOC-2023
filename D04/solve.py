data = []

while True:
    new = input()
    if new == "":
        break
    data.append(new)
    
result = sum([
        2**(len(win.intersection(my))-1)
        for line in data
        if (
            (win := set([int(x) for x in line.split(":")[1].split("|")[0].strip().split(" ") if x != ""])) or True
         ) and (
            (my := set([int(x) for x in line.split(":")[1].split("|")[1].strip().split(" ") if x != ""])) or True
         ) and len(win.intersection(my)) > 0
    ])

print("Part 1:", result)

points = [
        len(win.intersection(my))
        for line in data
        if (
            (win := set([int(x) for x in line.split(":")[1].split("|")[0].strip().split(" ") if x != ""])) or True
         ) and (
            (my := set([int(x) for x in line.split(":")[1].split("|")[1].strip().split(" ") if x != ""])) or True
         )
    ]

for i in range(len(points)-1, -1, -1):
    points[i] = sum(points[i:i+1+points[i]])

result = sum(points) + len(data)
print("Part 2:", result)