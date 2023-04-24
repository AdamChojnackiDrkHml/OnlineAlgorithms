import matplotlib.pyplot as plt
import numpy as np
import csv
from matplotlib.ticker import MaxNLocator

Ns = [20, 30, 40, 50, 60, 70, 80, 90, 100]
Distros = ["Uni", "Geo", "Hrm", "Dhr"]
Algs = ["FIFO", "FWF", "LRU", "LFU", "RM", "RAND"]

res = [[0 for x in Distros] for i in Ns] 

print(res)

for i in range(len(Distros)):
    for j in range(len(Ns)):
        filename = "data/labs/" + (str(Ns[j]) + "_" + str(Distros[i]) + "2.txt")
        res[j][i] = np.loadtxt(filename, unpack='False')
        

# print(data)

for j in range(len(Ns)):
    for i in range(len(Distros)):
        # myK = list(map(lambda x: Ns[j] * x, KsRatios))
        ret = [int(i) for i in (res[j][i][0])]
        print(ret)
        for a in range(1, len(Algs)):
            plt.plot(ret, res[j][i][a], label=Algs[a])
        
        plt.xlabel("Cache Size")
        plt.ylabel("Avg Cost")
        plt.xticks(ret)
        plt.title("N = " + str(Ns[j]) + " " + Distros[i])
        plt.legend(loc='upper right')
        plt.savefig("data/labs/graphs/NDist/" + str(Ns[j]) + Distros[i] , bbox_inches="tight")
        plt.close()
    # plt.scatter(Ns, res[i], label="B = " + str(Ks[i]), s=2)
    # plt.legend(loc='upper right')
    # plt.savefig("data/plots/exp5bHLL" + str(Ks[i]), bbox_inches="tight")
    # plt.close()

for j in range(len(Ns)):
    ret = [int(i) for i in (res[j][0][0])]
    
    for a in range(1, len(Algs)):
    
        # myK = list(map(lambda x: Ns[j] * x, KsRatios))
        print(ret)
        for i in range(len(Distros)):
            plt.plot(ret, res[j][i][a], label=Distros[i])
        
        plt.xlabel("Cache Size")
        plt.ylabel("Avg Cost")
        plt.xticks(ret)
        plt.title("N = " + str(Ns[j]) + " " + Distros[i])
        plt.legend(loc='upper right')
        plt.savefig("data/labs/graphs/AlgDist/" + str(Ns[j]) + Algs[a] , bbox_inches="tight")
        plt.close()