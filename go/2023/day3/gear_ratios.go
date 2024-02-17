package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func isLastLine(scanner *bufio.Scanner)  bool {
    if hasNext := scanner.Scan(); !hasNext {
        return true
    }
    return false
}

func SearchSurrounding(prevLine string, curLine string, nextLine string, idx int) bool {
    // Regex pattern for the special chars
    re := regexp.MustCompile(`[^A-Za-z0-9\.]`)
    var start_idx, end_idx int

    // Set the start and end indices (beginning of line, middle, end of line)
    if idx == 0 {
        start_idx = idx
        end_idx = idx + 1
    } else if idx == len(curLine) - 1 {
        start_idx = idx - 1
        end_idx = idx
    } else {
        start_idx = idx - 1
        end_idx = idx + 1
    }

    // Loop over the 3 chars adjacent on the prev and next line, 
    for i := start_idx; i < end_idx + 1; i++ {
        // Check the cur line left and right
        if curLine != "" {
            if i == idx - 1 || i == idx + 1 {
                matched := re.MatchString(string(curLine[i]))
                if matched {
                    return true
                }
            }
        }

        // Check the prev line
        if prevLine != "" {
            matched := re.MatchString(string(prevLine[i]))
            if matched {
                return true
            }
        }

        // Check the next line
        if nextLine != ""{
            matched := re.MatchString(string(nextLine[i]))
            if matched {
                return true
            }
        }
    }

    return false
}

func GetNumber(cur_line string, cur_idx int) (int, int) {
    cur_number := ""
    start_idx := 0
    re := regexp.MustCompile(`[0-9]`)

    // Search for beginning of number
    for i := cur_idx; i >= 0; i-- {
        if match := re.MatchString(string(cur_line[i])); !match {
            start_idx = i + 1
            break
        } 
    }

    // Gather up the digits in this number
    j := start_idx;
    for ; j < len(cur_line); j++ {
        if match := re.MatchString(string(cur_line[j])); match {
            cur_number += string(cur_line[j])
        } else {
            break
        }
    }

    // Convert cur_number to integer
    curNumberInt, err := strconv.Atoi(cur_number)
    if err != nil {
        fmt.Println("Error converting cur_number to int:", err)
        return -1, j
    }

    return curNumberInt, j
}

func FindRatio() int {
    sum := 0

    // Open the file and close it when this function finishes
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error occured when opening file")
        return 0
    }
    defer file.Close()

    // Create a scanner and read line by line
    scanner := bufio.NewScanner(file)

    line_idx := 0
    prev_line := ""
    cur_line := ""
    next_line := ""

    for scanner.Scan() {
        next_line = scanner.Text()
        
        if line_idx == 0 {
            cur_line = next_line
            line_idx += 1
            continue
        }

        // Search to see if this number is valid, and add it to the sum
        i := 0
        for i < len(cur_line) {
            if _, err := strconv.Atoi(string(cur_line[i])); err == nil {
                match := SearchSurrounding(prev_line, cur_line, next_line, i)
                if match {
                    curNumberInt, lastNumIdx := GetNumber(cur_line, i)
                    sum += curNumberInt

                    // Advance i to end of this number
                    i = lastNumIdx
                    continue
                }
            }

            i += 1
        }

        prev_line = cur_line
        cur_line = next_line
        line_idx += 1
    }

    // Search the last line
    next_line = ""
    i := 0
    for i < len(cur_line) {
        if _, err := strconv.Atoi(string(cur_line[i])); err == nil {
            match := SearchSurrounding(prev_line, cur_line, next_line, i)
            if match {
                curNumberInt, lastNumIdx := GetNumber(cur_line, i)
                sum += curNumberInt

                // Advance i to end of this number
                i = lastNumIdx
                continue
            }
        } 

        i += 1
    }

    return sum
}

func main() {
    sum := FindRatio()
    fmt.Println("The sum is: ", sum)
}
