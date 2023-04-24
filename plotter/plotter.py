import os
import matplotlib.pyplot as plt
import sys 
import numpy as np
from itertools import permutations as perm
from itertools import cycle


with open(sys.argv[1], 'r') as f:
    lines = f.read().splitlines()

# lines.pop() #remove last \n

# # problem = lines.pop(0)
# # num_of_algs = int(lines.pop(0))
# # algs = lines.pop(0).strip().split(" ")
# # num_of_dist = int(lines.pop(0))
# # dists = lines.pop(0).strip().split(" ")
# # size = int(lines.pop(0))
# # names = []
# # for i in algs:
# #     for j in dists:
# #         names.append(i+"_"+j)
        
# # num_of_plots = len(names)
# print(names)
# print(num_of_plots)

names = ["FIFO", "FWF",	"LRU", "LFU", "RM", "RAND"]



data = [[] for x in range (7)]
for row in lines:
    i = 0
    row = row.strip()
    nums = row.split(" ")
    for num in nums:
        data[i].append(float(num))
        i = i + 1

X = data[0]
# print(data)
data.pop(0)
cycol = cycle('bgrcmk')
# print(data)
minY = 1.0
maxY = 0
for i, d in enumerate(data):
    # print(X, d, i)
    minY = min(minY, float(min(d)))
    maxY = max(maxY, float(max(d)))
    plt.plot(X, d, color=next(cycol), label=names[i])
    
# print(maxY)
# print(minY)
plt.xlabel("Size")
plt.ylabel("Avg Cost")
# plt.yticks(np.arange(minY - 0.1, maxY + 0.1, 12))
# plt.title(problem + "_" + str(size))

plt.legend()

resName = os.path.dirname(sys.argv[1]) + '/graphs/' + os.path.basename(sys.argv[1]) + '.png'
# print(resName)
plt.savefig(resName)


