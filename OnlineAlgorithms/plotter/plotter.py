import matplotlib.pyplot as plt
import sys 
import itertools
from itertools import permutations as perm
from itertools import cycle


with open(sys.argv[1], 'r') as f:
    lines = f.read().splitlines()

problem = lines.pop(0)
num_of_algs = int(lines.pop(0))
algs = lines.pop(0).split(" ")
num_of_dist = int(lines.pop(0))
dists = lines.pop(0).split(" ")

names = []
for i in algs:
    for j in dists:
        names.append(i+"_"+j)
        
num_of_plots = len(names)
data = [[] for x in range (num_of_plots + 1)]
for row in lines:
    i = 0
    nums = row.split(" ")
    for num in nums:
        data[i].append(int(num))
        i = i + 1

X = data[0]
data.pop(0)

cycol = cycle('bgrcmk')
for i, d in enumerate(data):
    print(X, d, i)
    plt.plot(X, d, color=next(cycol), label=names[i])
    
plt.xlabel("Size")
plt.ylabel("Avg Cost")

plt.title(problem)

plt.legend()
plt.show()