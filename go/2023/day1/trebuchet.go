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

func SumWithLetters(lines []string) int {
    sum := 0
    digit_words := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    for _, line := range lines {
        var first, last int
        foundFirst := false

        for i, char := range line {
            if isDigit(char) {
                if (!foundFirst) {
                    foundFirst = true
                    first = charToInt(char)
                }

                last = charToInt(char)
            } else {
                for word, number := range digit_words {
                    found := true

                    // Check that word fits within remaining chars
                    if len(word) <= (len(line) - i){
                        for j, _ := range word {
                            if word[j] != line[i + j] {
                                found = false
                                break
                            }
                        }

                        if found {
                            if !foundFirst {
                                foundFirst = true
                                first = number
                            }

                            last = number

                            break
                        }
                    }
                }
            }
        }
        var digit = (first * 10) + last

        sum += digit
    }

    return sum
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
    sum1 := CalculateSum(lines)
    sum2 := SumWithLetters(lines)
    fmt.Printf("The sum without letters: %d\n", sum1)
    fmt.Printf("The sum with letters: %d\n", sum2)
}
