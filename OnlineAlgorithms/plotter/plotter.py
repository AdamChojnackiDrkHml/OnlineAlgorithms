import matplotlib.pyplot as plt
import sys 
import itertools
from itertools import permutations as perm
from itertools import cycle


with open(sys.argv[1], 'r') as f:
    lines = f.read().splitlines()

lines.pop() #remove last \n

problem = lines.pop(0)
num_of_algs = int(lines.pop(0))
algs = lines.pop(0).strip().split(" ")
num_of_dist = int(lines.pop(0))
dists = lines.pop(0).strip().split(" ")
size = int(lines.pop(0))
names = []
for i in algs:
    for j in dists:
        names.append(i+"_"+j)
        
num_of_plots = len(names)
print(names)
print(num_of_plots)
data = [[] for x in range (num_of_plots+1)]
for row in lines:
    i = 0
    row = row.strip()
    nums = row.split(" ")
    for num in nums:
        data[i].append(int(num))
        i = i + 1

X = data[0]
print(data)
data.pop(0)

cycol = cycle('bgrcmk')
print(data)
for i, d in enumerate(data):
    print(X, d, i)
    plt.plot(X, d, color=next(cycol), label=names[i])
    
plt.xlabel("Size")
plt.ylabel("Avg Cost")

plt.title(problem + "_" + str(size))

plt.legend()
plt.show()