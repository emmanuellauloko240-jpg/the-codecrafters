package main

import (
    "fmt"
    "strconv"
)

func add(a,b int) int{
    return a + b
}

func subtract(a,b int) int{
    return a-b
}

func multiply(a,b int) int{
    return a*b
}

func divide(a,b int) int{
 return a/b
}

func main(){

    for {
        var num1 string 
        fmt.Println("Enter number1: ")
        fmt.Scanln(&num1)
        a, err := strconv.Atoi(num1)
        if err != nil {
            fmt.Println("only want digit")
            continue
        }

        var num2 string
        fmt.Println("Enter number2: ")
        fmt.Scanln(&num2)

        b, err := strconv.Atoi(num2)
        if err != nil {
            fmt.Println("only want digit")
            continue
        }
        var operation string
        fmt.Println("choose operation(add, subtract, multiply, divide, Exit, Help)")
        fmt.Scanln(&operation)
  
        switch operation {
        case "add":
            fmt.Println(add(a, b))

        case "subtract": 
            fmt.Println(subtract(a, b))

        case "multiply":
            fmt.Println(multiply(a, b))

        case "divide":
            if a == 0 || b == 0 {
                fmt.Println("can not be divided by zero")
            }else {
            fmt.Println(multiply(a, b))
            }
        case "Exit": 
            fmt.Println("Goodbye")
            return

        case "Help":
            if operation == "Help"{
        
                fmt.Println("add +")
                fmt.Println("subtract -")
                fmt.Println("multiply *")
                fmt.Println("divide /")



            }
        default:
            fmt.Println("Invalid operation")
        }

    }
}