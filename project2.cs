using System;

class Program {
    static void Main() {
        Console.Write("Введите массу в килограммах: ");
        double kg = Convert.ToDouble(Console.ReadLine());
        
        double centner = kg / 100;
        
        Console.WriteLine($"Масса в центнерах: {centner}");
    }
}
