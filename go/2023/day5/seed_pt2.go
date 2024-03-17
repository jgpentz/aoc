package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
)

func read_input2(fname string) ([][]int, map[string][][]int, []string) {
    var seeds [][]int
    var maps = make(map[string][][]int)
    var map_names []string
    var cur_map_name string

    file, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if (strings.Contains(line, ":")) {
            // Store all seeds from the line with a prefix of "seeds: "
            if (strings.Contains(line, "seeds")) {
                seeds_str := strings.Split(strings.Split(line, ": ")[1], " ")
                for iseed, seed_str := range seeds_str {
                    if iseed % 2 == 0 {
                        seed_int, err := strconv.Atoi(seed_str)
                        if err != nil {
                            log.Fatal(err)
                        }
                        seed_range, err := strconv.Atoi(seeds_str[iseed + 1])
                        if err != nil {
                            log.Fatal(err)
                        }
                        
                        seeds = append(seeds, []int{seed_int, seed_range})
                    }
                }
            } else {
                // Create a new map entry
                cur_map_name = strings.Fields(line)[0]
                map_names = append(map_names, cur_map_name)
                maps[cur_map_name] = [][]int{}
            }
        } else if len(line) > 0 {
            // Store values for the current map
            values := strings.Fields(line)

            dst, err := strconv.Atoi(values[0])
            if err != nil {
                log.Fatal(err)
            }
            src, err := strconv.Atoi(values[1])
            if err != nil {
                log.Fatal(err)
            }
            cur_range, err := strconv.Atoi(values[2])
            if err != nil {
                log.Fatal(err)
            }

            map_entry := []int{dst, src, cur_range}
            maps[cur_map_name] = append(maps[cur_map_name], map_entry)
        }
    }

    return seeds, maps, map_names
}


func backwards(seeds [][]int, maps map[string][][]int, map_names []string) int {
    for i, j := 0, len(map_names) - 1; i < j; i, j = i + 1, j - 1{
        map_names[i], map_names[j] = map_names[j], map_names[i]
    }

    i := 0
    for {
        starting := i
        input := i
        for _, map_name := range map_names {
            for _, map_entry := range maps[map_name] {
                // Check if input is within source range, if it is then the destination
                // is destination plus the new offset (input - source), otherwise
                // the destination is just equal to the input
                if (map_entry[0] <= input) && (input <= map_entry[0] + map_entry[2]) {
                    input = map_entry[1] + (input - map_entry[0])
                    break
                }
            }
        }
        
        for _, seed := range seeds {
            if (seed[0] <= input) && (input <= seed[0] + seed[1]) {
                return starting
            }
        } 

        i += 1
    }
}

func main() {
    seeds, maps, map_names := read_input2("input_test.txt")

    lowest_location := backwards(seeds, maps, map_names)
    if lowest_location != 46 {
        panic(fmt.Sprintf("Test case failed! Expected: 46, computed: %v", lowest_location))
    }

    seeds, maps, map_names = read_input2("input.txt")

    lowest_location = backwards(seeds, maps, map_names)
    fmt.Println("Lowest location number is: ", lowest_location)
}
