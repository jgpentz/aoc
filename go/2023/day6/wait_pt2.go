package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func read_records2(fname string) map[string]int {
    records := make(map[string]int)
    f, err := os.Open(fname)
    if err != nil {
        log.Fatal("Error reading file")
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        title := strings.Split(scanner.Text(), ":")[0]
        vals := strings.Fields(strings.Split(scanner.Text(), ":")[1])
        val := strings.Join(vals, "")
        num, err := strconv.Atoi(val)
        if err != nil {
            fmt.Printf("Error converting string to integer: %v\n", err)
        }

        records[title] = num
    }

    return records
}

func find_times2(records map[string]int) int {
    sum := 1

    time := records["Time"]
    record := records["Distance"]

    for j := 0; j < time; j++ {
        if (time - j) * j > record {
            if time % 2 == 0 {
                sum = sum * ((((time / 2) - j) * 2) + 1)
            } else {
                sum = sum * ((((time / 2) - j) + 1) * 2)
            }

            break
        }
    }

    return sum
}

func main() {
    recs := read_records2("input_test.txt")
    sum := find_times2(recs)
    if sum != 71503 {
        panic(fmt.Sprintf("Test case failed! Expected: 71503, computed: %v", sum))
    }

    recs = read_records2("input.txt")
    sum = find_times2(recs)

    fmt.Printf("output: %v\n", sum)
}
