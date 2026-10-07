// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakehivetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/biglakehivetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BiglakeHiveTableStorageDescriptorOutputReference interface {
	cdktn.ComplexObject
	BucketCols() *[]*string
	SetBucketCols(val *[]*string)
	BucketColsInput() *[]*string
	Columns() BiglakeHiveTableStorageDescriptorColumnsList
	ColumnsInput() interface{}
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
	Compressed() interface{}
	SetCompressed(val interface{})
	CompressedInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	InternalValue() *BiglakeHiveTableStorageDescriptor
	SetInternalValue(val *BiglakeHiveTableStorageDescriptor)
	LocationUri() *string
	SetLocationUri(val *string)
	LocationUriInput() *string
	NumBuckets() *float64
	SetNumBuckets(val *float64)
	NumBucketsInput() *float64
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	SerdeInfo() BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference
	SerdeInfoInput() *BiglakeHiveTableStorageDescriptorSerdeInfo
	SkewedInfo() BiglakeHiveTableStorageDescriptorSkewedInfoOutputReference
	SkewedInfoInput() *BiglakeHiveTableStorageDescriptorSkewedInfo
	SortCols() BiglakeHiveTableStorageDescriptorSortColsList
	SortColsInput() interface{}
	StoredAsSubDirs() interface{}
	SetStoredAsSubDirs(val interface{})
	StoredAsSubDirsInput() interface{}
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
	PutColumns(value interface{})
	PutSerdeInfo(value *BiglakeHiveTableStorageDescriptorSerdeInfo)
	PutSkewedInfo(value *BiglakeHiveTableStorageDescriptorSkewedInfo)
	PutSortCols(value interface{})
	ResetBucketCols()
	ResetCompressed()
	ResetInputFormat()
	ResetLocationUri()
	ResetNumBuckets()
	ResetOutputFormat()
	ResetParameters()
	ResetSerdeInfo()
	ResetSkewedInfo()
	ResetSortCols()
	ResetStoredAsSubDirs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BiglakeHiveTableStorageDescriptorOutputReference
type jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) BucketCols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketCols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) BucketColsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) Columns() BiglakeHiveTableStorageDescriptorColumnsList {
	var returns BiglakeHiveTableStorageDescriptorColumnsList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) InternalValue() *BiglakeHiveTableStorageDescriptor {
	var returns *BiglakeHiveTableStorageDescriptor
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) LocationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) NumBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) NumBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) SerdeInfo() BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference {
	var returns BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference
	_jsii_.Get(
		j,
		"serdeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) SerdeInfoInput() *BiglakeHiveTableStorageDescriptorSerdeInfo {
	var returns *BiglakeHiveTableStorageDescriptorSerdeInfo
	_jsii_.Get(
		j,
		"serdeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) SkewedInfo() BiglakeHiveTableStorageDescriptorSkewedInfoOutputReference {
	var returns BiglakeHiveTableStorageDescriptorSkewedInfoOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) SkewedInfoInput() *BiglakeHiveTableStorageDescriptorSkewedInfo {
	var returns *BiglakeHiveTableStorageDescriptorSkewedInfo
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) SortCols() BiglakeHiveTableStorageDescriptorSortColsList {
	var returns BiglakeHiveTableStorageDescriptorSortColsList
	_jsii_.Get(
		j,
		"sortCols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) SortColsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) StoredAsSubDirs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) StoredAsSubDirsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBiglakeHiveTableStorageDescriptorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BiglakeHiveTableStorageDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewBiglakeHiveTableStorageDescriptorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.biglakeHiveTable.BiglakeHiveTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBiglakeHiveTableStorageDescriptorOutputReference_Override(b BiglakeHiveTableStorageDescriptorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.biglakeHiveTable.BiglakeHiveTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetBucketCols(val *[]*string) {
	if err := j.validateSetBucketColsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketCols",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetInternalValue(val *BiglakeHiveTableStorageDescriptor) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetLocationUri(val *string) {
	if err := j.validateSetLocationUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationUri",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetNumBuckets(val *float64) {
	if err := j.validateSetNumBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numBuckets",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetStoredAsSubDirs(val interface{}) {
	if err := j.validateSetStoredAsSubDirsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirs",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) PutColumns(value interface{}) {
	if err := b.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putColumns",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) PutSerdeInfo(value *BiglakeHiveTableStorageDescriptorSerdeInfo) {
	if err := b.validatePutSerdeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSerdeInfo",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) PutSkewedInfo(value *BiglakeHiveTableStorageDescriptorSkewedInfo) {
	if err := b.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) PutSortCols(value interface{}) {
	if err := b.validatePutSortColsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSortCols",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetBucketCols() {
	_jsii_.InvokeVoid(
		b,
		"resetBucketCols",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		b,
		"resetCompressed",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetLocationUri() {
	_jsii_.InvokeVoid(
		b,
		"resetLocationUri",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetNumBuckets() {
	_jsii_.InvokeVoid(
		b,
		"resetNumBuckets",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetSerdeInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetSerdeInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetSortCols() {
	_jsii_.InvokeVoid(
		b,
		"resetSortCols",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ResetStoredAsSubDirs() {
	_jsii_.InvokeVoid(
		b,
		"resetStoredAsSubDirs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

