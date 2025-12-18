package main
import "fmt"

func main() {
    var n int
    fmt.Print("Введите число n ( 1 <= n <= 999): ")
    fmt.Scanln(&n)
    
     if n < 1 || n > 999 {
        fmt.Println("Ошибка! Число n не должно быть меньше 1 или больше 999!")
    } else {
        var op1 int = n % 10
        var op2 int = n - op1
        var op3 int = op2 / 10
        var op4 int = op1 * 100
        var x int = op3 + op4
        
        fmt.Println("x = ", x)
    }
}
