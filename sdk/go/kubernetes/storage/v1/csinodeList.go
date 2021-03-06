// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CSINodeList is a collection of CSINode objects.
type CSINodeList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is the list of CSINode
	Items CSINodeTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCSINodeList registers a new resource with the given unique name, arguments, and options.
func NewCSINodeList(ctx *pulumi.Context,
	name string, args *CSINodeListArgs, opts ...pulumi.ResourceOption) (*CSINodeList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSINodeList")
	var resource CSINodeList
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSINodeList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSINodeList gets an existing CSINodeList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSINodeList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSINodeListState, opts ...pulumi.ResourceOption) (*CSINodeList, error) {
	var resource CSINodeList
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSINodeList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSINodeList resources.
type csinodeListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CSINode
	Items []CSINodeType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type CSINodeListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CSINode
	Items CSINodeTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CSINodeListState) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeListState)(nil)).Elem()
}

type csinodeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CSINode
	Items []CSINodeType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CSINodeList resource.
type CSINodeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CSINode
	Items CSINodeTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CSINodeListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeListArgs)(nil)).Elem()
}

type CSINodeListInput interface {
	pulumi.Input

	ToCSINodeListOutput() CSINodeListOutput
	ToCSINodeListOutputWithContext(ctx context.Context) CSINodeListOutput
}

func (CSINodeList) ElementType() reflect.Type {
	return reflect.TypeOf((*CSINodeList)(nil)).Elem()
}

func (i CSINodeList) ToCSINodeListOutput() CSINodeListOutput {
	return i.ToCSINodeListOutputWithContext(context.Background())
}

func (i CSINodeList) ToCSINodeListOutputWithContext(ctx context.Context) CSINodeListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodeListOutput)
}

type CSINodeListOutput struct {
	*pulumi.OutputState
}

func (CSINodeListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CSINodeListOutput)(nil)).Elem()
}

func (o CSINodeListOutput) ToCSINodeListOutput() CSINodeListOutput {
	return o
}

func (o CSINodeListOutput) ToCSINodeListOutputWithContext(ctx context.Context) CSINodeListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CSINodeListOutput{})
}
