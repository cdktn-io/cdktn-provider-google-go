// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthentication
	SetInternalValue(val *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthentication)
	Msso() ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMssoOutputReference
	MssoInput() *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMsso
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrellixIam() ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIamOutputReference
	TrellixIamInput() *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIam
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
	PutMsso(value *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMsso)
	PutTrellixIam(value *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIam)
	ResetMsso()
	ResetTrellixIam()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference
type jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) InternalValue() *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthentication {
	var returns *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) Msso() ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMssoOutputReference {
	var returns ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMssoOutputReference
	_jsii_.Get(
		j,
		"msso",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) MssoInput() *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMsso {
	var returns *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMsso
	_jsii_.Get(
		j,
		"mssoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) TrellixIam() ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIamOutputReference {
	var returns ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIamOutputReference
	_jsii_.Get(
		j,
		"trellixIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) TrellixIamInput() *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIam {
	var returns *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIam
	_jsii_.Get(
		j,
		"trellixIamInput",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference_Override(c ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference)SetInternalValue(val *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) PutMsso(value *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMsso) {
	if err := c.validatePutMssoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMsso",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) PutTrellixIam(value *ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIam) {
	if err := c.validatePutTrellixIamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrellixIam",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) ResetMsso() {
	_jsii_.InvokeVoid(
		c,
		"resetMsso",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) ResetTrellixIam() {
	_jsii_.InvokeVoid(
		c,
		"resetTrellixIam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

