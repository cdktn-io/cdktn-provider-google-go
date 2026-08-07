// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/colabschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference interface {
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
	FailurePolicy() *string
	SetFailurePolicy(val *string)
	FailurePolicyInput() *string
	// Experimental.
	Fqn() *string
	GcsOutputDirectory() *string
	SetGcsOutputDirectory(val *string)
	GcsOutputDirectoryInput() *string
	InternalValue() *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	SetInternalValue(val *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig)
	ParameterValues() *map[string]*string
	SetParameterValues(val *map[string]*string)
	ParameterValuesInput() *map[string]*string
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
	ResetFailurePolicy()
	ResetParameterValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
type jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GcsOutputDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GcsOutputDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InternalValue() *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig {
	var returns *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ParameterValues() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ParameterValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.colabSchedule.ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference_Override(c ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.colabSchedule.ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetGcsOutputDirectory(val *string) {
	if err := j.validateSetGcsOutputDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputDirectory",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetInternalValue(val *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetParameterValues(val *map[string]*string) {
	if err := j.validateSetParameterValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterValues",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ResetParameterValues() {
	_jsii_.InvokeVoid(
		c,
		"resetParameterValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

