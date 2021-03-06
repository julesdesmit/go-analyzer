// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/julesdesmit/go-analyzer/pkg/runner"
	"github.com/urfave/cli/v2"
)

// CLIFlags contains all possible application flags.
var CLIFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "run all custom lints",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "license",
		Aliases: []string{"l"},
		Usage:   "check if all source files have the correct license header",
		Value:   false,
	},
}

func main() {
	app := &cli.App{
		Name:      "go-analyzer",
		Usage:     "Performs custom lint checks on Golang repositories.",
		Copyright: "Copyright (c) 2022 Jules de Smit",
		Version:   "0.1.0",
		Flags:     CLIFlags,
		Action:    action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func action(ctx *cli.Context) error {
	if ctx.NumFlags() == 0 {
		cli.ShowAppHelpAndExit(ctx, 1)
	}

	dirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if ctx.Bool("all") {
		if errs := runner.RunAll(dirPath); errs != nil {
			for _, err := range errs {
				fmt.Println(err)
			}

			return errors.New("lint failed")
		}

		fmt.Println("lint successful!")
		return nil
	}

	for _, flag := range CLIFlags {
		if ctx.Bool(flag.String()) {
			if errs := runner.Run(flag.String(), dirPath); errs != nil {
				for _, err := range errs {
					fmt.Println(err)
				}

				return errors.New("lint failed")
			}
		}
	}

	fmt.Println("lint successful!")
	return nil
}
