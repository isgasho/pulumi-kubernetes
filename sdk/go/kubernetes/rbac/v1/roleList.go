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

// RoleList is a collection of Roles
type RoleList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of Roles
	Items RoleTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewRoleList registers a new resource with the given unique name, arguments, and options.
func NewRoleList(ctx *pulumi.Context,
	name string, args *RoleListArgs, opts ...pulumi.ResourceOption) (*RoleList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("RoleList")
	var resource RoleList
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1:RoleList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleList gets an existing RoleList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleListState, opts ...pulumi.ResourceOption) (*RoleList, error) {
	var resource RoleList
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1:RoleList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleList resources.
type roleListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of Roles
	Items []RoleType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type RoleListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of Roles
	Items RoleTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrInput
}

func (RoleListState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleListState)(nil)).Elem()
}

type roleListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of Roles
	Items []RoleType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a RoleList resource.
type RoleListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of Roles
	Items RoleTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrInput
}

func (RoleListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleListArgs)(nil)).Elem()
}

type RoleListInput interface {
	pulumi.Input

	ToRoleListOutput() RoleListOutput
	ToRoleListOutputWithContext(ctx context.Context) RoleListOutput
}

func (RoleList) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleList)(nil)).Elem()
}

func (i RoleList) ToRoleListOutput() RoleListOutput {
	return i.ToRoleListOutputWithContext(context.Background())
}

func (i RoleList) ToRoleListOutputWithContext(ctx context.Context) RoleListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleListOutput)
}

type RoleListOutput struct {
	*pulumi.OutputState
}

func (RoleListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleListOutput)(nil)).Elem()
}

func (o RoleListOutput) ToRoleListOutput() RoleListOutput {
	return o
}

func (o RoleListOutput) ToRoleListOutputWithContext(ctx context.Context) RoleListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoleListOutput{})
}
