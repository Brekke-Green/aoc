filename = "input.txt"
def calculate(filename):
    f = open(filename, "r")
    result = 0
    input_str = f.read()
    input_str = input_str.split("\n")
    for line in input_str:
        first_half, second_half = set(line[:len(line)//2]), set(line[len(line)//2:]) 
        for item in first_half:
            if item in second_half:
                if item.islower():
                    result += ord(item) - 96
                else:
                    result += ord(item) - 38
    

    return result

print(calculate(filename))
