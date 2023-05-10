package main

import (
    "fmt"
//    "os"
)

var version string


/*
type RuntimeOpts struct {
    configFile string
    loglevel string
    showVersion bool
}

func ErrorAndExit(msg string) {
    fmt.Printf(
        "ERR: %s\nusage: %s",
        msg,
        os.Args[0],
    )
    os.Exit(2)
}
*/

func main() {
    fmt.Println(version)
}


