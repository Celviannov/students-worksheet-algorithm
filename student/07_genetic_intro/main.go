package main

import "fmt"

func fitness(candidate string) int {
    // TODO: count the number of '1' characters.
    // Hint: for i := 0; i < len(candidate); i++ { ... }
    return 0
}

func main() {
    a, b := "1001", "0110"
    // TODO: crossover after character 2 to produce 1010.
    // TODO: mutate its second character to 1, producing 1110.
    fmt.Println(a, b, fitness(a))
}
