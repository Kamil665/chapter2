print "Введите двузначное число: "
number = gets.chomp.to_i

unless number.between?(10, 99)
  puts "Ошибка: число должно быть двузначным!"
  exit
end

puts " "

#а)
tens = number / 10

#б)
units = number % 10

#в)
summa = tens + units

#г)
multi = tens * units

puts "Исходное число: #{number}"
puts "а) Число десятков: #{tens}"
puts "б) Число единиц: #{units}"
puts "в) Сумма цифр: #{summa}"
puts "г) Произведение цифр: #{multi}"
