// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference interface {
	cdktn.ComplexObject
	BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResourceOutputReference
	BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResource
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
	F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1ResourceOutputReference
	F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1Resource
	// Experimental.
	Fqn() *string
	GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResourceOutputReference
	GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResource
	GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResourceOutputReference
	GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResource
	InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResource
	SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResource)
	ModelDescription() *string
	SetModelDescription(val *string)
	ModelDescriptionInput() *string
	RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResourceOutputReference
	RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource
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
	PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResource)
	PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1Resource)
	PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResource)
	PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResource)
	PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource)
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

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference
type jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResourceOutputReference
	_jsii_.Get(
		j,
		"bigqueryResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResource
	_jsii_.Get(
		j,
		"bigqueryResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) DisplayLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) DisplayLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1ResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1ResourceOutputReference
	_jsii_.Get(
		j,
		"f1Resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1Resource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1Resource
	_jsii_.Get(
		j,
		"f1ResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResourceOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResource
	_jsii_.Get(
		j,
		"googleCloudStorageResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResourceOutputReference
	_jsii_.Get(
		j,
		"googleDriveResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResource
	_jsii_.Get(
		j,
		"googleDriveResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResourceOutputReference
	_jsii_.Get(
		j,
		"rawFileResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource
	_jsii_.Get(
		j,
		"rawFileResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) UseRag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) UseRagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRagInput",
		&returns,
	)
	return returns
}


func NewAgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference_Override(a AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetDisplayLabel(val *string) {
	if err := j.validateSetDisplayLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayLabel",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetModelDescription(val *string) {
	if err := j.validateSetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDescription",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)SetUseRag(val interface{}) {
	if err := j.validateSetUseRagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRag",
		val,
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResource) {
	if err := a.validatePutBigqueryResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBigqueryResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1Resource) {
	if err := a.validatePutF1ResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putF1Resource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResource) {
	if err := a.validatePutGoogleCloudStorageResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleCloudStorageResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResource) {
	if err := a.validatePutGoogleDriveResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleDriveResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource) {
	if err := a.validatePutRawFileResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRawFileResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetBigqueryResource() {
	_jsii_.InvokeVoid(
		a,
		"resetBigqueryResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetDisplayLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetF1Resource() {
	_jsii_.InvokeVoid(
		a,
		"resetF1Resource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetGoogleCloudStorageResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleCloudStorageResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetGoogleDriveResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleDriveResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetModelDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetRawFileResource() {
	_jsii_.InvokeVoid(
		a,
		"resetRawFileResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ResetUseRag() {
	_jsii_.InvokeVoid(
		a,
		"resetUseRag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

