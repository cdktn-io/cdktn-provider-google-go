// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference interface {
	cdktn.ComplexObject
	BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResourceOutputReference
	BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource
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
	F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1ResourceOutputReference
	F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource
	// Experimental.
	Fqn() *string
	GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResourceOutputReference
	GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource
	GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResourceOutputReference
	GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource
	InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource
	SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource)
	ModelDescription() *string
	SetModelDescription(val *string)
	ModelDescriptionInput() *string
	RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResourceOutputReference
	RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource
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
	PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource)
	PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource)
	PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource)
	PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource)
	PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource)
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

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference
type jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResourceOutputReference
	_jsii_.Get(
		j,
		"bigqueryResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource
	_jsii_.Get(
		j,
		"bigqueryResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) DisplayLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) DisplayLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1ResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1ResourceOutputReference
	_jsii_.Get(
		j,
		"f1Resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource
	_jsii_.Get(
		j,
		"f1ResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResourceOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource
	_jsii_.Get(
		j,
		"googleCloudStorageResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResourceOutputReference
	_jsii_.Get(
		j,
		"googleDriveResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource
	_jsii_.Get(
		j,
		"googleDriveResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResourceOutputReference
	_jsii_.Get(
		j,
		"rawFileResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource
	_jsii_.Get(
		j,
		"rawFileResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) UseRag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) UseRagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRagInput",
		&returns,
	)
	return returns
}


func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference_Override(a AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetDisplayLabel(val *string) {
	if err := j.validateSetDisplayLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayLabel",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetModelDescription(val *string) {
	if err := j.validateSetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDescription",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)SetUseRag(val interface{}) {
	if err := j.validateSetUseRagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRag",
		val,
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource) {
	if err := a.validatePutBigqueryResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBigqueryResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource) {
	if err := a.validatePutF1ResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putF1Resource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource) {
	if err := a.validatePutGoogleCloudStorageResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleCloudStorageResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource) {
	if err := a.validatePutGoogleDriveResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleDriveResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource) {
	if err := a.validatePutRawFileResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRawFileResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetBigqueryResource() {
	_jsii_.InvokeVoid(
		a,
		"resetBigqueryResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetDisplayLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetF1Resource() {
	_jsii_.InvokeVoid(
		a,
		"resetF1Resource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetGoogleCloudStorageResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleCloudStorageResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetGoogleDriveResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleDriveResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetModelDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetRawFileResource() {
	_jsii_.InvokeVoid(
		a,
		"resetRawFileResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ResetUseRag() {
	_jsii_.InvokeVoid(
		a,
		"resetUseRag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

