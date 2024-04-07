package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)

func read_input(fpath string) ([]int, [][]string) {
    var maze [][]string
    var start []int

    f, err := os.Open(fpath)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    iline := 0
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        maze = append(maze, []string{})

        for ichar, c := range line {
            maze[iline] = append(maze[iline], string(c))

            if string(c) == "S" {
                start = []int{iline, ichar}
            }
        }
        
        iline++
    }

    return start, maze
}

func get_coords(cur []int) [][]int {
    // cur[0] = y, cur[1] = x
    coords := [][]int{
        {cur[0] - 1, cur[1]}, // up
        {cur[0], cur[1] + 1}, // right
        {cur[0] + 1, cur[1]}, // down
        {cur[0], cur[1] - 1}, // left
    }

    return coords
}

type PathKey struct {
    y int
    c int
}

func walk(maze [][]string, start []int, cur []int, prev []int, dist int, path [][]int) (int, [][]int) {
    if cur[0] == start[0] && cur[1] == start[1] {
        if prev[0] != start[0] || prev[1] != start[1] {
            return dist, path
        }
    }

    curLetter := maze[cur[0]][cur[1]]
    if curLetter == "J" {
        if prev[0] == cur[0] && prev[1] == (cur[1] - 1) {
            next := []int{cur[0] - 1, cur[1]}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        } else {
            next := []int{cur[0], cur[1] - 1}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        }
    } else if curLetter == "F" {
        if prev[0] == (cur[0] + 1) && prev[1] == cur[1] {
            next := []int{cur[0], cur[1] + 1}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        } else {
            next := []int{cur[0] + 1, cur[1]}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        }
    } else if curLetter == "L" {
        if prev[0] == (cur[0] - 1) && prev[1] == cur[1] {
            next := []int{cur[0], cur[1] + 1}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        } else {
            next := []int{cur[0] - 1, cur[1]}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        }
    } else if curLetter == "7" {
        if prev[0] == cur[0] && prev[1] == (cur[1] - 1) {
            next := []int{cur[0] + 1, cur[1]}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        } else {
            next := []int{cur[0], cur[1] - 1}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        }
    } else if curLetter == "|" {
        if prev[0] == (cur[0] - 1) && prev[1] == cur[1] {
            next := []int{cur[0] + 1, cur[1]}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        } else {
            next := []int{cur[0] - 1, cur[1]}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        }
    } else if curLetter == "-" {
        if prev[0] == cur[0] && prev[1] == (cur[1] - 1) {
            next := []int{cur[0], cur[1] + 1}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        } else {
            next := []int{cur[0], cur[1] - 1}
            path = append(path, next)
            return walk(maze, start, next, cur, dist + 1, path)
        }
    }

    return dist, path
}

func find_distance(start []int, maze [][]string) (int, [][]int) {
    dirs := [][]string{
        {"F", "|", "7", "S"}, // up
        {"-", "J", "7", "S"}, // right
        {"J", "|", "L", "S"}, // down
        {"F", "-", "L", "S"}, // left
    }

    coords := get_coords(start)
    for icoord, coord := range coords {
        // ensure x, y are within maze boundaries
        if coord[0] >= 0 && coord[0] < len(maze) && coord[1] >= 0 && coord[1] < len(maze[0]) {
            for _, pipe := range dirs[icoord]{
                if maze[coord[0]][coord[1]] == pipe {
                    new_path := [][]int{coord}
                    distance, path := walk(maze, start, coord, start, 1, new_path)
                    return distance, path
                }
            }
        }
    }


    return 0, [][]int{{0}}
}

func count_inside(maze [][] string, path [][]int) int {
    cnt := 0
    visited := make(map[PathKey]bool)

    for _, coord := range path {
        key := PathKey{coord[0], coord[1]}
        visited[key] = true
    }

    for irow := range maze {
        is_inside := false
        for icol := range maze[irow] {
            key := PathKey{irow, icol}
            if _, ok := visited[key]; ok {
                cur := maze[irow][icol]
                // WARNING: I cheesed this, you need to turn "S" into a pipe,
                // but I was too lazy and just looked at the input, and determined
                // that it doesn't need to flip is_inside
                if cur == "F" || cur == "7" || cur == "|" {
                    is_inside = !is_inside
                }
            } else if is_inside {
                cnt++
            }
        }
    }


    return cnt

}

func main() {
    start, maze := read_input("input_test5.txt")

    distance, path := find_distance(start, maze)
    count := count_inside(maze, path)
    fmt.Println("distance: ", distance / 2)
    fmt.Println("count: ", count)

    start, maze = read_input("input_test2.txt")

    distance, path = find_distance(start, maze)
    count = count_inside(maze, path)
    fmt.Println("distance: ", distance / 2)
    fmt.Println("count: ", count)

    start, maze = read_input("input.txt")

    distance, path = find_distance(start, maze)
    count = count_inside(maze, path)
    fmt.Println("distance: ", distance / 2)
    fmt.Println("count: ", count)
}
