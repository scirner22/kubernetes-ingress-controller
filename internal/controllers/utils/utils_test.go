package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	netv1 "k8s.io/api/networking/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ingressWithClass(class string) *netv1.Ingress {
	return &netv1.Ingress{
		Spec: netv1.IngressSpec{IngressClassName: &class},
	}
}

func TestMatchesIngressClassName(t *testing.T) {
	for idx, tt := range []struct {
		obj             client.Object
		controllerClass string
		isDefault       bool

		want bool
	}{

		{obj: ingressWithClass(""), isDefault: false, controllerClass: "foo", want: false},
		{obj: ingressWithClass(""), isDefault: false, controllerClass: "kozel", want: false},
		{obj: ingressWithClass(""), isDefault: true, controllerClass: "", want: true},
		{obj: ingressWithClass(""), isDefault: true, controllerClass: "custom", want: true},
		{obj: ingressWithClass(""), isDefault: true, controllerClass: "foo", want: true},
		{obj: ingressWithClass(""), isDefault: true, controllerClass: "killer", want: true},
		{obj: ingressWithClass(""), isDefault: true, controllerClass: "killer", want: true},
		{obj: ingressWithClass("custom"), isDefault: false, controllerClass: "foo", want: false},
		{obj: ingressWithClass("custom"), isDefault: false, controllerClass: "kozel", want: false},
		{obj: ingressWithClass("custom"), isDefault: true, controllerClass: "custom", want: true},
		{obj: ingressWithClass("custom"), isDefault: true, controllerClass: "foo", want: false},
		{obj: ingressWithClass("custom"), isDefault: true, controllerClass: "kozel", want: false},
		{obj: ingressWithClass("foo"), isDefault: false, controllerClass: "foo", want: true},
		{obj: ingressWithClass("foo"), isDefault: true, controllerClass: "foo", want: true},
		{obj: ingressWithClass("kozel"), isDefault: false, controllerClass: "kozel", want: true},
		{obj: ingressWithClass("kozel"), isDefault: true, controllerClass: "kozel", want: true},
	} {
		t.Run(fmt.Sprintf("test case %d", idx), func(t *testing.T) {
			got := MatchesIngressClassName(tt.obj, tt.controllerClass, tt.isDefault)
			require.Equal(t, tt.want, got)
		})
	}
}
