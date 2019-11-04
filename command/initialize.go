package command

import (
	"fmt"

	"github.com/getshway/shway/initialize"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "setup shway and load zsh config",
		Long:  "setup shway and load zsh config",
		Run:   initCommand,
	}
)

func init() {
	RootCmd.AddCommand(initCmd)
}

func initCommand(cmd *cobra.Command, args []string) {
	out, err := initialize.Run()
	if err != nil {
		Exit(err, 1)
	}
	fmt.Println(out)
}
