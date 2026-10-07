// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/datalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference interface {
	cdktn.ComplexObject
	AdjustByImageFindings() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsOutputReference
	AdjustByImageFindingsInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings
	AdjustByMatchingInfoTypes() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypesOutputReference
	AdjustByMatchingInfoTypesInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes
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
	InternalValue() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule
	SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule)
	LikelihoodAdjustment() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustmentOutputReference
	LikelihoodAdjustmentInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment
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
	PutAdjustByImageFindings(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings)
	PutAdjustByMatchingInfoTypes(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes)
	PutLikelihoodAdjustment(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment)
	ResetAdjustByImageFindings()
	ResetAdjustByMatchingInfoTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference
type jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByImageFindings() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsOutputReference
	_jsii_.Get(
		j,
		"adjustByImageFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByImageFindingsInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings
	_jsii_.Get(
		j,
		"adjustByImageFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByMatchingInfoTypes() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypesOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypesOutputReference
	_jsii_.Get(
		j,
		"adjustByMatchingInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByMatchingInfoTypesInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes
	_jsii_.Get(
		j,
		"adjustByMatchingInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) InternalValue() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) LikelihoodAdjustment() DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustmentOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustmentOutputReference
	_jsii_.Get(
		j,
		"likelihoodAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) LikelihoodAdjustmentInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment
	_jsii_.Get(
		j,
		"likelihoodAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference_Override(d DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) PutAdjustByImageFindings(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings) {
	if err := d.validatePutAdjustByImageFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdjustByImageFindings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) PutAdjustByMatchingInfoTypes(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes) {
	if err := d.validatePutAdjustByMatchingInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdjustByMatchingInfoTypes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) PutLikelihoodAdjustment(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment) {
	if err := d.validatePutLikelihoodAdjustmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLikelihoodAdjustment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ResetAdjustByImageFindings() {
	_jsii_.InvokeVoid(
		d,
		"resetAdjustByImageFindings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ResetAdjustByMatchingInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetAdjustByMatchingInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

