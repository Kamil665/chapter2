program swap_digits
    implicit none
    integer :: number, tens, units, swapped_number
    
    write(*,'(A)',advance='no') 'Введите двузначное число: '
    read(*,*) number
    
    if (number < 10 .or. number > 99) then
        write(*,*) 'Ошибка: число должно быть двузначным'
        stop
    end if
    
    tens = number / 10
    units = mod(number, 10)
    
    swapped_number = units * 10 + tens
    
    write(*,*) 'Инверсия числа: ', swapped_number
    
end program swap_digits
