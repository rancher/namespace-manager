package cmd

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:     "create [namespace]",
	Aliases: []string{"command"},
	Short:   "Create namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if err := create(cmd, args); err != nil {
			log.Fatal(err)
		}
	},
}

func create(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Expected one argument")
	}
	namespace := args[0]
	return applyNamespace(clientset, namespace)
}
