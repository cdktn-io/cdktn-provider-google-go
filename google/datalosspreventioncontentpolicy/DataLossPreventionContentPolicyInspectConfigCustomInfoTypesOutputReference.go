// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/datalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference interface {
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
	DetectionRules() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesList
	DetectionRulesInput() interface{}
	Dictionary() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionaryOutputReference
	DictionaryInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary
	ExclusionType() *string
	SetExclusionType(val *string)
	ExclusionTypeInput() *string
	FileLabelInfoType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeOutputReference
	FileLabelInfoTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType
	// Experimental.
	Fqn() *string
	InfoType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoTypeOutputReference
	InfoTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Likelihood() *string
	SetLikelihood(val *string)
	LikelihoodInput() *string
	MetadataKeyValueExpression() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpressionOutputReference
	MetadataKeyValueExpressionInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression
	Regex() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegexOutputReference
	RegexInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex
	SensitivityScore() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScoreOutputReference
	SensitivityScoreInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore
	StoredType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredTypeOutputReference
	StoredTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType
	SurrogateType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	SurrogateTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType
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
	PutDetectionRules(value interface{})
	PutDictionary(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary)
	PutFileLabelInfoType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType)
	PutInfoType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType)
	PutMetadataKeyValueExpression(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression)
	PutRegex(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex)
	PutSensitivityScore(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore)
	PutStoredType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType)
	PutSurrogateType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType)
	ResetDetectionRules()
	ResetDictionary()
	ResetExclusionType()
	ResetFileLabelInfoType()
	ResetLikelihood()
	ResetMetadataKeyValueExpression()
	ResetRegex()
	ResetSensitivityScore()
	ResetStoredType()
	ResetSurrogateType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference
type jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) DetectionRules() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesList {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesList
	_jsii_.Get(
		j,
		"detectionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) DetectionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Dictionary() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionaryOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) DictionaryInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ExclusionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ExclusionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) FileLabelInfoType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeOutputReference
	_jsii_.Get(
		j,
		"fileLabelInfoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) FileLabelInfoTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType
	_jsii_.Get(
		j,
		"fileLabelInfoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InfoType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoTypeOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoTypeOutputReference
	_jsii_.Get(
		j,
		"infoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InfoTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType
	_jsii_.Get(
		j,
		"infoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Likelihood() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) LikelihoodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) MetadataKeyValueExpression() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpressionOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpressionOutputReference
	_jsii_.Get(
		j,
		"metadataKeyValueExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) MetadataKeyValueExpressionInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression
	_jsii_.Get(
		j,
		"metadataKeyValueExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Regex() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegexOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) RegexInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SensitivityScore() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScoreOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScoreOutputReference
	_jsii_.Get(
		j,
		"sensitivityScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SensitivityScoreInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore
	_jsii_.Get(
		j,
		"sensitivityScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) StoredType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredTypeOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredTypeOutputReference
	_jsii_.Get(
		j,
		"storedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) StoredTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType
	_jsii_.Get(
		j,
		"storedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SurrogateType() DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateTypeOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	_jsii_.Get(
		j,
		"surrogateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SurrogateTypeInput() *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType {
	var returns *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType
	_jsii_.Get(
		j,
		"surrogateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference_Override(d DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetExclusionType(val *string) {
	if err := j.validateSetExclusionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionType",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetLikelihood(val *string) {
	if err := j.validateSetLikelihoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likelihood",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutDetectionRules(value interface{}) {
	if err := d.validatePutDetectionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDetectionRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutDictionary(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary) {
	if err := d.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDictionary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutFileLabelInfoType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType) {
	if err := d.validatePutFileLabelInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFileLabelInfoType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutInfoType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType) {
	if err := d.validatePutInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInfoType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutMetadataKeyValueExpression(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression) {
	if err := d.validatePutMetadataKeyValueExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetadataKeyValueExpression",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutRegex(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex) {
	if err := d.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRegex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutSensitivityScore(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore) {
	if err := d.validatePutSensitivityScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSensitivityScore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutStoredType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType) {
	if err := d.validatePutStoredTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStoredType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutSurrogateType(value *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType) {
	if err := d.validatePutSurrogateTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSurrogateType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetDetectionRules() {
	_jsii_.InvokeVoid(
		d,
		"resetDetectionRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		d,
		"resetDictionary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetExclusionType() {
	_jsii_.InvokeVoid(
		d,
		"resetExclusionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetFileLabelInfoType() {
	_jsii_.InvokeVoid(
		d,
		"resetFileLabelInfoType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetLikelihood() {
	_jsii_.InvokeVoid(
		d,
		"resetLikelihood",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetMetadataKeyValueExpression() {
	_jsii_.InvokeVoid(
		d,
		"resetMetadataKeyValueExpression",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetSensitivityScore() {
	_jsii_.InvokeVoid(
		d,
		"resetSensitivityScore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetStoredType() {
	_jsii_.InvokeVoid(
		d,
		"resetStoredType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetSurrogateType() {
	_jsii_.InvokeVoid(
		d,
		"resetSurrogateType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

