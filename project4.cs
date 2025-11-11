using System;

class Program {
    static void Main() {
        Console.Write("Введите расстояние в метрах: ");
        double metr = Convert.ToDouble(Console.ReadLine());
        
        double km = metr / 1000;
        
        Console.WriteLine($"Расстояние в километрах: {km}");
    }
}
