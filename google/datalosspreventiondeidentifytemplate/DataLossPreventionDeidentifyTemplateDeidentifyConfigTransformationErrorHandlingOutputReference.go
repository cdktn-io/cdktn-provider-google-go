// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondeidentifytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/datalosspreventiondeidentifytemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference interface {
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
	InternalValue() *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandling
	SetInternalValue(val *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandling)
	LeaveUntransformed() DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformedOutputReference
	LeaveUntransformedInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThrowError() DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowErrorOutputReference
	ThrowErrorInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError
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
	PutLeaveUntransformed(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed)
	PutThrowError(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError)
	ResetLeaveUntransformed()
	ResetThrowError()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference
type jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) InternalValue() *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandling {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) LeaveUntransformed() DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformedOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformedOutputReference
	_jsii_.Get(
		j,
		"leaveUntransformed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) LeaveUntransformedInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed
	_jsii_.Get(
		j,
		"leaveUntransformedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ThrowError() DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowErrorOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowErrorOutputReference
	_jsii_.Get(
		j,
		"throwError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ThrowErrorInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError
	_jsii_.Get(
		j,
		"throwErrorInput",
		&returns,
	)
	return returns
}


func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference_Override(d DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference)SetInternalValue(val *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) PutLeaveUntransformed(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingLeaveUntransformed) {
	if err := d.validatePutLeaveUntransformedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLeaveUntransformed",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) PutThrowError(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingThrowError) {
	if err := d.validatePutThrowErrorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putThrowError",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ResetLeaveUntransformed() {
	_jsii_.InvokeVoid(
		d,
		"resetLeaveUntransformed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ResetThrowError() {
	_jsii_.InvokeVoid(
		d,
		"resetThrowError",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigTransformationErrorHandlingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

