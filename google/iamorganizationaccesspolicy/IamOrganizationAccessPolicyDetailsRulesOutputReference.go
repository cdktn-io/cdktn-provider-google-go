// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamorganizationaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/iamorganizationaccesspolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IamOrganizationAccessPolicyDetailsRulesOutputReference interface {
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
	Conditions() IamOrganizationAccessPolicyDetailsRulesConditionsList
	ConditionsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Effect() *string
	SetEffect(val *string)
	EffectInput() *string
	ExcludedPrincipals() *[]*string
	SetExcludedPrincipals(val *[]*string)
	ExcludedPrincipalsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Operation() IamOrganizationAccessPolicyDetailsRulesOperationOutputReference
	OperationInput() *IamOrganizationAccessPolicyDetailsRulesOperation
	Principals() *[]*string
	SetPrincipals(val *[]*string)
	PrincipalsInput() *[]*string
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
	PutConditions(value interface{})
	PutOperation(value *IamOrganizationAccessPolicyDetailsRulesOperation)
	ResetConditions()
	ResetDescription()
	ResetExcludedPrincipals()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamOrganizationAccessPolicyDetailsRulesOutputReference
type jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) Conditions() IamOrganizationAccessPolicyDetailsRulesConditionsList {
	var returns IamOrganizationAccessPolicyDetailsRulesConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) Effect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) EffectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ExcludedPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ExcludedPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) Operation() IamOrganizationAccessPolicyDetailsRulesOperationOutputReference {
	var returns IamOrganizationAccessPolicyDetailsRulesOperationOutputReference
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) OperationInput() *IamOrganizationAccessPolicyDetailsRulesOperation {
	var returns *IamOrganizationAccessPolicyDetailsRulesOperation
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) Principals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) PrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"principalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIamOrganizationAccessPolicyDetailsRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IamOrganizationAccessPolicyDetailsRulesOutputReference {
	_init_.Initialize()

	if err := validateNewIamOrganizationAccessPolicyDetailsRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.iamOrganizationAccessPolicy.IamOrganizationAccessPolicyDetailsRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIamOrganizationAccessPolicyDetailsRulesOutputReference_Override(i IamOrganizationAccessPolicyDetailsRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.iamOrganizationAccessPolicy.IamOrganizationAccessPolicyDetailsRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetEffect(val *string) {
	if err := j.validateSetEffectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effect",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetExcludedPrincipals(val *[]*string) {
	if err := j.validateSetExcludedPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPrincipals",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetPrincipals(val *[]*string) {
	if err := j.validateSetPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principals",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) PutConditions(value interface{}) {
	if err := i.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putConditions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) PutOperation(value *IamOrganizationAccessPolicyDetailsRulesOperation) {
	if err := i.validatePutOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOperation",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ResetConditions() {
	_jsii_.InvokeVoid(
		i,
		"resetConditions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ResetExcludedPrincipals() {
	_jsii_.InvokeVoid(
		i,
		"resetExcludedPrincipals",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamOrganizationAccessPolicyDetailsRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

