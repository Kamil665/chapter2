num = 4568

n1 = int(num / 1000)
n2 = int((num % 1000) / 100)
n3 = int(((num % 1000) % 100) / 10)
n4 = int(((num % 1000) % 100) % 10)

print("Четырёхзначное число: ", num)
print("а) ", n4, n3, n2, n1)
print("б) ", n2, n1, n4, n3)
print("в) ", n1, n3, n2, n4)
print("г) 1) ", n3, n4, n1, n2)

x1 = num % 100
x2 = x1 * 100
x3 = int(num / 100)
x4 = x2 + x3

print("г) 2) ", x4)
