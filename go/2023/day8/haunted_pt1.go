package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
)

func read_input(fname string) (string, map[string][]string) {
    var directions string
    desert_map := make(map[string][]string)

    f, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    iline := 0
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        if iline == 0 {
            directions = scanner.Text()
            iline += 1
        } else {
            if len(scanner.Text()) == 0 {
                continue
            }
            current_node:= strings.Split(scanner.Text(), " = ")[0]
            next_nodes := strings.Split(scanner.Text(), " = ")[1]
            next_l := strings.Split(next_nodes, ", ")[0][1:]
            next_r := strings.Split(strings.Split(next_nodes, ", ")[1], ")")[0]

            desert_map[current_node] = []string{next_l, next_r}
        }
    }

    return directions, desert_map
}

func walk(dirs string, desert_map map[string][]string) int {
    steps := 0
    cur_loc := "AAA"
    dirs_index := 0

    for cur_loc != "ZZZ" {
        steps += 1
        
        if string(dirs[dirs_index]) == "L" {
            cur_loc = desert_map[cur_loc][0]
        } else {

            cur_loc = desert_map[cur_loc][1]
        }

        dirs_index += 1
        if dirs_index == len(dirs) {
            dirs_index = 0
        }
    }

    return steps
}

func main() {
    directions, desert_map := read_input("input_test.txt")
    steps := walk(directions, desert_map)

    fmt.Println(steps)

    directions, desert_map = read_input("input.txt")
    steps = walk(directions, desert_map)

    fmt.Println(steps)
}
