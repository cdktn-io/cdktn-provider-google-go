// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/networkserviceshttproute/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkServicesHttpRouteRulesActionDirectResponseOutputReference interface {
	cdktn.ComplexObject
	BytesBody() *string
	SetBytesBody(val *string)
	BytesBodyInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkServicesHttpRouteRulesActionDirectResponse
	SetInternalValue(val *NetworkServicesHttpRouteRulesActionDirectResponse)
	Status() *float64
	SetStatus(val *float64)
	StatusInput() *float64
	StringBody() *string
	SetStringBody(val *string)
	StringBodyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetBytesBody()
	ResetStringBody()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesHttpRouteRulesActionDirectResponseOutputReference
type jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) BytesBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bytesBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) BytesBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bytesBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) InternalValue() *NetworkServicesHttpRouteRulesActionDirectResponse {
	var returns *NetworkServicesHttpRouteRulesActionDirectResponse
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) Status() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) StatusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) StringBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) StringBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesHttpRouteRulesActionDirectResponseOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkServicesHttpRouteRulesActionDirectResponseOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesHttpRouteRulesActionDirectResponseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionDirectResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesHttpRouteRulesActionDirectResponseOutputReference_Override(n NetworkServicesHttpRouteRulesActionDirectResponseOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionDirectResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetBytesBody(val *string) {
	if err := j.validateSetBytesBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesBody",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetInternalValue(val *NetworkServicesHttpRouteRulesActionDirectResponse) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetStatus(val *float64) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetStringBody(val *string) {
	if err := j.validateSetStringBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringBody",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) ResetBytesBody() {
	_jsii_.InvokeVoid(
		n,
		"resetBytesBody",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) ResetStringBody() {
	_jsii_.InvokeVoid(
		n,
		"resetStringBody",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionDirectResponseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

