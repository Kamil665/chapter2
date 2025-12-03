package main
import "fmt"

func main() {
    var nach int = 237
    
    var op1 int = nach % 100
    var op2 int = op1 * 10
    var x int = op2 + (nach / 100)
    
    fmt.Println("Число конечное = ", nach)
    fmt.Println("Число x = ", x)
}
