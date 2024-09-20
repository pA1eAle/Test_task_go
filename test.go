package main

import (
    "fmt"
)


func getInput() (int8, int8, string) {
    var a, b int8
    var operator string

    fmt.Print("Input: ")
    fmt.Scanln(&a, &operator, &b)

    return a, b, operator
}


func main() {
    a, b, operator := getInput()

    var result int8

    if a > 0 {
        if a <= 10 {
            if b > 0 {
                if b <= 10 {
                    switch operator {
                    case "+":
                        result = a + b
                    case "-":
                        result = a - b
                    case "*":
                        result = a * b
                    case "/":
                        if b == 0 {
                            fmt.Println("Ошибка: деление на ноль!")
                            result = 0
                        }
                        result = a / b
                    default:
                        fmt.Println("Неизвестный оператор")
                        return
                    }
                } else {
                    fmt.Println("Некорректное число")
                }
            } else {
                fmt.Println("Некорректное число")
            }
        } else {
            fmt.Println("Некорректное число")
        }
        
    } else {
        fmt.Println("Некорректное число")
    }

    fmt.Printf("Output: %d", result)
}


