package main
import "fmt"

func main() {
    var num int
    fmt.Print("Введите натуральное число n (n > 99): ")
    fmt.Scan(&num)
    
    if num <= 99 {
        fmt.Print("Ошибка! Число должно быть больше 99!")
    } else {
        var sot int = num / 100
        var des int = (num % 100) / 10
        
        fmt.Println("Число десятков: ", des)
        fmt.Println("Число сотен: ", sot)
    }
}
