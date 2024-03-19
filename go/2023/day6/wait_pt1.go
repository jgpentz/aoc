package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func read_records(fname string) map[string][]int {
    records := make(map[string][]int)
    f, err := os.Open(fname)
    if err != nil {
        log.Fatal("Error reading file")
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        title := strings.Split(scanner.Text(), ":")[0]
        vals := strings.Fields(strings.Split(scanner.Text(), ":")[1])

        int_vals := make([]int, len(vals))

        for i, str := range vals {
            num, err := strconv.Atoi(str)
            if err != nil {
                fmt.Printf("Error converting string to integer: %v\n", err)
            }

            int_vals[i] = num
        }
        records[title] = int_vals
    }

    return records
}

func find_times(records map[string][]int) int {
    sum := 1

    for i, time := range records["Time"] {
        record := records["Distance"][i]

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
    }

    return sum
}

func main() {
    recs := read_records("input_test.txt")
    sum := find_times(recs)
    if sum != 288 {
        panic(fmt.Sprintf("Test case failed! Expected: 288, computed: %v", sum))
    }

    recs = read_records("input.txt")
    sum = find_times(recs)

    fmt.Printf("output: %v\n", sum)
}
