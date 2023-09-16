package main


import (
    "flag"
    "fmt"
    "os"
    "runtime"

    //"github.com/BurntSushi/toml"
    "github.com/brentp/goluaez"
    "github.com/brentp/xopen"
)


// shared/shared.go
type Config struct {
    Annotation     []Annotation
    PostAnnotation []PostAnnotation
    // base path to prepend to all files.
    Base string
}

// shared/shared.go
// Annotation holds information about the annotation files parsed from the toml config.
type Annotation struct {
    File    string
    Ops     []string
    Fields  []string
    Columns []int
    // the names in the output.
    Names []string
}

// api/api.go
type PostAnnotation struct {
    Fields []string
    Op     string
    Name   string
    Type   string

    code string

    // use 8 of these to avoid contention in parallel contexts.
    mus [8]chan int
    Vms [8]*goluaez.State
}


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
    // github.com/brentp/xopen v0.0.0-20181116180855-111b45cadc7d
    if !(xopen.Exists(queryFile) || queryFile == "") {
        fmt.Fprintf(os.Stderr, "\nERROR: can't find query file: %s\n", queryFile)
        os.Exit(2)
    }

    // Required
    runtime.GOMAXPROCS(4)

    // Trigger
    _, err := os.Open(tomlFile)
    if err != nil {
        panic(err)
    }

}
