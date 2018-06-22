package main

import (
    "fmt"
    "errors"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func add(args ... int) (int) {
    result := 0

    for _, arg := range args {
        result  += arg
    }
    return result
}

func sub(args ... int) (int) {
    result := 0

    for _, arg := range args {
        result  -= arg
    }
    return result
}

func mul(args ... int) (int) {
    result := args[0]

    for _, arg := range args[1:] {
        result  *= arg
    }
    return result
}

func div(args ... int) (int, error) {
    result := args[0]

    for _, arg := range args[1:] {
        if arg == 0 {
            return 0, errors.New("zero division")
        }
        result /= arg
    }
    return result, nil
}

func toInt(arg []string) ([]int) {
    result := make([]int, len(arg))

    for i,v := range arg {
        val,err := strconv.Atoi(v)
        if err == nil {
            result[i] = val
        } else {
            fmt.Println(err)
        }
    }
    return result
}


func main() {
    fmt.Println("Hello world")
    line := ""
    in := bufio.NewReader(os.Stdin)
    for line != "-1" {
        line,err := in.ReadString('\n')

        if err == nil {
            line = line[:len(line) - 1]
            strs := strings.Split(line, " ")
            operation := strs[0]
            args := toInt(strs[1:])
            var result int

            switch operation {
            case "sub":
                result = sub(args...)
            case "mul":
                result = mul(args...)
            case "div":
                result,_ = div(args...)
            default:
                result = add(args...)
            }
            fmt.Println(result)
        } else {
            fmt.Println(err)
        }
    }
}
