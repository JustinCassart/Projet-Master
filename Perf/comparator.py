import os
import sys
import matplotlib.pyplot as plt

def data_extractor(filename):
    iterations = []
    execution_times = []
    allocations = []
    memories = []
    with open(filename) as file:
        for line in file:
            if line[:5] == "Bench":
                splitted = line.split()
                iterations.append(int(splitted[1]))
                execution_times.append(int(splitted[2]))
                allocations.append(int(splitted[6]))
                memories.append(int(splitted[4]))
    return sum(iterations) / len(iterations), \
            (sum(execution_times) / len(execution_times) / 1000), \
            sum(allocations) / len(allocations), \
            sum(memories) / len(memories)

def display(times):
    for key in times.keys():
        for key2 in times[key].keys():
            for key3 in times[key][key2].keys():
                print(key, end=" ")
                print(key2, end=" ")
                print(key3, end=" ")
                print(times[key][key2][key3], end="\n")

def show_pattern_equal_text(times):
    for alphabet_size in times.keys():
        fig, ax = plt.subplots()
        x = []
        m1 = []
        m2 = []
        keys = list(times[alphabet_size].keys())
        keys.sort()
        for key2 in keys:
            x.append(key2)
            v1, v2 = times[alphabet_size][key2][key2]
            m1.append(v1[1])
            m2.append(v2[1])
        ax.plot(x, m1, label="Méthode1")
        ax.plot(x, m2, label="Méthode2")
        ax.set_xlabel("Taille du mot")
        ax.set_ylabel("Temps en ms")
        title = f"Motif et texte identique (alphabet de taille {alphabet_size})"
        ax.set_title(title)
        ax.legend()
        plt.savefig("/home/justin/Pictures/Projet/" + title + f"-{x[-1]}.png")
        # plt.vlines(500, 0, 4000)
        plt.show()
        plt.close()

def show_len_text(times):
    for alphabet_size in times.keys():
        for pattern_size in times[alphabet_size].keys():
            fig, ax = plt.subplots()
            x = list(times[alphabet_size][pattern_size])
            x.sort()
            m1 = []
            m2 = []
            for text_size in x:
                v1, v2 = times[alphabet_size][pattern_size][text_size]
                m1.append(v1[1])
                m2.append(v2[1])
            ax.plot(x, m1, label="Méthode1")
            ax.plot(x, m2, label="Méthode2")
            ax.set_xlabel("Taille du texte")
            ax.set_ylabel("Temps en ms")
            title = f"Alphabet de taille {alphabet_size} et motif de taille {pattern_size}"
            ax.set_title(title)
            ax.legend()
            plt.savefig("/home/justin/Pictures/Projet/" + title + f"-{x[-1]}.png")
            plt.close()

def show_len_text_allocations(times):
    for alphabet_size in times.keys():
        for pattern_size in times[alphabet_size].keys():
            fig, ax = plt.subplots()
            x = list(times[alphabet_size][pattern_size].keys())
            x.sort()
            m1 = []
            m2 = []
            for text_size in x:
                x.append(text_size)
                v1, v2 = times[alphabet_size][pattern_size][text_size]
                m1.append(v1[2])
                m2.append(v2[2])
            ax.plot(x, m1, label="Méthode1")
            ax.plot(x, m2, label="Méthode2")
            ax.set_xlabel("Taille du texte")
            ax.set_ylabel("Nombre d'allocations mémoire")
            title = f"Allocation pour alphabet de taille {alphabet_size} et motif de taille {pattern_size}"
            ax.set_title(title)
            ax.legend()
            plt.savefig("/home/justin/Pictures/Projet/" + title + f"-{x[-1]}.png")
            plt.close()

def show_pattern_equal_text_allocation(times):
    for alphabet_size in times.keys():
        fig, ax = plt.subplots()
        x = []
        m1 = []
        m2 = []
        for key2 in times[alphabet_size].keys():
            x.append(key2)
            v1, v2 = times[alphabet_size][key2][key2]
            m1.append(v1[2])
            m2.append(v2[2])
        ax.plot(x, m1, label="Méthode1")
        ax.plot(x, m2, label="Méthode2")
        ax.set_xlabel("Taille du mot")
        ax.set_ylabel("Nombre d'allocations mémoire")
        title = f"Allocation Motif et texte identique (alphabet de taille {alphabet_size})"
        ax.set_title(title)
        ax.legend()
        plt.savefig("/home/justin/Pictures/Projet/" + title + f"-{x[-1]}.png")
        plt.close()

