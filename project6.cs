using System;

class Program {
    static void Main() {
        Console.Write("Введите кол-во секунд с начала суток: ");
        int sek = Convert.ToInt32(Console.ReadLine());
        
        //а)
        int hours = sek / 3600;
        Console.WriteLine($"Полных часов прошло: {hours}");
        
        //б)
        int min = (sek % 3600) / 60;
        Console.WriteLine($"Полных минут прошло с начала очередного часа: {min}");
        
        //в)
        int s = (sek % 3600) % 60;
        Console.WriteLine($"Полных секунд прошло с начала очередной минуты: {s}");
    }
}
