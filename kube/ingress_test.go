package kube

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/networking/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery/fake"
	testclient "k8s.io/client-go/kubernetes/fake"

	"testing"
)

func TestIngress(t *testing.T) {
	cs := testclient.NewSimpleClientset(
		// first element is completely empty, but getIngresses should handle this gracefully
		&v1.Ingress{},
		// this is a "real" ingres complete with HTTP rule and path
		&v1.Ingress{
			TypeMeta: v12.TypeMeta{
				Kind:       "Ingress",
				APIVersion: "networking.k8s.io/v1",
			},
			ObjectMeta: v12.ObjectMeta{
				Namespace: "default",
				Name:      "ingrid",
			},
			Spec: v1.IngressSpec{
				Rules: []v1.IngressRule{
					{
						Host: "hase.io",
						IngressRuleValue: v1.IngressRuleValue{HTTP: &v1.HTTPIngressRuleValue{
							Paths: []v1.HTTPIngressPath{{
								Path: "/",
							}},
						},
						},
					},
				},
			},
		},
	)
	// Fake server version or version.Must test will panic
	// thanks: https://itnext.io/testing-kubernetes-go-applications-f1f87502b6ef
	cs.Discovery().(*fake.FakeDiscovery).FakedServerVersion = &version.Info{
		Major:      "1",
		Minor:      "28",
		GitVersion: "v1.28.4",
	}
	vi, _ := cs.Discovery().ServerVersion()
	assert.Equal(t, "28", vi.Minor)
	ingresses := getIngresses(client{clientset: cs})
	assert.Equal(t, 1, len(ingresses), "expected one valid ingress and a skipped one")
	assert.Equal(t, "https://hase.io/", ingresses[0])
}
