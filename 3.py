lst = list(input().split())

ans = set(lst[0])
for i in range(1, len(lst)):
    ans.intersection_update(set(lst[i]))

print(sorted(list(ans)))