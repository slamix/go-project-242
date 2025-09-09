package main

import (
	code "code"
	"context"
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args().Slice()
			if len(args) == 0 {
				return fmt.Errorf("filepath is required")
			}

			path := args[0]
			human, all, recursive := cmd.Bool("human"), cmd.Bool("all"), cmd.Bool("recursive")
			res, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return err
			}
			fmt.Println(res)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
