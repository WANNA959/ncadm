package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"ncadm/pkg/cmds"
	"os"
)

func main() {
	app := cmds.NewApp()
	app.Commands = []*cli.Command{
		cmds.NewCreateTokenCommand(),
		cmds.NewGetTokenCommand(),
		cmds.NewCheckConnStateCommand(),
		cmds.NewUnRegisterCommand(),
		cmds.NewCheckHealthCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error options: %s\n", err.Error())
		os.Exit(-1)
	}
}
