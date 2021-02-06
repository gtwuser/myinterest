package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"

	"os"
)

type Context struct {
	Debug bool
}

type RmCmd struct {
	Force     bool `help:"Force removal."`
	Recursive bool `help:"Recursively remove files." env:"CXDL_INGESTION_EXECUTION_ID"`

	Paths []string `arg name:"path" help:"Paths to remove." type:"path"`
}

func (r *RmCmd) Run(ctx *Context) error {
	log.Info().Msgf(`CXDL_INGESTION_EXECUTION_ID %v`, r.Recursive)
	fmt.Println("rm", r.Paths)
	return nil
}

type LsCmd struct {
	Paths []string `arg optional name:"path" help:"Paths to list." type:"path"`
}

func (l *LsCmd) Run(ctx *Context) error {
	fmt.Println("ls", l.Paths)
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Rm RmCmd `cmd help:"Remove files."`
	Ls LsCmd `cmd help:"List paths."`
}

func main() {
	os.Setenv(`CXDL_INGESTION_EXECUTION_ID`, `true`)
	ctx := kong.Parse(&cli,
		kong.Name(os.Args[0]),
		kong.Description("Cx Data Lake Ingestion"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{Compact: true, Summary: true}))
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})

	log.Info().Msgf("env %v", os.Getenv(`CXDL_INGESTION_EXECUTION_ID`))
	ctx.FatalIfErrorf(err)
}