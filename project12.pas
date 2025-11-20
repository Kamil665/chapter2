program ThreeDigitNumber;
var
   number, hundreds, tens, units: integer;
   summa, multi: integer;
begin
   write('Введите трёхзначное число: ');
   readln(number);
   
   hundreds := number div 100;
   tens := (number div 10) mod 10;
   units := (number mod 100) mod 10;
   
   summa := hundreds + tens + units;
   multi := hundreds * tens * units;
   
   writeln();
   writeln('Число единиц в нём: ', units);
   writeln('Число десятков в нём: ', tens);
   writeln('Сумма его цифр: ', summa);
   writeln('Произведение его цифр: ', multi);
end.