def text(times, position, title, xlabel, ylabel):
    for alphabet_size in times.keys():
        for pattern_size in times[alphabet_size].keys():
            fig, ax = plt.subplots()
            x = list(times[alphabet_size][pattern_size].keys())
            x.sort()
            m1 = []
            m2 = []
            for text_size in x:
                v1, v2 = times[alphabet_size][pattern_size][text_size]
                m1.append(v1[position])
                m2.append(v2[position])
            ax.plot(x, m1, label="Méthode1")
            ax.plot(x, m2, label="Méthode2")
            ax.set_xlabel(xlabel)
            ax.set_ylabel(ylabel)
            titl = title + f"alphabet de taille {alphabet_size} et motif de taille {pattern_size}"
            ax.set_title(titl)
            ax.legend()
            plt.savefig("/home/justin/Pictures/Projet/" + titl + f"-{x[-1]}.png")
            plt.close()

def equal(times, position, title, xlabel, ylabel):
    for alphabet_size in times.keys():
        fig, ax = plt.subplots()
        m1 = []
        m2 = []
        x = list(times[alphabet_size].keys())
        x.sort()
        for key2 in x:
            v1, v2 = times[alphabet_size][key2][key2]
            m1.append(v1[position])
            m2.append(v2[position])
        ax.plot(x, m1, label="Méthode1")
        ax.plot(x, m2, label="Méthode2")
        ax.set_xlabel(xlabel)
        ax.set_ylabel(ylabel)
        titl = f"Motif et texte identique (alphabet de taille {alphabet_size})"
        ax.set_title(titl)
        ax.legend()
        plt.savefig("/home/justin/Pictures/Projet/" + title + f"-{x[-1]}.png")
        # plt.show()
        plt.close()

def pattern_in_text():
    times, allocs, n_allocs = {}, {}, {}
    line = "Le professeur pouvait trouver cette situation fort simple, mais la pensée de me promener sous la masse des eaux ne laissa pas de me préoccuper. Et cependant, que les plaines et les montagnes de l’Islande fussent suspendues sur notre tête, ou les flots de l’Atlantique, cela différait peu, en somme, du moment que la charpente granitique était solide. Du reste, je m’habituai promptement à cette idée, car le couloir, tantôt droit, tantôt sinueux, capricieux dans ses pentes comme dans ses détours, mais toujours courant au sud-est, et toujours s’enfonçant davantage, nous conduisit rapidement à de grandes profondeurs."
    size = 32
    while size <= len(line):
        pattern = line[:size]
        stream = os.popen(f"go test ./tests -run=None -bench=Fonction -benchmem -args '{pattern}'")
        out = stream.readlines()
        f1 = out[3].split()
        f2 = out[4].split()
        times[len(pattern)] = (int(f1[2])*1E-6, int(f2[2])*1E-6)
        allocs[len(pattern)] = (int(f1[4]), int(f2[4]))
        n_allocs[len(pattern)] = (int(f1[6]), int(f2[6]))
        size += 32

    _, ax = plt.subplots(2)
    x = list(times.keys())
    x.sort()
    y1_times = [times[i][0] for i in x]
    y2_times = [times[i][1] for i in x]

    for i in range(2):
        ax[i].plot(x, y1_times, label="Méthode1")
        ax[i].plot(x, y2_times, label="Méthode2")
        ax[i].set_xlabel("Taille motif")
        ax[i].set_ylabel("Temps en ms")
        titl = f"Temps de recherche pour un motif existant dans le texte"
        ax[i].set_title(titl)
        ax[i].legend()
    for i in range(32, 600, 32):
        ax[1].vlines(i, 0, 100)
    plt.show()
    plt.close()

