import sys

temp = [] 
res = dict()

with open(sys.argv[1], 'r') as f:
    Lines = f.readlines()
    count = 0
    d={}
    for line in Lines: 
        print line
        x = line.split("?")
        res_dct = {x[i]: x[i + 1] for i in range(0, len(x), 2)}        
        d.update(res_dct)


for key, val in d.items(): 
   if val not in temp: 
      temp.append(val) 
      res[key] = val 

for key, value in res.iteritems():
    c=key+"?"+value
    w = open("output.txt", "a")
    w.write(c)
    w.close()
