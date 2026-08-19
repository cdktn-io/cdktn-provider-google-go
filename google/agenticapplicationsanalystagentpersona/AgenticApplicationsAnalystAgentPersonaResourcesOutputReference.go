// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgenticApplicationsAnalystAgentPersonaResourcesOutputReference interface {
	cdktn.ComplexObject
	BigqueryResource() AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResourceOutputReference
	BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource
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
	DisplayLabel() *string
	SetDisplayLabel(val *string)
	DisplayLabelInput() *string
	F1Resource() AgenticApplicationsAnalystAgentPersonaResourcesF1ResourceOutputReference
	F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesF1Resource
	// Experimental.
	Fqn() *string
	GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResourceOutputReference
	GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResource
	GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResourceOutputReference
	GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelDescription() *string
	SetModelDescription(val *string)
	ModelDescriptionInput() *string
	RawFileResource() AgenticApplicationsAnalystAgentPersonaResourcesRawFileResourceOutputReference
	RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesRawFileResource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseRag() interface{}
	SetUseRag(val interface{})
	UseRagInput() interface{}
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
	PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource)
	PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaResourcesF1Resource)
	PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResource)
	PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource)
	PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaResourcesRawFileResource)
	ResetBigqueryResource()
	ResetDisplayLabel()
	ResetF1Resource()
	ResetGoogleCloudStorageResource()
	ResetGoogleDriveResource()
	ResetModelDescription()
	ResetRawFileResource()
	ResetUseRag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersonaResourcesOutputReference
type jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) BigqueryResource() AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResourceOutputReference
	_jsii_.Get(
		j,
		"bigqueryResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource {
	var returns *AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource
	_jsii_.Get(
		j,
		"bigqueryResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) DisplayLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) DisplayLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) F1Resource() AgenticApplicationsAnalystAgentPersonaResourcesF1ResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaResourcesF1ResourceOutputReference
	_jsii_.Get(
		j,
		"f1Resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesF1Resource {
	var returns *AgenticApplicationsAnalystAgentPersonaResourcesF1Resource
	_jsii_.Get(
		j,
		"f1ResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResourceOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResource {
	var returns *AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResource
	_jsii_.Get(
		j,
		"googleCloudStorageResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResourceOutputReference
	_jsii_.Get(
		j,
		"googleDriveResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource {
	var returns *AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource
	_jsii_.Get(
		j,
		"googleDriveResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) RawFileResource() AgenticApplicationsAnalystAgentPersonaResourcesRawFileResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaResourcesRawFileResourceOutputReference
	_jsii_.Get(
		j,
		"rawFileResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaResourcesRawFileResource {
	var returns *AgenticApplicationsAnalystAgentPersonaResourcesRawFileResource
	_jsii_.Get(
		j,
		"rawFileResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) UseRag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) UseRagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRagInput",
		&returns,
	)
	return returns
}


func NewAgenticApplicationsAnalystAgentPersonaResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AgenticApplicationsAnalystAgentPersonaResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAgenticApplicationsAnalystAgentPersonaResourcesOutputReference_Override(a AgenticApplicationsAnalystAgentPersonaResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetDisplayLabel(val *string) {
	if err := j.validateSetDisplayLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayLabel",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetModelDescription(val *string) {
	if err := j.validateSetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDescription",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)SetUseRag(val interface{}) {
	if err := j.validateSetUseRagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRag",
		val,
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource) {
	if err := a.validatePutBigqueryResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBigqueryResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaResourcesF1Resource) {
	if err := a.validatePutF1ResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putF1Resource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResource) {
	if err := a.validatePutGoogleCloudStorageResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleCloudStorageResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource) {
	if err := a.validatePutGoogleDriveResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleDriveResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaResourcesRawFileResource) {
	if err := a.validatePutRawFileResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRawFileResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetBigqueryResource() {
	_jsii_.InvokeVoid(
		a,
		"resetBigqueryResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetDisplayLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetF1Resource() {
	_jsii_.InvokeVoid(
		a,
		"resetF1Resource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetGoogleCloudStorageResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleCloudStorageResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetGoogleDriveResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleDriveResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetModelDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetRawFileResource() {
	_jsii_.InvokeVoid(
		a,
		"resetRawFileResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ResetUseRag() {
	_jsii_.InvokeVoid(
		a,
		"resetUseRag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

