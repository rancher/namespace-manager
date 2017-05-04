package cmd

import (
	"os/user"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig string
	clientset  *kubernetes.Clientset
)

var RootCmd = &cobra.Command{
	Use: "auth-manager",
}

func init() {
	cobra.OnInitialize(func() {
		if err := initClient(); err != nil {
			log.Fatal(err)
		}
	})
	viper.SetEnvPrefix("namespace_manager")
	viper.AutomaticEnv()
	RootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "")
}

func initClient() error {
	if kubeconfig == "" {
		user, err := user.Current()
		if err != nil {
			return err
		}
		kubeconfig = path.Join(user.HomeDir, ".kube", "config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}
