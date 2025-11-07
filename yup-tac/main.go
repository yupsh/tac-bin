package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tac"
)

const (
	flagSeparator = "separator"
	flagBefore    = "before"
	flagRegex     = "regex"
)

func main() {
	app := &cli.App{
		Name:  "tac",
		Usage: "concatenate and print files in reverse",
		UsageText: `tac [OPTIONS] [FILE...]

   Write each FILE to standard output, last line first.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagSeparator,
				Aliases: []string{"s"},
				Usage:   "use STRING as the separator instead of newline",
			},
			&cli.BoolFlag{
				Name:    flagBefore,
				Aliases: []string{"b"},
				Usage:   "attach the separator before instead of after",
			},
			&cli.BoolFlag{
				Name:    flagRegex,
				Aliases: []string{"r"},
				Usage:   "interpret the separator as a regular expression",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "tac: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, yup.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagSeparator) {
		params = append(params, Separator(c.String(flagSeparator)))
	}
	if c.Bool(flagBefore) {
		params = append(params, Before)
	}
	if c.Bool(flagRegex) {
		params = append(params, Regex)
	}

	// Create and execute the tac command
	cmd := Tac(params...)
	return yup.Run(cmd)
}
