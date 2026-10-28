// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/pubsubsubscription/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PubsubSubscriptionBigtableConfigOutputReference interface {
	cdktn.ComplexObject
	AppProfileId() *string
	SetAppProfileId(val *string)
	AppProfileIdInput() *string
	ColumnFamilyMapping() PubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference
	ColumnFamilyMappingInput() *PubsubSubscriptionBigtableConfigColumnFamilyMapping
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
	InternalValue() *PubsubSubscriptionBigtableConfig
	SetInternalValue(val *PubsubSubscriptionBigtableConfig)
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	State() *string
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WriteMetadata() interface{}
	SetWriteMetadata(val interface{})
	WriteMetadataInput() interface{}
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
	PutColumnFamilyMapping(value *PubsubSubscriptionBigtableConfigColumnFamilyMapping)
	ResetAppProfileId()
	ResetColumnFamilyMapping()
	ResetServiceAccountEmail()
	ResetWriteMetadata()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubSubscriptionBigtableConfigOutputReference
type jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) AppProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) AppProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ColumnFamilyMapping() PubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference {
	var returns PubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference
	_jsii_.Get(
		j,
		"columnFamilyMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ColumnFamilyMappingInput() *PubsubSubscriptionBigtableConfigColumnFamilyMapping {
	var returns *PubsubSubscriptionBigtableConfigColumnFamilyMapping
	_jsii_.Get(
		j,
		"columnFamilyMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) InternalValue() *PubsubSubscriptionBigtableConfig {
	var returns *PubsubSubscriptionBigtableConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) WriteMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) WriteMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeMetadataInput",
		&returns,
	)
	return returns
}


func NewPubsubSubscriptionBigtableConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PubsubSubscriptionBigtableConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubSubscriptionBigtableConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.pubsubSubscription.PubsubSubscriptionBigtableConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPubsubSubscriptionBigtableConfigOutputReference_Override(p PubsubSubscriptionBigtableConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.pubsubSubscription.PubsubSubscriptionBigtableConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetAppProfileId(val *string) {
	if err := j.validateSetAppProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appProfileId",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetInternalValue(val *PubsubSubscriptionBigtableConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference)SetWriteMetadata(val interface{}) {
	if err := j.validateSetWriteMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeMetadata",
		val,
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) PutColumnFamilyMapping(value *PubsubSubscriptionBigtableConfigColumnFamilyMapping) {
	if err := p.validatePutColumnFamilyMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putColumnFamilyMapping",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ResetAppProfileId() {
	_jsii_.InvokeVoid(
		p,
		"resetAppProfileId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ResetColumnFamilyMapping() {
	_jsii_.InvokeVoid(
		p,
		"resetColumnFamilyMapping",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ResetWriteMetadata() {
	_jsii_.InvokeVoid(
		p,
		"resetWriteMetadata",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

