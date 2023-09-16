package main


import (
    "flag"
    "fmt"
    "os"
)


func main() {

    flag.Parse()
    inFiles := flag.Args()

    if len(inFiles) != 2 {
        fmt.Printf(`Usage: %s <file_a> <file_b>`, os.Args[0])
        fmt.Println()
        os.Exit(2)
    }

    file_a := inFiles[0]
    file_b := inFiles[1]

    _, err_a := os.Stat(file_a)
    if err_a != nil {
        fmt.Fprintf(os.Stderr, "\nERROR: failed os.Stat call on: %s\n", file_a)
        panic(err_a)
    }

    _, err_b := os.Open(file_b)
    if err_b != nil {
        fmt.Fprintf(os.Stderr, "\nERROR: failed os.Open call on: %s\n", file_b)
        panic(err_b)
    }

}
