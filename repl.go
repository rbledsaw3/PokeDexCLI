package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func startRepl() {
    reader := bufio.NewScanner(os.Stdin)
    if reader == nil {
        fmt.Println("Error creating reader")
        return
    }
    for {
        fmt.Print("Pokedex > ")
        reader.Scan()

        words := cleanInput(reader.Text())
        if len(words) == 0 {
            continue
        }
        
        commandName := words[0]
        fmt.Printf("Your command was: %s\n", commandName)
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

