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
	RootCmd.AddCommand(addCmd)
}

var (
	role string
)

var addCmd = &cobra.Command{
	Use:     "add [user] [namespace]",
	Aliases: []string{"command"},
	Short:   "Add user to namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if err := add(cmd, args); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&role, "role", "r", "view", "")
}

func add(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("Expected two arguments")
	}

	user := args[0]
	namespace := args[1]

	if !contains(roles, role) {
		return fmt.Errorf("Invalid role %s", role)
	}

	return applyRoleBinding(clientset, namespace, v1alpha1.RoleBinding{
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("%s-%s", namespace, role),
		},
		RoleRef: v1alpha1.RoleRef{
			Kind: "ClusterRole",
			Name: role,
		},
	}, user, "")
}
