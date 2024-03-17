package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
)

func read_input(fname string) ([]int, map[string][][]int, []string) {
    var seeds []int
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
                for _, seed_str := range seeds_str {
                    seed_int, err := strconv.Atoi(seed_str)
                    if err != nil {
                        log.Fatal(err)
                    }

                    seeds = append(seeds, seed_int)
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

func find_lowest_location(seeds []int, maps map[string][][]int, map_names []string) int {
    outputs := seeds
    for _, map_name := range map_names {
        inputs := outputs
        outputs = []int{}
        for _, input := range inputs {
            new_output := input
            for _, map_entry := range maps[map_name] {
                // Check if input is within source range, if it is then the destination
                // is destination plus the new offset (input - source), otherwise
                // the destination is just equal to the input
                if (map_entry[1] <= input) && (input <= map_entry[1] + map_entry[2]) {
                    new_output = map_entry[0] + (input - map_entry[1])
                    break
                }
            }

            outputs = append(outputs, new_output)
        }
    }

    lowest_location := outputs[0]
    for _, output := range outputs {
        if output < lowest_location {
            lowest_location = output
        }
    }

    return lowest_location
}

func main() {
    seeds, maps, map_names := read_input("input_test.txt")

    lowest_location := find_lowest_location(seeds, maps, map_names)
    if lowest_location != 35 {
        panic(fmt.Sprintf("Test case failed! Expected: 35, computed: %v", lowest_location))
    }

    seeds, maps, map_names = read_input("input.txt")

    lowest_location = find_lowest_location(seeds, maps, map_names)
    fmt.Println("Lowest location number is: ", lowest_location)
}
