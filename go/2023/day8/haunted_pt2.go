package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
)

func read_input(fname string) (string, []string, map[string][]string) {
    var directions string
    desert_map := make(map[string][]string)
    var starting_nodes []string

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
            current_node := strings.Split(scanner.Text(), " = ")[0]
            next_nodes := strings.Split(scanner.Text(), " = ")[1]
            next_l := strings.Split(next_nodes, ", ")[0][1:]
            next_r := strings.Split(strings.Split(next_nodes, ", ")[1], ")")[0]

            if string(current_node[2]) == "A" {
                starting_nodes = append(starting_nodes, current_node)
            }

            desert_map[current_node] = []string{next_l, next_r}
        }
    }

    return directions, starting_nodes, desert_map
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func walk(dirs string, starting_nodes []string, desert_map map[string][]string) int {
    var lcms []int

    for _, node := range starting_nodes {
        steps := 0
        dirs_index := 0
        cur_node := node
        for string(cur_node[2]) != "Z" {
            steps += 1

            if string(dirs[dirs_index]) == "L" {
                cur_node = desert_map[cur_node][0]
            } else {

                cur_node = desert_map[cur_node][1]
            }

            dirs_index += 1
            if dirs_index == len(dirs) {
                dirs_index = 0
            }
        }

        lcms = append(lcms, steps)
    }

    lcm := LCM(lcms[0], lcms[1], lcms[2:]...)
    
    return lcm
}

func main() {
    directions, starting_nodes, desert_map := read_input("input_test2.txt")
    steps := walk(directions, starting_nodes, desert_map)
    fmt.Println(steps)

    directions, starting_nodes, desert_map = read_input("input.txt")
    steps = walk(directions, starting_nodes, desert_map)

    fmt.Println(steps)
}
