lst1 = list(map(int, input().split()))
lst2 = list(map(int, input().split()))

for item in sorted(zip(lst2, lst1)):
    print(item[1], end=" ")