package main
import "fmt"

func main() {
    var n int
    fmt.Print("Введите натуральное число n (n > 9): ")
    fmt.Scan(&n)
    
    if n <= 9 {
        fmt.Print("Ошибка! Число должно быть больше 9!")
    } else {
        var num1 int = n / 10
        var num2 int = n % 10
        
        fmt.Println("Число единиц: ", num2)
        fmt.Println("Число десятков: ", num1)
    }
}
