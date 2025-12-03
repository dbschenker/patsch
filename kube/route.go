package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	gwclient "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

func FindHTTPRoutes(kubeconfig string) []string {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the Gateway API clientset
	clientset, err := gwclient.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return getHTTPRoutes(clientset)
}

func getHTTPRoutes(c gwclient.Interface) []string {
	var ret []string

	httpRouteList, err := c.GatewayV1().HTTPRoutes(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, route := range httpRouteList.Items {
		// iterate hostnames and rules
		for _, host := range route.Spec.Hostnames {
			for _, rule := range route.Spec.Rules {
				for _, match := range rule.Matches {
					if match.Path != nil && *match.Path.Value != "" {
						ret = append(ret, fmt.Sprintf("https://%s%s", string(host), *match.Path.Value))
					}
				}
			}
		}
	}

	return ret
}
