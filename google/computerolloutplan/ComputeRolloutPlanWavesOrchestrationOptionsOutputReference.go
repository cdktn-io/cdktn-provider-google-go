// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/computerolloutplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeRolloutPlanWavesOrchestrationOptionsOutputReference interface {
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
	Delays() ComputeRolloutPlanWavesOrchestrationOptionsDelaysList
	DelaysInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRolloutPlanWavesOrchestrationOptions
	SetInternalValue(val *ComputeRolloutPlanWavesOrchestrationOptions)
	MaxConcurrentLocations() *float64
	SetMaxConcurrentLocations(val *float64)
	MaxConcurrentLocationsInput() *float64
	MaxConcurrentResourcesPerLocation() *float64
	SetMaxConcurrentResourcesPerLocation(val *float64)
	MaxConcurrentResourcesPerLocationInput() *float64
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
	PutDelays(value interface{})
	ResetDelays()
	ResetMaxConcurrentLocations()
	ResetMaxConcurrentResourcesPerLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRolloutPlanWavesOrchestrationOptionsOutputReference
type jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) Delays() ComputeRolloutPlanWavesOrchestrationOptionsDelaysList {
	var returns ComputeRolloutPlanWavesOrchestrationOptionsDelaysList
	_jsii_.Get(
		j,
		"delays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) DelaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) InternalValue() *ComputeRolloutPlanWavesOrchestrationOptions {
	var returns *ComputeRolloutPlanWavesOrchestrationOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) MaxConcurrentLocations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) MaxConcurrentLocationsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) MaxConcurrentResourcesPerLocation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentResourcesPerLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) MaxConcurrentResourcesPerLocationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentResourcesPerLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRolloutPlanWavesOrchestrationOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeRolloutPlanWavesOrchestrationOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRolloutPlanWavesOrchestrationOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeRolloutPlan.ComputeRolloutPlanWavesOrchestrationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRolloutPlanWavesOrchestrationOptionsOutputReference_Override(c ComputeRolloutPlanWavesOrchestrationOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeRolloutPlan.ComputeRolloutPlanWavesOrchestrationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference)SetInternalValue(val *ComputeRolloutPlanWavesOrchestrationOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference)SetMaxConcurrentLocations(val *float64) {
	if err := j.validateSetMaxConcurrentLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentLocations",
		val,
	)
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference)SetMaxConcurrentResourcesPerLocation(val *float64) {
	if err := j.validateSetMaxConcurrentResourcesPerLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentResourcesPerLocation",
		val,
	)
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) PutDelays(value interface{}) {
	if err := c.validatePutDelaysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDelays",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) ResetDelays() {
	_jsii_.InvokeVoid(
		c,
		"resetDelays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) ResetMaxConcurrentLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConcurrentLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) ResetMaxConcurrentResourcesPerLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConcurrentResourcesPerLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRolloutPlanWavesOrchestrationOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

