package main
import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
    "strings"
)

func read_input(fname string) [][]int {
    var history [][]int

    f, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        cur_str := strings.Fields(scanner.Text())
        cur_hist := []int{}
        for _, num_str := range cur_str {
            num_int, err := strconv.Atoi(num_str)
            if err != nil {
                log.Fatal(err)
            }
            cur_hist = append(cur_hist, num_int)
        }

        history = append(history, cur_hist)
    }

    return history
}

func predict(hist [][]int) int {
    sum := 0

    for _, cur := range hist {
        datasets := [][]int{cur}
        i := 0
        all_zeros := false
        for all_zeros == false {
            all_zeros = true
            for _, num := range datasets[i] {
                if num != 0 {
                    all_zeros = false
                }
            }
            
            if all_zeros == true {
                break
            }

            dataset := []int{}
            for inum := range datasets[i] {
                if inum != 0 {
                    dataset = append(dataset, datasets[i][inum] - datasets[i][inum - 1])
                }
            }
            
            datasets = append(datasets, dataset)

            i += 1
        }

        for i > 0 {
            target := datasets[i][0]
            prediction := datasets[i - 1][0] - target
            datasets[i - 1] = append([]int{prediction}, datasets[i - 1]...)

            i -= 1
        }
        sum += datasets[0][0]
    }
    
    return sum
}

func main() {
    hist := read_input("input_test.txt")
    sum := predict(hist)
    if sum != 2 {
        panic(fmt.Sprintf("Test case failed! Expected: 2, computed: %v", sum))
    }
    hist = read_input("input.txt")
    sum = predict(hist)
    fmt.Printf("The sum is: %v\n", sum)
}
