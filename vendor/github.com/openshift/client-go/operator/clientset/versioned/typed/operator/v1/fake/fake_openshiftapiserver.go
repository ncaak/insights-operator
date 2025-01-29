// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	typedoperatorv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeOpenShiftAPIServers implements OpenShiftAPIServerInterface
type fakeOpenShiftAPIServers struct {
	*gentype.FakeClientWithListAndApply[*v1.OpenShiftAPIServer, *v1.OpenShiftAPIServerList, *operatorv1.OpenShiftAPIServerApplyConfiguration]
	Fake *FakeOperatorV1
}

func newFakeOpenShiftAPIServers(fake *FakeOperatorV1) typedoperatorv1.OpenShiftAPIServerInterface {
	return &fakeOpenShiftAPIServers{
		gentype.NewFakeClientWithListAndApply[*v1.OpenShiftAPIServer, *v1.OpenShiftAPIServerList, *operatorv1.OpenShiftAPIServerApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("openshiftapiservers"),
			v1.SchemeGroupVersion.WithKind("OpenShiftAPIServer"),
			func() *v1.OpenShiftAPIServer { return &v1.OpenShiftAPIServer{} },
			func() *v1.OpenShiftAPIServerList { return &v1.OpenShiftAPIServerList{} },
			func(dst, src *v1.OpenShiftAPIServerList) { dst.ListMeta = src.ListMeta },
			func(list *v1.OpenShiftAPIServerList) []*v1.OpenShiftAPIServer {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.OpenShiftAPIServerList, items []*v1.OpenShiftAPIServer) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
