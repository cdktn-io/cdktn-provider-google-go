// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computezonevmextensionpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/computezonevmextensionpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference interface {
	cdktn.ComplexObject
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
	ExtensionName() *string
	SetExtensionName(val *string)
	ExtensionNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PinnedVersion() *string
	SetPinnedVersion(val *string)
	PinnedVersionInput() *string
	StringConfig() *string
	SetStringConfig(val *string)
	StringConfigInput() *string
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
	ResetPinnedVersion()
	ResetStringConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference
type jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ExtensionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ExtensionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) PinnedVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinnedVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) PinnedVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinnedVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) StringConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) StringConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeZoneVmExtensionPolicyExtensionPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeZoneVmExtensionPolicy.ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference_Override(c ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeZoneVmExtensionPolicy.ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetExtensionName(val *string) {
	if err := j.validateSetExtensionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionName",
		val,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetPinnedVersion(val *string) {
	if err := j.validateSetPinnedVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinnedVersion",
		val,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetStringConfig(val *string) {
	if err := j.validateSetStringConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringConfig",
		val,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ResetPinnedVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetPinnedVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ResetStringConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetStringConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeZoneVmExtensionPolicyExtensionPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

