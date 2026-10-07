// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/vertexairagcorpus/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference interface {
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
	InternalValue() *VertexAiRagCorpusVectorDbConfigRagManagedDbAnn
	SetInternalValue(val *VertexAiRagCorpusVectorDbConfigRagManagedDbAnn)
	LeafCount() *float64
	SetLeafCount(val *float64)
	LeafCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TreeDepth() *float64
	SetTreeDepth(val *float64)
	TreeDepthInput() *float64
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
	ResetLeafCount()
	ResetTreeDepth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference
type jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) InternalValue() *VertexAiRagCorpusVectorDbConfigRagManagedDbAnn {
	var returns *VertexAiRagCorpusVectorDbConfigRagManagedDbAnn
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) LeafCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) LeafCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) TreeDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"treeDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) TreeDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"treeDepthInput",
		&returns,
	)
	return returns
}


func NewVertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagCorpus.VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference_Override(v VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagCorpus.VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference)SetInternalValue(val *VertexAiRagCorpusVectorDbConfigRagManagedDbAnn) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference)SetLeafCount(val *float64) {
	if err := j.validateSetLeafCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leafCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference)SetTreeDepth(val *float64) {
	if err := j.validateSetTreeDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treeDepth",
		val,
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) ResetLeafCount() {
	_jsii_.InvokeVoid(
		v,
		"resetLeafCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) ResetTreeDepth() {
	_jsii_.InvokeVoid(
		v,
		"resetTreeDepth",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigRagManagedDbAnnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

