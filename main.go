package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	hy "github.com/nod/halyard/halyard"
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

    app := &cli.App{
/*
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("added task: ", cCtx.Args().First())
                    return nil
                },
            },
            {
                Name:    "complete",
                Aliases: []string{"c"},
                Usage:   "complete a task on the list",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("completed task: ", cCtx.Args().First())
                    return nil
                },
            },
            {
                Name:    "template",
                Aliases: []string{"t"},
                Usage:   "options for task templates",
                Subcommands: []*cli.Command{
                    {
                        Name:  "add",
                        Usage: "add a new template",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("new task template: ", cCtx.Args().First())
                            return nil
                        },
                    },
                    {
                        Name:  "remove",
                        Usage: "remove an existing template",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("removed task template: ", cCtx.Args().First())
                            return nil
                        },
                    },
                },
            },
        },
*/
    }

	log := hy.GetLogger()

    cfg := hy.GetConfig()
    cfg.HttpListenUri = "localhost:6488"
    cfg.DbUri = "/tmp/tmp.db"
    hy.StartHTTPServer(cfg.HttpListenUri)

    if err := app.Run(os.Args); err != nil {
        log.Error(fmt.Sprintf("%v", err))
    }
}
