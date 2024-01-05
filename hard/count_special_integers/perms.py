from math import factorial

def perm(n, k):
    return factorial(n) / factorial(n - k)
  
print(perm(5, 3))
