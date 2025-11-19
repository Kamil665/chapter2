function find_month(n)
    return (n % 12) + 1
end

print("n = 3 -> x = " .. find_month(3))

for i = 0, 15 do
    print(string.format("n = %2d -> x = %2d", i, find_month(i)))
end
