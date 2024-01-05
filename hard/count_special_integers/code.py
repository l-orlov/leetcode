from math import factorial

def countSpecialNumbers(N):
    L = list(map(int, str(N + 1)))
    n = len(L)

    res = 0
    for i in range(n - 1):
        res += 9 * perm(9, i)
    
    s = set()
    for i, x in enumerate(L):
        for y in range(i == 0, x):
            if y not in s:
                res += perm(9 - i, n - i - 1)
        if x in s: break
        s.add(x)
    return res

def perm(n, k):
    return factorial(n) / factorial(n - k)

print(countSpecialNumbers(435))