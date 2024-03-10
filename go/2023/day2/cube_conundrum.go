package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "slices"
    "regexp"
    "strconv"
)

type DayEntry struct {
    blue []int
    red []int
    green []int
}

func ReadInput() map[int]DayEntry{
    dayData := make(map[int]DayEntry)

    file, err := os.Open("./input.txt")
    if err != nil {
        fmt.Println("Error occured during reading")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Create a custom split function that separates on semicolons
    onSemicolon := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        for i := 0; i < len(data); i++ {
            if data[i] == ';' {
                return i + 1, data[:i], nil
            }
        }
        if !atEOF {
            return 0, nil, nil
        }
        return 0, data, bufio.ErrFinalToken
    }

    onComma := func(data []byte, atEOF bool) (advace int, token []byte, err error) {
        for i := 0; i < len(data); i++ {
            if data[i] == ',' {
                return i + 1, data[:i], nil
            }
        }
        if !atEOF {
            return 0, nil, nil
        }
        return 0, data, bufio.ErrFinalToken
    }

    // Go through each line, storing the semicolon separated items
    dayNumber := 1
    for scanner.Scan() {
        semicolonScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
        semicolonScanner.Split(onSemicolon)

        for semicolonScanner.Scan() {
            commaScanner := bufio.NewScanner(strings.NewReader(semicolonScanner.Text()))
            commaScanner.Split(onComma)

            for commaScanner.Scan() {
                // Read out the text e.g. '4 blue'
                text := strings.TrimSpace(commaScanner.Text())

                pattern := `^(\d+)\s+([a-zA-Z]+)$`
                regex := regexp.MustCompile(pattern)
                matches := regex.FindStringSubmatch(text)
                
                if len(matches) == 3 {
                    valStr := matches[1]
                    color := matches[2]

                    val, _ := strconv.Atoi(valStr)


                    switch color {
                    case "blue":
                        if entry, ok := dayData[dayNumber]; ok {
                            entry.blue = append(entry.blue, val)
                            dayData[dayNumber] = entry

                        } else {
                            newEntry := DayEntry {
                                blue: []int{val},
                            }
                            dayData[dayNumber] = newEntry
                        }
                    case "red":
                        if entry, ok := dayData[dayNumber]; ok {
                            entry.red = append(entry.red, val)
                            dayData[dayNumber] = entry
                        } else {
                            newEntry := DayEntry {
                                red: []int{val},
                            }
                            dayData[dayNumber] = newEntry
                        }
                    case "green":
                        if entry, ok := dayData[dayNumber]; ok {
                            entry.green = append(entry.green, val)
                            dayData[dayNumber] = entry
                        } else {
                            newEntry := DayEntry {
                                green: []int{val},
                            }
                            dayData[dayNumber] = newEntry
                        }
                    }
                }
            }
        }
        dayNumber += 1
    }
    
    return dayData
}

func CountCubes(data map[int]DayEntry) int {
    sum := 0
    maxRed := 12
    maxGreen := 13
    maxBlue := 14

    for day, entry := range data {
        dayValid := true
        for _, val := range entry.blue {
            if val > maxBlue {
                dayValid = false
                break
            }
        }
        for _, val := range entry.red {
            if val > maxRed {
                dayValid = false
                break
            }
        }
        for _, val := range entry.green {
            if val > maxGreen {
                dayValid = false
                break
            }
        }
        
        if dayValid {
            sum += day
        }
    } 

    return sum
}

func FewestCubes(data map[int]DayEntry) int {
    sum := 0

    for _, entry := range data {
        maxBlue := slices.Max(entry.blue)
        maxRed := slices.Max(entry.red)
        maxGreen := slices.Max(entry.green)

        power := maxBlue * maxRed * maxGreen
        sum += power
    }

    return sum
}

func main () {
    data := ReadInput()
    sum := CountCubes(data)
    fmt.Println("The sum is:", sum)
    sum_fewest := FewestCubes(data)
    fmt.Println("The sum is:", sum_fewest)
}
