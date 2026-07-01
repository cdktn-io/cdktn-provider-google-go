// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregionresizerequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/computeregionresizerequest/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference interface {
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
	ErrorInfo() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsErrorInfoList
	// Experimental.
	Fqn() *string
	Help() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsHelpList
	InternalValue() *ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetails
	SetInternalValue(val *ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetails)
	LocalizedMessage() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsLocalizedMessageList
	QuotaInfo() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference
type jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ErrorInfo() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsErrorInfoList {
	var returns ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsErrorInfoList
	_jsii_.Get(
		j,
		"errorInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) Help() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsHelpList {
	var returns ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsHelpList
	_jsii_.Get(
		j,
		"help",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) InternalValue() *ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetails {
	var returns *ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) LocalizedMessage() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsLocalizedMessageList {
	var returns ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsLocalizedMessageList
	_jsii_.Get(
		j,
		"localizedMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) QuotaInfo() ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoList {
	var returns ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoList
	_jsii_.Get(
		j,
		"quotaInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeRegionResizeRequest.ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference_Override(c ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeRegionResizeRequest.ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetInternalValue(val *ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

