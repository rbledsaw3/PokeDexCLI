package main

import ("strings")

func cleanInput(text string) []string {
    // splits text into words based on whitespace and makes them lowercase
    words := []string{}
    for _, word := range strings.Fields(text) {
        words = append(words, strings.ToLower(word))
    }
    return words
}

func main() {

}
