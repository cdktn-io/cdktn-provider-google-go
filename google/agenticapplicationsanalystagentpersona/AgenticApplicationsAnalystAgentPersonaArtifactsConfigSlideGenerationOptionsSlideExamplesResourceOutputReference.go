// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference interface {
	cdktn.ComplexObject
	BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference
	BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource
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
	F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1ResourceOutputReference
	F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1Resource
	// Experimental.
	Fqn() *string
	GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResourceOutputReference
	GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResource
	GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResourceOutputReference
	GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResource
	InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResource
	SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResource)
	ModelDescription() *string
	SetModelDescription(val *string)
	ModelDescriptionInput() *string
	RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResourceOutputReference
	RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResource
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
	PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource)
	PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1Resource)
	PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResource)
	PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResource)
	PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResource)
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

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference
type jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference
	_jsii_.Get(
		j,
		"bigqueryResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource
	_jsii_.Get(
		j,
		"bigqueryResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) DisplayLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) DisplayLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1ResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1ResourceOutputReference
	_jsii_.Get(
		j,
		"f1Resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1Resource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1Resource
	_jsii_.Get(
		j,
		"f1ResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResourceOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResource
	_jsii_.Get(
		j,
		"googleCloudStorageResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResourceOutputReference
	_jsii_.Get(
		j,
		"googleDriveResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResource
	_jsii_.Get(
		j,
		"googleDriveResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResourceOutputReference
	_jsii_.Get(
		j,
		"rawFileResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResource
	_jsii_.Get(
		j,
		"rawFileResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) UseRag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) UseRagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRagInput",
		&returns,
	)
	return returns
}


func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference_Override(a AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetDisplayLabel(val *string) {
	if err := j.validateSetDisplayLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayLabel",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetModelDescription(val *string) {
	if err := j.validateSetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDescription",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)SetUseRag(val interface{}) {
	if err := j.validateSetUseRagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRag",
		val,
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource) {
	if err := a.validatePutBigqueryResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBigqueryResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1Resource) {
	if err := a.validatePutF1ResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putF1Resource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResource) {
	if err := a.validatePutGoogleCloudStorageResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleCloudStorageResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResource) {
	if err := a.validatePutGoogleDriveResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleDriveResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResource) {
	if err := a.validatePutRawFileResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRawFileResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetBigqueryResource() {
	_jsii_.InvokeVoid(
		a,
		"resetBigqueryResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetDisplayLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetF1Resource() {
	_jsii_.InvokeVoid(
		a,
		"resetF1Resource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetGoogleCloudStorageResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleCloudStorageResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetGoogleDriveResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleDriveResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetModelDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetRawFileResource() {
	_jsii_.InvokeVoid(
		a,
		"resetRawFileResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ResetUseRag() {
	_jsii_.InvokeVoid(
		a,
		"resetUseRag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

