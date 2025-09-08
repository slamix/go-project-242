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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args().Slice()
			if len(args) == 0 {
				return fmt.Errorf("filepath is required")
			}

			path := args[0]

			res, err := code.GetPathSize(path, false, false, false)
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
