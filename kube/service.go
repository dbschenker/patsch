package kube

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const externalDnsKey = "external-dns.alpha.kubernetes.io/hostname"

func FindServices(kubeconfig string) []string {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return getServices(client{clientset: clientset})
}

func getServices(c client) []string {
	var ret []string
	serviceList, err := c.clientset.CoreV1().Services(v1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	services := serviceList.Items
	if len(services) > 0 {
		for _, services := range services {
			if services.Spec.Type != v1.ServiceTypeLoadBalancer {
				continue
			}

			a, ok := services.Annotations[externalDnsKey]
			if !ok {
				continue
			}
			ret = append(ret, fmt.Sprintf("https://%s", a))
		}
	}
	return ret
}
