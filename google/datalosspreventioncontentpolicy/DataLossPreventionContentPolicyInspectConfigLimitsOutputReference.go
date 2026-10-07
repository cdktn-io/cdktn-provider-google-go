// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/datalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLossPreventionContentPolicyInspectConfigLimitsOutputReference interface {
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
	InternalValue() *DataLossPreventionContentPolicyInspectConfigLimits
	SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigLimits)
	MaxFindingsPerInfoType() DataLossPreventionContentPolicyInspectConfigLimitsMaxFindingsPerInfoTypeList
	MaxFindingsPerInfoTypeInput() interface{}
	MaxFindingsPerItem() *float64
	SetMaxFindingsPerItem(val *float64)
	MaxFindingsPerItemInput() *float64
	MaxFindingsPerRequest() *float64
	SetMaxFindingsPerRequest(val *float64)
	MaxFindingsPerRequestInput() *float64
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
	PutMaxFindingsPerInfoType(value interface{})
	ResetMaxFindingsPerInfoType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionContentPolicyInspectConfigLimitsOutputReference
type jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) InternalValue() *DataLossPreventionContentPolicyInspectConfigLimits {
	var returns *DataLossPreventionContentPolicyInspectConfigLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) MaxFindingsPerInfoType() DataLossPreventionContentPolicyInspectConfigLimitsMaxFindingsPerInfoTypeList {
	var returns DataLossPreventionContentPolicyInspectConfigLimitsMaxFindingsPerInfoTypeList
	_jsii_.Get(
		j,
		"maxFindingsPerInfoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) MaxFindingsPerInfoTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxFindingsPerInfoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) MaxFindingsPerItem() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) MaxFindingsPerItemInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerItemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) MaxFindingsPerRequest() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) MaxFindingsPerRequestInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionContentPolicyInspectConfigLimitsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataLossPreventionContentPolicyInspectConfigLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionContentPolicyInspectConfigLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionContentPolicyInspectConfigLimitsOutputReference_Override(d DataLossPreventionContentPolicyInspectConfigLimitsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference)SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference)SetMaxFindingsPerItem(val *float64) {
	if err := j.validateSetMaxFindingsPerItemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFindingsPerItem",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference)SetMaxFindingsPerRequest(val *float64) {
	if err := j.validateSetMaxFindingsPerRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFindingsPerRequest",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) PutMaxFindingsPerInfoType(value interface{}) {
	if err := d.validatePutMaxFindingsPerInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMaxFindingsPerInfoType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) ResetMaxFindingsPerInfoType() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFindingsPerInfoType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

