
def x(n):
    if n == 0: return 0
    elif n == 1: return 1
    return (8)*x(n-1) - 16 * x(n-2)

print(x(3))   