package main


import (
    "flag"
    "fmt"
    "os"
    "runtime"
)


func main() {

    flag.Parse()
    inFiles := flag.Args()

    if len(inFiles) != 2 {
        fmt.Printf(`Usage: %s <config.toml> <input.vcf>`, os.Args[0])
        fmt.Println()
        os.Exit(2)
    }

    tomlFile := inFiles[0]
    queryFile := inFiles[1]

    // REQUIRED
	_, err_query := os.Stat(queryFile)
    if err_query != nil {
        fmt.Fprintf(os.Stderr, "\nERROR: can't find query file: %s\n", queryFile)
        os.Exit(2)
    }

    // Required
    runtime.GOMAXPROCS(4)

    // Trigger
    _, err_toml := os.Open(tomlFile)
    if err_toml != nil {
        panic(err_toml)
    }

}
