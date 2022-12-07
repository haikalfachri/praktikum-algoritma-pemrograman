x = int(input())
while x > 0:
    div = x % 10
    x = x // 10
    print(div, end=" ")