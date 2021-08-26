package handler

import (
	"demo/app/cli/internal/svc"
	"github.com/urfave/cli/v2"
)

func RegisterHandlers(app *cli.App, serverCtx *svc.ServiceContext) {
	outputFlag := &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Value:   "",
		Usage:   "output file path",
	}
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list data",
			Action:  ListHandler(serverCtx),
			Flags:   []cli.Flag{outputFlag},
		},
		&cli.Command{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create dta",
			Action:  CreateHandler(serverCtx),
			Flags:   []cli.Flag{outputFlag},
		},
	)
}
