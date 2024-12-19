package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/rbledsaw3/pokedexcli/internal/pokeapi"
)

type config struct {
    pokeapiClient       pokeapi.Client
    nextLocationsURL    *string
    prevLocationsURL    *string
}

func startRepl(cfg *config) {
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

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Printf("Unknown command")
            continue
        }
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

type cliCommand struct {
    name        string
    description string
    callback    func(*config) error
}

func getCommands() map[string]cliCommand{
    return map[string]cliCommand{
        "help": {
            name: "help",
            description: "Displays a help messsage",
            callback: commandHelp,
        },
        "map": {
            name: "map",
            description: "Displays the next page of location areas",
            callback: commandMapf,
        },
        "mapb": {
            name: "mapb",
            description: "Displays the previous page of location areas",
            callback: commandMapb,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
    }
}
