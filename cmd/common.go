package cmd

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/rbac/v1alpha1"
)

var (
	roles = []string{
		"admin",
		"edit",
		"view",
	}
)

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func applyNamespace(clientset *kubernetes.Clientset, namespace string) error {
	_, err := clientset.CoreV1().Namespaces().Get(namespace)
	if err == nil {
		return nil
	}
	_, err = clientset.CoreV1().Namespaces().Create(&v1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: namespace,
		},
	})
	return err
}

func applyRoleBinding(clientset *kubernetes.Clientset, namespace string, roleBinding v1alpha1.RoleBinding, addUser, removeUser string) error {
	subject := v1alpha1.Subject{
		Kind: "User",
		Name: addUser,
	}

	existingBinding, err := clientset.RbacV1alpha1().RoleBindings(namespace).Get(roleBinding.ObjectMeta.Name)
	if err == nil {
		var subjects []v1alpha1.Subject
		for _, existingSubject := range existingBinding.Subjects {
			if existingSubject.Kind == "User" && existingSubject.Name == removeUser {
				continue
			}
			subjects = append(subjects, existingSubject)
		}

		if addUser != "" {
			subjects = append(subjects, subject)
		}

		roleBinding.RoleRef = existingBinding.RoleRef
		roleBinding.Subjects = subjects

		_, err = clientset.RbacV1alpha1().RoleBindings(namespace).Update(&roleBinding)
		return err
	}

	if addUser != "" {
		roleBinding.Subjects = []v1alpha1.Subject{subject}
	}

	_, err = clientset.RbacV1alpha1().RoleBindings(namespace).Create(&roleBinding)
	return err
}
