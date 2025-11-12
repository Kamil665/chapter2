using System;

class Program {
    static void Main() {
        int pr = 542 * 130;
        int square = 130;
        
        int area = square * square;
        
        int res = pr / area;
        
        Console.WriteLine($"{res} квадрата(ов) можно отрезать от прямоугольника с размерами 542 * 130");
    }
}
