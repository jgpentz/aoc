package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
    "log"
    "strconv"
    "slices"
)

type CamelHand struct {
    hand string
    bid int
}

// Define the order of playing cards
var cardOrder = map[string]int{
    "J": 0,
    "2":  1,
    "3":  2,
    "4":  3,
    "5":  4,
    "6":  5,
    "7":  6,
    "8":  7,
    "9":  8,
    "T":  9,
    "Q":  10,
    "K":  11,
    "A":  12,
}

func read_input(fname string) []CamelHand {
    var CamelHands []CamelHand

    f, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        var Hand CamelHand

        Hand.hand = strings.Fields(scanner.Text())[0]
        Hand.bid, err = strconv.Atoi(strings.Fields(scanner.Text())[1])
        if err != nil {
            log.Fatal(err)
        }

        CamelHands = append(CamelHands, Hand)
    }

    return CamelHands
}

func create_type_lists(Hands []CamelHand) map[string][]CamelHand {
    type_list := make(map[string][]CamelHand)
    for _, Hand := range Hands {
        card_count := make(map[string]int)
        largest_cnt := 0
        second_largest_cnt := 0
        num_pairs := 0

        for _, card := range Hand.hand {
            card_count[string(card)] += 1
        }

        for _, cnt := range card_count {
             // Full house
            if (largest_cnt == 3 && cnt == 2) || (largest_cnt == 2 && cnt == 3) {
                largest_cnt = 3
                second_largest_cnt = 2
            }

           if cnt > largest_cnt {
                largest_cnt = cnt
            }

            // Count pairs
            if cnt == 2 {
                num_pairs += 1
            }
        }
        
        if largest_cnt == 5 {
            type_list["five_kind"] = append(type_list["five_kind"], Hand)
        } else if largest_cnt == 4 {
            if card_count["J"] == 1  || card_count["J"] == 4{
                type_list["five_kind"] = append(type_list["five_kind"], Hand)
            } else {
                type_list["four_kind"] = append(type_list["four_kind"], Hand)
            }
        } else if largest_cnt == 3 && second_largest_cnt == 2 {
            if card_count["J"] == 3 || card_count["J"] == 2 {
                type_list["five_kind"] = append(type_list["five_kind"], Hand)
            } else {
                type_list["full_house"] = append(type_list["full_house"], Hand)
            }
        } else if largest_cnt == 3 {
            if card_count["J"] == 1 || card_count["J"] == 3 {
                type_list["four_kind"] = append(type_list["four_kind"], Hand)
            } else {
                type_list["three_kind"] = append(type_list["three_kind"], Hand)
            }
        } else if num_pairs == 2 {
            if card_count["J"] == 2 {
                type_list["four_kind"] = append(type_list["four_kind"], Hand)
            } else if card_count["J"] == 1 {
                type_list["full_house"] = append(type_list["full_house"], Hand)
            } else {
                type_list["two_pair"] = append(type_list["two_pair"], Hand)
            }
        } else if num_pairs == 1 {
            if card_count["J"] == 2 || card_count["J"] == 1 {
                type_list["three_kind"] = append(type_list["three_kind"], Hand)
            } else {
                type_list["one_pair"] = append(type_list["one_pair"], Hand)
            }
        } else {
            if card_count["J"] == 1 {
                type_list["one_pair"] = append(type_list["one_pair"], Hand)
            } else {
                type_list["high_card"] = append(type_list["high_card"], Hand)
            }
        }
    }

    return type_list
}

func sort_lists(type_list map[string][]CamelHand) {
    for list := range type_list {
        // Sort each list in increasing order of winning hands
        slices.SortFunc(type_list[list], func(a, b CamelHand) int {
            for i := range a.hand {
                if cardOrder[string(a.hand[i])] < cardOrder[string(b.hand[i])] {
                    return -1
                } else if cardOrder[string(a.hand[i])] > cardOrder[string(b.hand[i])] {
                    return 1
                }
            }

            return 0
        })
    }
}

func sum_ranks(type_list map[string][]CamelHand) int {
    sum := 0
    list_names := []string {
        "high_card", 
        "one_pair", 
        "two_pair", 
        "three_kind", 
        "full_house", 
        "four_kind", 
        "five_kind",
    }

    rank := 1
    for _, name := range list_names {
        for _, hand := range type_list[name] {
            sum += hand.bid * rank
            rank += 1
        }
    }

    return sum
}


func main() {
    // ---- TEST CASE -----
    Hands := read_input("input_test.txt")
    type_list := create_type_lists(Hands)
    sort_lists(type_list)
    sum := sum_ranks(type_list)

    if sum != 6839 {
        panic(fmt.Sprintf("Test case failed! Expected: 6839, computed: %v", sum))
    }

    // ---- REAL INPUT -----
    Hands = read_input("input.txt")
    type_list = create_type_lists(Hands)
    sort_lists(type_list)
    sum = sum_ranks(type_list)

    fmt.Printf("total winnings: %v\n", sum)
}
