filename = "input.txt"
def calculate(filename):
    f = open(filename, "r")
    result = 0
    input_str = f.read()
    input_arr = input_str.split("\n")
    index = 0
    while index <= len(input_arr) - 3:
        #counter = dict()
        set1, set2, set3 = set(input_arr[index]), set(input_arr[index+1]), set(input_arr[index+2])

        lst = [set1, set2, set3]
        in_all = set.intersection(*lst)

        for i in in_all: 
            item = i

        if item.islower():
            result += ord(item) - 96
        else:
            result += ord(item) - 38

        index+=3
    

    return result

print(calculate(filename))
