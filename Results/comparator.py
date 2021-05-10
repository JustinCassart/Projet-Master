import os

METHOD1 = "BenchmarkBigShiftAnd"
METHOD2 = "BenchmarkBigShiftAndMultiMask"

def equalcomparator():
    res1 = os.system(f"go test -run=NONE -bench={METHOD1}")
    # print(res1)

def data_extractor(filename):
    iterations = []
    execution_times = []
    allocations = []
    with open(filename) as file:
        for line in file:
            if line[:5] == "Bench":
                splitted = line.split()
                iterations.append(int(splitted[1]))
                execution_times.append(int(splitted[2]))
                allocations.append(int(splitted[6]))
    return iterations, execution_times, allocations

if __name__ == "__main__":
    files = os.listdir(".")[:-1]
    times = {}

    for file in files:
        values = data_extractor(file)
        splitted = file.split("_")
        k1 = int(splitted[1])
        k2 = int(splitted[2])
        k3 = int(splitted[3].split(".")[0])
        if not k1 in d.keys():
            d[k1] = {}
            d[k1][k2] = {}
            d[k1][k2][k3] = []
        elif not k2 in d[k1].keys():
            d[k1][k2] = {}
            d[k1][k2][k3] = []
        elif not k3 in d[k1][k2].keys():
            d[k1][k2][k3] = []
        d[k1][k2][k3].append(values)
    
    