import matplotlib.pyplot as plt
import numpy as np
import csv

Ns = [20, 30, 40, 50, 60, 70, 80, 90, 100]
KsRatios = [1.0 / 10.0, 1.0 / 9.0, 1.0 / 8.0, 1.0 / 7.0, 1.0 / 6.0, 1.0 / 5.0]
Distros = ["Uni", "Geo", "Hrm", "Dhr"]
Algs = ["FIFO", "FWF", "LRU", "LFU", "RM", "RAND"]

res = [[0 for x in Distros] for i in Ns] 

print(res)

for i in range(len(Distros)):
    for j in range(len(Ns)):
        filename = "data/labs/" + (str(Ns[j]) + "_" + str(Distros[i]) + ".txt")
        res[j][i] = np.loadtxt(filename, unpack='False')
        

# print(data)

for j in range(len(Ns)):
    for i in range(len(Distros)):
        myK = list(map(lambda x: Ns[j] * x, KsRatios))
        for a in range(len(Algs)):
            plt.plot(myK, res[j][i][a], label=Algs[a])
        
        plt.xlabel("Cache Size")
        plt.ylabel("Avg Cost")
        plt.title("N = " + str(Ns[j]) + " " + Distros[i])
        plt.legend(loc='upper right')
        plt.savefig("data/labs/graphs/NDist/" + str(Ns[j]) + Distros[i] , bbox_inches="tight")
        plt.close()
    # plt.scatter(Ns, res[i], label="B = " + str(Ks[i]), s=2)
    # plt.legend(loc='upper right')
    # plt.savefig("data/plots/exp5bHLL" + str(Ks[i]), bbox_inches="tight")
    # plt.close()

# for a in range(len(Algs)):   
#     for i in range(len(Distros)):
#         for j in range(len(KsRatios)):
#             plt.plot(Ns, res[j][i][a], label=str(Ns[j]))
        
#         plt.xlabel("Cache Size")
#         plt.ylabel("Avg Cost")
#         plt.title(Algs[a] + " " + Distros[i])
#         plt.legend(loc='upper right')
#         plt.savefig("data/labs/graphs/AlgDist/" + Distros[i] , bbox_inches="tight")
#         plt.close()
#         print(a)