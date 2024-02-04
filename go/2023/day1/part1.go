package main
import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

func ReadData() ([]string) {
    file, err := os.Open("./input.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, "reading standard imput:", err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard imput:", err)
    }

    return lines
}

func isDigit(char rune) bool {
    return unicode.IsDigit(char)
}

func charToInt(char rune) int {
    return int(char - '0')
}

func CalculateSum(lines []string) int {
    sum := 0
    for _, line := range lines {
        var first, last int
        foundFirst := false

        for _, char := range line {
            // Get the first and last digit
            if isDigit(char) {
                if(!foundFirst) {
                    foundFirst = true
                    first = charToInt(char)
                }

                last = charToInt(char)
            }
        }

        // Combine the digits
        var digit = (first * 10) + last

        // Add the new number to the sum
        sum += digit
    }

    return sum
}

func main() {
    lines := ReadData()
    sum := CalculateSum(lines)
    fmt.Printf("The sum is: %d\n", sum)
}
