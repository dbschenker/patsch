package kube

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-version"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type client struct {
	clientset kubernetes.Interface
}

func FindIngresses(kubeconfig string) []string {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return getIngresses(client{clientset: clientset})
}

func getIngresses(c client) []string {
	var ret []string
	serverVersion, err := c.clientset.Discovery().ServerVersion()
	if err != nil {
		panic(err.Error())
	}
	apiVersion := version.Must(version.NewVersion(serverVersion.String()))
	hasExt, err := version.NewConstraint("< 1.22")
	if err != nil {
		panic(err.Error())
	}
	if hasExt.Check(apiVersion) {
		ingressList, err := c.clientset.ExtensionsV1beta1().Ingresses(v1.NamespaceAll).List(context.TODO(), v12.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		ingresses := ingressList.Items
		if len(ingresses) > 0 {
			for _, ingress := range ingresses {
				for _, rule := range ingress.Spec.Rules {
					if rule.Host != "" {
						for _, p := range rule.HTTP.Paths {
							ret = append(ret, fmt.Sprintf("https://%s%s", rule.Host, p.Path))
						}
					}
				}
			}
		}
	}

	ingressList, err := c.clientset.NetworkingV1().Ingresses(v1.NamespaceAll).List(context.TODO(), v12.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ingresses := ingressList.Items
	if len(ingresses) > 0 {
		for _, ingress := range ingresses {
			for _, rule := range ingress.Spec.Rules {
				if rule.Host != "" && rule.HTTP != nil {
					for _, p := range rule.HTTP.Paths {
						ret = append(ret, fmt.Sprintf("https://%s%s", rule.Host, p.Path))
					}
				}
			}
		}
	}
	return ret
}
