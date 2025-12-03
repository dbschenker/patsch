package kube

import (
	"testing"

	"github.com/dbschenker/patsch/util"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	gwfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"
)

func TestHTTPRoute(t *testing.T) {
	cs := gwfake.NewClientset(
		&gwv1.HTTPRoute{
			TypeMeta: metav1.TypeMeta{
				Kind:       "HTTPRoute",
				APIVersion: "gateway.networking.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "route1",
			},
			Spec: gwv1.HTTPRouteSpec{
				Hostnames: []gwv1.Hostname{"example.com"},
				Rules: []gwv1.HTTPRouteRule{
					{
						Matches: []gwv1.HTTPRouteMatch{
							{
								Path: &gwv1.HTTPPathMatch{
									Type:  util.Addr(gwv1.PathMatchPathPrefix),
									Value: util.Addr("/"),
								},
							},
						},
					},
				},
			},
		},
	)

	got := getHTTPRoutes(cs)
	assert.Equal(t, 1, len(got), "expected one valid HTTPRoute")
	assert.Equal(t, "https://example.com/", got[0])
}
