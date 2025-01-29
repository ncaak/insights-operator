// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	networkv1alpha1 "github.com/openshift/api/network/v1alpha1"
	applyconfigurationsnetworkv1alpha1 "github.com/openshift/client-go/network/applyconfigurations/network/v1alpha1"
	scheme "github.com/openshift/client-go/network/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// DNSNameResolversGetter has a method to return a DNSNameResolverInterface.
// A group's client should implement this interface.
type DNSNameResolversGetter interface {
	DNSNameResolvers(namespace string) DNSNameResolverInterface
}

// DNSNameResolverInterface has methods to work with DNSNameResolver resources.
type DNSNameResolverInterface interface {
	Create(ctx context.Context, dNSNameResolver *networkv1alpha1.DNSNameResolver, opts v1.CreateOptions) (*networkv1alpha1.DNSNameResolver, error)
	Update(ctx context.Context, dNSNameResolver *networkv1alpha1.DNSNameResolver, opts v1.UpdateOptions) (*networkv1alpha1.DNSNameResolver, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, dNSNameResolver *networkv1alpha1.DNSNameResolver, opts v1.UpdateOptions) (*networkv1alpha1.DNSNameResolver, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*networkv1alpha1.DNSNameResolver, error)
	List(ctx context.Context, opts v1.ListOptions) (*networkv1alpha1.DNSNameResolverList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkv1alpha1.DNSNameResolver, err error)
	Apply(ctx context.Context, dNSNameResolver *applyconfigurationsnetworkv1alpha1.DNSNameResolverApplyConfiguration, opts v1.ApplyOptions) (result *networkv1alpha1.DNSNameResolver, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, dNSNameResolver *applyconfigurationsnetworkv1alpha1.DNSNameResolverApplyConfiguration, opts v1.ApplyOptions) (result *networkv1alpha1.DNSNameResolver, err error)
	DNSNameResolverExpansion
}

// dNSNameResolvers implements DNSNameResolverInterface
type dNSNameResolvers struct {
	*gentype.ClientWithListAndApply[*networkv1alpha1.DNSNameResolver, *networkv1alpha1.DNSNameResolverList, *applyconfigurationsnetworkv1alpha1.DNSNameResolverApplyConfiguration]
}

// newDNSNameResolvers returns a DNSNameResolvers
func newDNSNameResolvers(c *NetworkV1alpha1Client, namespace string) *dNSNameResolvers {
	return &dNSNameResolvers{
		gentype.NewClientWithListAndApply[*networkv1alpha1.DNSNameResolver, *networkv1alpha1.DNSNameResolverList, *applyconfigurationsnetworkv1alpha1.DNSNameResolverApplyConfiguration](
			"dnsnameresolvers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *networkv1alpha1.DNSNameResolver { return &networkv1alpha1.DNSNameResolver{} },
			func() *networkv1alpha1.DNSNameResolverList { return &networkv1alpha1.DNSNameResolverList{} },
		),
	}
}
