num = int(279)

n1 = int(num / 100)

n2 = int((num % 100) / 10)

n3 = int((num % 100) % 10)

print("Трёхзначное число: ", num)
print("После перестановки: ", n2, n1, n3)
