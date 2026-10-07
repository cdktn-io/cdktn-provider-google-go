// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/datalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference interface {
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
	HotwordRegex() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexOutputReference
	HotwordRegexInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex
	InternalValue() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRule
	SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRule)
	LikelihoodAdjustment() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentOutputReference
	LikelihoodAdjustmentInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment
	Proximity() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference
	ProximityInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
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
	PutHotwordRegex(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex)
	PutLikelihoodAdjustment(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment)
	PutProximity(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference
type jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) HotwordRegex() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexOutputReference
	_jsii_.Get(
		j,
		"hotwordRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) HotwordRegexInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex
	_jsii_.Get(
		j,
		"hotwordRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) InternalValue() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) LikelihoodAdjustment() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentOutputReference
	_jsii_.Get(
		j,
		"likelihoodAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) LikelihoodAdjustmentInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment
	_jsii_.Get(
		j,
		"likelihoodAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) Proximity() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference
	_jsii_.Get(
		j,
		"proximity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) ProximityInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	_jsii_.Get(
		j,
		"proximityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference_Override(d DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference)SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) PutHotwordRegex(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) {
	if err := d.validatePutHotwordRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHotwordRegex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) PutLikelihoodAdjustment(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) {
	if err := d.validatePutLikelihoodAdjustmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLikelihoodAdjustment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) PutProximity(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) {
	if err := d.validatePutProximityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProximity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

