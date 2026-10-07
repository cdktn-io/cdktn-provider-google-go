// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/datalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference interface {
	cdktn.ComplexObject
	AdjustmentRule() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference
	AdjustmentRuleInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule
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
	ExclusionRule() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference
	ExclusionRuleInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule
	// Experimental.
	Fqn() *string
	HotwordRule() DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRuleOutputReference
	HotwordRuleInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutAdjustmentRule(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule)
	PutExclusionRule(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule)
	PutHotwordRule(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule)
	ResetAdjustmentRule()
	ResetExclusionRule()
	ResetHotwordRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference
type jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) AdjustmentRule() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference
	_jsii_.Get(
		j,
		"adjustmentRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) AdjustmentRuleInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule
	_jsii_.Get(
		j,
		"adjustmentRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ExclusionRule() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference
	_jsii_.Get(
		j,
		"exclusionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ExclusionRuleInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule
	_jsii_.Get(
		j,
		"exclusionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) HotwordRule() DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRuleOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRuleOutputReference
	_jsii_.Get(
		j,
		"hotwordRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) HotwordRuleInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule
	_jsii_.Get(
		j,
		"hotwordRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference_Override(d DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) PutAdjustmentRule(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule) {
	if err := d.validatePutAdjustmentRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdjustmentRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) PutExclusionRule(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule) {
	if err := d.validatePutExclusionRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExclusionRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) PutHotwordRule(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule) {
	if err := d.validatePutHotwordRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHotwordRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ResetAdjustmentRule() {
	_jsii_.InvokeVoid(
		d,
		"resetAdjustmentRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ResetExclusionRule() {
	_jsii_.InvokeVoid(
		d,
		"resetExclusionRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ResetHotwordRule() {
	_jsii_.InvokeVoid(
		d,
		"resetHotwordRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

