// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/datalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference interface {
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
	Dictionary() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	DictionaryInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary
	ExcludeByHotword() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	ExcludeByHotwordInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	ExcludeByImageFindings() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsOutputReference
	ExcludeByImageFindingsInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings
	ExcludeInfoTypes() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	ExcludeInfoTypesInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	// Experimental.
	Fqn() *string
	InternalValue() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule
	SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule)
	MatchingType() *string
	SetMatchingType(val *string)
	MatchingTypeInput() *string
	Regex() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	RegexInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex
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
	PutDictionary(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary)
	PutExcludeByHotword(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword)
	PutExcludeByImageFindings(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings)
	PutExcludeInfoTypes(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	PutRegex(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex)
	ResetDictionary()
	ResetExcludeByHotword()
	ResetExcludeByImageFindings()
	ResetExcludeInfoTypes()
	ResetRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference
type jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Dictionary() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) DictionaryInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotword() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	_jsii_.Get(
		j,
		"excludeByHotword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotwordInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	_jsii_.Get(
		j,
		"excludeByHotwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByImageFindings() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsOutputReference
	_jsii_.Get(
		j,
		"excludeByImageFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByImageFindingsInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings
	_jsii_.Get(
		j,
		"excludeByImageFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypes() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	_jsii_.Get(
		j,
		"excludeInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypesInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	_jsii_.Get(
		j,
		"excludeInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) InternalValue() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Regex() DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegexOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) RegexInput() *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex {
	var returns *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference_Override(d DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetInternalValue(val *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetMatchingType(val *string) {
	if err := j.validateSetMatchingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingType",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutDictionary(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary) {
	if err := d.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDictionary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeByHotword(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword) {
	if err := d.validatePutExcludeByHotwordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeByHotword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeByImageFindings(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings) {
	if err := d.validatePutExcludeByImageFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeByImageFindings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeInfoTypes(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) {
	if err := d.validatePutExcludeInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeInfoTypes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutRegex(value *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex) {
	if err := d.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRegex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		d,
		"resetDictionary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeByHotword() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeByHotword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeByImageFindings() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeByImageFindings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