def pattern_notin_text():
    times, allocs, n_allocs = {}, {}, {}
    line = "Le professeur pouvait trouver cette situation fort simple, mais la pensée de me promener sous la masse des eaux ne laissa pas de me préoccuper. Et cependant, que les plaines et les montagnes de l’Islande fussent suspendues sur notre tête, ou les flots de l’Atlantique, cela différait peu, en somme, du moment que la charpente granitique était solide. Du reste, je m’habituai promptement à cette idée, car le couloir, tantôt droit, tantôt sinueux, capricieux dans ses pentes comme dans ses détours, mais toujours courant au sud-est, et toujours s’enfonçant davantage, nous conduisit rapidement à de grandes profondeurs."
    size = 32
    while size <= len(line):
        pattern = line[:size-2]+"zz"
        stream = os.popen(f"go test ./tests -run=None -bench=Fonction -benchmem -args '{pattern}'")
        out = stream.readlines()
        f1 = out[3].split()
        f2 = out[4].split()
        times[len(pattern)] = (int(f1[2])*1E-6, int(f2[2])*1E-6)
        allocs[len(pattern)] = (int(f1[4]), int(f2[4]))
        n_allocs[len(pattern)] = (int(f1[6]), int(f2[6]))
        size += 32

    _, ax = plt.subplots(2)
    x = list(times.keys())
    x.sort()
    y1_times = [times[i][0] for i in x]
    y2_times = [times[i][1] for i in x]

    for i in range(2):
        ax[i].plot(x, y1_times, label="Méthode1")
        ax[i].plot(x, y2_times, label="Méthode2")
        ax[i].set_xlabel("Taille motif")
        ax[i].set_ylabel("Temps en ms")
        titl = f"Temps de recherche pour un motif existant dans le texte"
        ax[i].set_title(titl)
        ax[i].legend()
    for i in range(32, 600, 32):
        ax[1].vlines(i, 0, 100)
    plt.show()
    plt.close()

def pre_processing(fct: str, permutation: bool):
    times, allocs, n_allocs = {}, {}, {}
    line = "Le professeur pouvait trouver cette situation fort simple, mais la pensée de me promener sous la masse des eaux ne laissa pas de me préoccuper. Et cependant, que les plaines et les montagnes de l’Islande fussent suspendues sur notre tête, ou les flots de l’Atlantique, cela différait peu, en somme, du moment que la charpente granitique était solide. Du reste, je m’habituai promptement à cette idée, car le couloir, tantôt droit, tantôt sinueux, capricieux dans ses pentes comme dans ses détours, mais toujours courant au sud-est, et toujours s’enfonçant davantage, nous conduisit rapidement à de grandes profondeurs."
    size = 32
    maxytimes = 0
    maxbytes = 0
    maxallocs = 0
    while size <= len(line):
        pattern = line[:size]
        if permutation:
            pattern = pattern[:len(pattern)-2] + "zz"
        stream = os.popen(f"go test ./tests -run=None -bench={fct} -benchmem -args '{pattern}'")
        out = stream.readlines()
        f1 = out[3].split()
        f2 = out[4].split()
        times[len(pattern)] = (int(f1[2])*1E-6, int(f2[2])*1E-6)
        maxytimes = max(maxytimes, int(f1[2])*1E-6, int(f2[2])*1E-6)
        allocs[len(pattern)] = (int(f1[4]), int(f2[4]))
        maxbytes = max(maxbytes, int(f1[4]), int(f2[4]))
        n_allocs[len(pattern)] = (int(f1[6]), int(f2[6]))
        maxallocs = max(maxallocs, int(f1[6]), int(f2[6]))
        size += 32

    print(maxytimes, maxbytes, maxallocs)
    _, ax = plt.subplots(2)
    x = list(times.keys())
    x.sort()
    y1_times = [times[i][0] for i in x]
    y2_times = [times[i][1] for i in x]

    y1_bytes = [allocs[i][0] for i in x]
    y2_bytes = [allocs[i][1] for i in x]

    y1_allocs = [n_allocs[i][0] for i in x]
    y2_allocs = [n_allocs[i][1] for i in x]

    for i in range(2):
        ax[i].plot(x, y1_times, label="Méthode1")
        ax[i].plot(x, y2_times, label="Méthode2")
        ax[i].set_xlabel("Taille motif")
        ax[i].set_ylabel("Temps en ms")
        titl = f"Temps de recherche pour un motif existant dans le texte"
        ax[i].set_title(titl)
        ax[i].legend()
    for i in range(32, 600, 32):
        ax[1].vlines(i, 0, maxytimes)
    plt.show()
    plt.close()

    _, ax = plt.subplots(2)
    for i in range(2):
        ax[i].plot(x, y1_allocs, label="Méthode1")
        ax[i].plot(x, y2_allocs, label="Méthode2")
        ax[i].set_xlabel("Taille motif")
        ax[i].set_ylabel("Nombre d'allocations")
        titl = f"Nombre d'allocations effectuées pour un motif existant dans le texte"
        ax[i].set_title(titl)
        ax[i].legend()
    for i in range(32, 600, 32):
        ax[1].vlines(i, 0, maxallocs)
    plt.show()
    plt.close()

    _, ax = plt.subplots(2)
    for i in range(2):
        ax[i].plot(x, y1_bytes, label="Méthode1")
        ax[i].plot(x, y2_bytes, label="Méthode2")
        ax[i].set_xlabel("Taille motif")
        ax[i].set_ylabel("Nombre d'octets alloués")
        titl = f"Quantité d'octets alloués pour un motif existant dans le texte"
        ax[i].set_title(titl)
        ax[i].legend()
    for i in range(32, 600, 32):
        ax[1].vlines(i, 0, maxbytes)
    plt.show()
    plt.close()

