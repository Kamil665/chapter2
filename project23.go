package main
import "fmt"

func main() {
    var num int
    fmt.Print("Введите натуральное число n (n > 999): ")
    fmt.Scan(&num)
    
    if num <= 999 {
        fmt.Print("Ошибка! Число должно быть больше 999!")
    } else {
        var tis int = num / 1000
        var sot int = (num % 1000) / 100
        
        fmt.Println("Число сотен: ", sot)
        fmt.Println("Число тысяч: ", tis)
    }
}
