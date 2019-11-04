package command

import (
	"fmt"

	"github.com/getshway/shway/project"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list project names",
		Long:  "list",
		Run:   listCommand,
	}
	loadCmd = &cobra.Command{
		Use:   "load",
		Short: "load project",
		Long:  "load project",
		Run:   loadCommand,
	}
	setCmd = &cobra.Command{
		Use:   "set",
		Short: "set new project",
		Long:  "set new project",
		Run:   setCommand,
	}
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "update a project",
		Long:  "update a project",
		Run:   updateCommand,
	}
)

func init() {
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(loadCmd)
	RootCmd.AddCommand(setCmd)
	RootCmd.AddCommand(updateCmd)
}

func listCommand(cmd *cobra.Command, args []string) {
	ps, err := project.List()
	if err != nil {
		Exit(err, 1)
	}
	for _, p := range ps {
		fmt.Println(p)
	}
}

func loadCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		Exit(fmt.Errorf("parameters are not enough: shway project load (project name)"), 1)
	}
	out, err := project.Load(args[0])
	if err != nil {
		Exit(err, 1)
	}
	fmt.Println(out)
}

func setCommand(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		Exit(fmt.Errorf("parameters are not enough: shway project set (project name) (project git repo URI)"), 1)
	}
	if err := project.Set(args[0], args[1]); err != nil {
		Exit(err, 1)
	}
}

func updateCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		Exit(fmt.Errorf("parameters are not enough: shway project update (project name)"), 1)
	}
	if err := project.Update(args[0]); err != nil {
		Exit(err, 1)
	}
}
