package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "   hello   world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "a b   v   x  r",
            expected: []string{"a", "b", "v", "x", "r"},
        },
        {
            input: " 13 9sdf jf \r",
            expected: []string{"13", "9sdf", "jf"},
        },
        {
            input: "Charmander Bulbasaur PIKACHU",
            expected: []string{"charmander", "bulbasaur", "pikachu"},
        },
    }
    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("cleanInput(%q) == %q, expected %q", c.input, actual, c.expected)
        }
        for i := range actual {
            word := actual[i]
            expected := c.expected[i]
            if word != expected {
                t.Errorf("cleanInput(%q) == %q, expected %q", c.input, actual, c.expected)
            }
        }
    }
}
