lst = list(input().split())
letters = set()

for i in range(len(lst)):
    lst[i] = sorted(lst[i])
    for letter in lst[i]:
        letters.add(letter)

ans = []
while len(lst[0]) > 0:
    for letter in letters:
        flag = True

        for word in lst:
            if letter in word:
                word.remove(letter)
            else:
                flag = False
                break
    
        if flag:
            ans.append(letter)

print(*sorted(ans))
