package cmd

import (
	"errors"
	"fmt"

	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/rbac/v1alpha1"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:     "remove [user] [namespace]",
	Aliases: []string{"command"},
	Short:   "Remove user from namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if err := remove(cmd, args); err != nil {
			log.Fatal(err)
		}
	},
}

func remove(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("Expected two arguments")
	}

	user := args[0]
	namespace := args[1]

	for _, role := range roles {
		if err := applyRoleBinding(clientset, namespace, v1alpha1.RoleBinding{
			ObjectMeta: v1.ObjectMeta{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s-%s", namespace, role),
			},
			RoleRef: v1alpha1.RoleRef{
				Kind: "ClusterRole",
				Name: role,
			},
		}, "", user); err != nil {
			return err
		}
	}

	return nil
}
