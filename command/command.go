package command

import (
	"fmt"
	"log"
	"os"

	"github.com/getshway/shway/domain"
	"github.com/spf13/cobra"
)

var (
	// RootCmd sets task command config
	RootCmd = &cobra.Command{
		Use: "shway",
		Run: func(cmd *cobra.Command, args []string) {
			versionCommand(cmd, args)
			cmd.Usage()
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  "Show version",
		Run:   versionCommand,
	}
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

func versionCommand(cmd *cobra.Command, args []string) {
	fmt.Println(domain.Version)
}

// Exit finishs requests
func Exit(err error, codes ...int) {
	log.SetOutput(os.Stderr)
	log.SetPrefix("shway: ")
	log.SetFlags(0)

	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 1
	}
	log.Print(err.Error())
	os.Exit(code)
}

// Run runs CLI action
func Run() {
	if err := RootCmd.Execute(); err != nil {
		Exit(fmt.Errorf("failed to run: %s", err.Error()), 1)
	}
}
