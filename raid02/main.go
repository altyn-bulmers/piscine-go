    package main

    import (
        "fmt"
  //      "os"

        "github.com/01-edu/z01"
    )

    func main() {
  //      arguments := os.Args[1:]
        arguments := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
        //arguments = []string{".964....1.", "1...6.5.4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}

        if checkInput(arguments) == true {
            table := [9][9]rune{}
            table = fillTable(table, arguments)

            if isSolve(&table) == true {
                for y := 0; y < 9; y++ {
                    for x := 0; x < 9; x++ {
                        if x != 8 {
                            z01.PrintRune(rune(table[y][x]))
                            z01.PrintRune(32)
                        } else {
                            z01.PrintRune(rune(table[y][x]))
                        }
                    }
                    z01.PrintRune(10)
                }
            } else {
                fmt.Println("Error") // checkInput is not True
            }
        }
    }



    func checkInput(args []string) bool {
        if len(args) != 9 {
            fmt.Println("Error") // Input length is out of range
            return false
        }

        for i := 0; i < len(args); i++ {
            if len(args[i]) != 9 {
                fmt.Println("Error") //  Input length is out of range
                return false
            }
        }
        for i := 0; i < len(args); i++ {
            for _, value := range args[i] {
                if value == 47 || value == 48 {
                    fmt.Println("Error") // Input is not correct
                    return false
                } else if value < 46 || value > 57 {
                    fmt.Println("Error") // Input is not correct
                    return false
                }
            }
        }
        return true
    }

    //fills with input arguments
    func fillTable(table [9][9]rune, args []string) [9][9]rune {
        for i := range args {
            for j := range args[i] {
                table[i][j] = rune(args[i][j])
            }
        }
        return table
    }

    //counts remaining empty cells
    func isDots(table *[9][9]rune) bool {
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                if table[i][j] == '.' {
                    return true
                }
            }
        }
        return false
    }

    // check if it fits

    func isValid(table *[9][9]rune, x int, y int, z rune) bool {
        // check double int
        for i := 0; i < 9; i++ {
            if z == table[i][x] {
                return false
            }
        }

        for j := 0; j < 9; j++ {
            if z == table[y][j] {
                return false
            }
        }

        //square check
        a := x / 3
        b := y / 3

        for k := 3 * a; k < 3*(a+1); k++ {
            for l := 3 * b; l < 3*(b+1); l++ {
                if z == table[l][k] {
                    return false
                }
            }
        }
        return true
    }

    //backtracking solver

    func isSolve(table *[9][9]rune) bool {
        if !isDots(table) {
            return true
        }
        for y := 0; y < 9; y++ {
            for x := 0; x < 9; x++ {
                if table[y][x] == '.' {
                    for z := '1'; z <= '9'; z++ {
                        if isValid(table, x, y, z) {
                            table[y][x] = z
                            if isSolve(table) {
                                return true
                            }
                        }
                        table[y][x] = '.'
                    }
                    return false
                }
            }
        }
        return false
    }