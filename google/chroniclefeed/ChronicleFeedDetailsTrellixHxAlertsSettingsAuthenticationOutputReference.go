// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference interface {
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
	InternalValue() *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthentication
	SetInternalValue(val *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthentication)
	Msso() ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMssoOutputReference
	MssoInput() *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMsso
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrellixIam() ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIamOutputReference
	TrellixIamInput() *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIam
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
	PutMsso(value *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMsso)
	PutTrellixIam(value *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIam)
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

// The jsii proxy struct for ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference
type jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) InternalValue() *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthentication {
	var returns *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) Msso() ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMssoOutputReference {
	var returns ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMssoOutputReference
	_jsii_.Get(
		j,
		"msso",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) MssoInput() *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMsso {
	var returns *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMsso
	_jsii_.Get(
		j,
		"mssoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) TrellixIam() ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIamOutputReference {
	var returns ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIamOutputReference
	_jsii_.Get(
		j,
		"trellixIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) TrellixIamInput() *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIam {
	var returns *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIam
	_jsii_.Get(
		j,
		"trellixIamInput",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference_Override(c ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference)SetInternalValue(val *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) PutMsso(value *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationMsso) {
	if err := c.validatePutMssoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMsso",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) PutTrellixIam(value *ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationTrellixIam) {
	if err := c.validatePutTrellixIamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrellixIam",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) ResetMsso() {
	_jsii_.InvokeVoid(
		c,
		"resetMsso",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) ResetTrellixIam() {
	_jsii_.InvokeVoid(
		c,
		"resetTrellixIam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsTrellixHxAlertsSettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