if __name__ == "__main__":
    # foldername = sys.argv[1]
    # files = os.listdir(foldername)
    # if len(files) == 0:
    #     raise "no files"
    # d = {}

    # for file in files:
    #     if not os.path.isdir(file):
    #         values = data_extractor(foldername + "/" + file)
    #         splitted = file.split("_")
    #         k1 = int(splitted[1])
    #         k2 = int(splitted[2])
    #         k3 = int(splitted[3].split(".")[0])
    #         if not k1 in d.keys():
    #             d[k1] = {}
    #             d[k1][k2] = {}
    #             d[k1][k2][k3] = []
    #         elif not k2 in d[k1].keys():
    #             d[k1][k2] = {}
    #             d[k1][k2][k3] = []
    #         elif not k3 in d[k1][k2].keys():
    #             d[k1][k2][k3] = []
    #         d[k1][k2][k3].append(values)
    
    # # display(d)
    # # show_pattern_equal_text(d)
    # # show_len_text(d)
    # # show_len_text_allocations(d)
    # # show_pattern_equal_text_allocation(d)
    # equal(d, 1, "temps", "Taille des mots", "Temps en ms")
    # equal(d, 2, "allocations", "Taille des mots", "Nombres d'allocations")
    # equal(d, 3, "mémoire", "Taille des mots", "Mémoire utilisée (B)")
    # times, allocs, n_allocs = {}, {}, {}
    # pattern = "a"*5
    # while len(pattern) <= 300:
    #     stream = os.popen(f"go test ./tests -run=None -bench=Fonction -benchmem -benchtime=100x -args {pattern}")
    #     out = stream.readlines()
    #     f1 = out[3].split()
    #     f2 = out[4].split()
    #     times[len(pattern)] = (int(f1[2])*1E-6, int(f2[2])*1E-6)
    #     allocs[len(pattern)] = (int(f1[4]), int(f2[4]))
    #     n_allocs[len(pattern)] = (int(f1[6]), int(f2[6]))
    #     pattern += "a"*50
    # print(times)
    # fig, ax = plt.subplots()
    # x = list(times.keys())
    # x.sort()
    # y1 = [times[i][0] for i in x]
    # y2 = [times[i][1] for i in x]

    # ax.plot(x, y1, label="Méthode1")
    # ax.plot(x, y2, label="Méthode2")
    # ax.set_xlabel("Taille motif")
    # ax.set_ylabel("Temps en ms")
    # titl = f"Temps de recherche pour un motif inexistant dans le texte"
    # ax.set_title(titl)
    # ax.legend()
    # # plt.savefig("/home/justin/Pictures/Projet/" + title + f"-{x[-1]}.png")
    # plt.show()
    # plt.close()
    # pattern_in_text()
    # pattern_notin_text()
    # pre_processing("PreShiftAndFunction", False)
    pre_processing("ShiftAndFunction", False)
    # pre_processing("ShiftAndFunction", True) 