package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
)

func main() {
    var r *bufio.Reader

    if len(os.Args) > 1 {
        args := strings.Join(os.Args[1:], "\n")
        r = bufio.NewReader(strings.NewReader(args))
    } else {
        r = bufio.NewReader(os.Stdin)
    }

    roll(r)
}

func roll(in *bufio.Reader) {
    for {
        r, _, err := in.ReadRune()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

        fmt.Print(string(r))
    }
}
