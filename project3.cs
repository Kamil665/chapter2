using System;

class Program {
    static void Main() {
        Console.Write("Введите массу в килограммах: ");
        double kg = Convert.ToDouble(Console.ReadLine());
        
        double tonn = kg / 1000;
        
        Console.WriteLine($"Масса в тоннах: {tonn}");
    }
}
