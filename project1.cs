using System;

class Program {
    static void Main() {
        Console.Write("Введите расстояние в сантиметрах: ");
        double cm = Convert.ToDouble(Console.ReadLine());
        
        double metr = cm / 100;
        
        Console.WriteLine($"Расстояние в метрах: {metr}");
    }
}
