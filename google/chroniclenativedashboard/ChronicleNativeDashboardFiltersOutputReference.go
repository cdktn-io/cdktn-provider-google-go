// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclenativedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclenativedashboard/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleNativeDashboardFiltersOutputReference interface {
	cdktn.ComplexObject
	ChartIds() *[]*string
	SetChartIds(val *[]*string)
	ChartIdsInput() *[]*string
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
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	FieldPath() *string
	SetFieldPath(val *string)
	FieldPathInput() *string
	FilterOperatorAndFieldValues() ChronicleNativeDashboardFiltersFilterOperatorAndFieldValuesList
	FilterOperatorAndFieldValuesInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsMandatory() interface{}
	SetIsMandatory(val interface{})
	IsMandatoryInput() interface{}
	IsStandardTimeRangeFilter() interface{}
	SetIsStandardTimeRangeFilter(val interface{})
	IsStandardTimeRangeFilterEnabled() interface{}
	SetIsStandardTimeRangeFilterEnabled(val interface{})
	IsStandardTimeRangeFilterEnabledInput() interface{}
	IsStandardTimeRangeFilterInput() interface{}
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
	PutFilterOperatorAndFieldValues(value interface{})
	ResetChartIds()
	ResetDataSource()
	ResetDisplayName()
	ResetFieldPath()
	ResetFilterOperatorAndFieldValues()
	ResetId()
	ResetIsMandatory()
	ResetIsStandardTimeRangeFilter()
	ResetIsStandardTimeRangeFilterEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleNativeDashboardFiltersOutputReference
type jsiiProxy_ChronicleNativeDashboardFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ChartIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"chartIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ChartIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"chartIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) FieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) FieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) FilterOperatorAndFieldValues() ChronicleNativeDashboardFiltersFilterOperatorAndFieldValuesList {
	var returns ChronicleNativeDashboardFiltersFilterOperatorAndFieldValuesList
	_jsii_.Get(
		j,
		"filterOperatorAndFieldValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) FilterOperatorAndFieldValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterOperatorAndFieldValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) IsMandatory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMandatory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) IsMandatoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMandatoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleNativeDashboardFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ChronicleNativeDashboardFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleNativeDashboardFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleNativeDashboardFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleNativeDashboard.ChronicleNativeDashboardFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewChronicleNativeDashboardFiltersOutputReference_Override(c ChronicleNativeDashboardFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleNativeDashboard.ChronicleNativeDashboardFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetChartIds(val *[]*string) {
	if err := j.validateSetChartIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chartIds",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetFieldPath(val *string) {
	if err := j.validateSetFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldPath",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetIsMandatory(val interface{}) {
	if err := j.validateSetIsMandatoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMandatory",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetIsStandardTimeRangeFilter(val interface{}) {
	if err := j.validateSetIsStandardTimeRangeFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isStandardTimeRangeFilter",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetIsStandardTimeRangeFilterEnabled(val interface{}) {
	if err := j.validateSetIsStandardTimeRangeFilterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isStandardTimeRangeFilterEnabled",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) PutFilterOperatorAndFieldValues(value interface{}) {
	if err := c.validatePutFilterOperatorAndFieldValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFilterOperatorAndFieldValues",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetChartIds() {
	_jsii_.InvokeVoid(
		c,
		"resetChartIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		c,
		"resetDataSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		c,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetFieldPath() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetFilterOperatorAndFieldValues() {
	_jsii_.InvokeVoid(
		c,
		"resetFilterOperatorAndFieldValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetIsMandatory() {
	_jsii_.InvokeVoid(
		c,
		"resetIsMandatory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetIsStandardTimeRangeFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetIsStandardTimeRangeFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ResetIsStandardTimeRangeFilterEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetIsStandardTimeRangeFilterEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleNativeDashboardFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

