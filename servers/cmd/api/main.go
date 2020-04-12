package main

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
)

// Exit statuses
const (
	ExitOK  = 0
	ExitErr = 1
)

var (
	portFlag       = &cli.UintFlag{Name: "port", Aliases: []string{"p"}, Value: 8000}
	dbUserFlag     = &cli.StringFlag{Name: "db-user", Value: "", Required: true}
	dbPasswordFlag = &cli.StringFlag{Name: "db-password", Value: ""}
	dbHostFlag     = &cli.StringFlag{Name: "db-host", Value: "localhost"}
	dbPortFlag     = &cli.UintFlag{Name: "db-port", Value: 3306}
	dbNameFlag     = &cli.StringFlag{Name: "db-name", Value: "", Required: true}
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Usage = ""
	app.HideVersion = true

	app.Flags = []cli.Flag{
		portFlag,
		dbUserFlag,
		dbPasswordFlag,
		dbHostFlag,
		dbPortFlag,
		dbNameFlag,
	}
	app.Action = execute

	return app
}

func main() {
	app := newApp()
	os.Exit(printOnError(app.Run(os.Args)))
}

func printOnError(err error) int {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return ExitErr
	}
	return ExitOK
}
