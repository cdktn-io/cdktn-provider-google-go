// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/cloudrunv2instance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudRunV2InstanceVolumesOutputReference interface {
	cdktn.ComplexObject
	CloudSqlInstance() CloudRunV2InstanceVolumesCloudSqlInstanceOutputReference
	CloudSqlInstanceInput() *CloudRunV2InstanceVolumesCloudSqlInstance
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
	EmptyDir() CloudRunV2InstanceVolumesEmptyDirOutputReference
	EmptyDirInput() *CloudRunV2InstanceVolumesEmptyDir
	// Experimental.
	Fqn() *string
	Gcs() CloudRunV2InstanceVolumesGcsOutputReference
	GcsInput() *CloudRunV2InstanceVolumesGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() CloudRunV2InstanceVolumesNfsOutputReference
	NfsInput() *CloudRunV2InstanceVolumesNfs
	Secret() CloudRunV2InstanceVolumesSecretOutputReference
	SecretInput() *CloudRunV2InstanceVolumesSecret
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
	PutCloudSqlInstance(value *CloudRunV2InstanceVolumesCloudSqlInstance)
	PutEmptyDir(value *CloudRunV2InstanceVolumesEmptyDir)
	PutGcs(value *CloudRunV2InstanceVolumesGcs)
	PutNfs(value *CloudRunV2InstanceVolumesNfs)
	PutSecret(value *CloudRunV2InstanceVolumesSecret)
	ResetCloudSqlInstance()
	ResetEmptyDir()
	ResetGcs()
	ResetNfs()
	ResetSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunV2InstanceVolumesOutputReference
type jsiiProxy_CloudRunV2InstanceVolumesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) CloudSqlInstance() CloudRunV2InstanceVolumesCloudSqlInstanceOutputReference {
	var returns CloudRunV2InstanceVolumesCloudSqlInstanceOutputReference
	_jsii_.Get(
		j,
		"cloudSqlInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) CloudSqlInstanceInput() *CloudRunV2InstanceVolumesCloudSqlInstance {
	var returns *CloudRunV2InstanceVolumesCloudSqlInstance
	_jsii_.Get(
		j,
		"cloudSqlInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) EmptyDir() CloudRunV2InstanceVolumesEmptyDirOutputReference {
	var returns CloudRunV2InstanceVolumesEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) EmptyDirInput() *CloudRunV2InstanceVolumesEmptyDir {
	var returns *CloudRunV2InstanceVolumesEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) Gcs() CloudRunV2InstanceVolumesGcsOutputReference {
	var returns CloudRunV2InstanceVolumesGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GcsInput() *CloudRunV2InstanceVolumesGcs {
	var returns *CloudRunV2InstanceVolumesGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) Nfs() CloudRunV2InstanceVolumesNfsOutputReference {
	var returns CloudRunV2InstanceVolumesNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) NfsInput() *CloudRunV2InstanceVolumesNfs {
	var returns *CloudRunV2InstanceVolumesNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) Secret() CloudRunV2InstanceVolumesSecretOutputReference {
	var returns CloudRunV2InstanceVolumesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) SecretInput() *CloudRunV2InstanceVolumesSecret {
	var returns *CloudRunV2InstanceVolumesSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudRunV2InstanceVolumesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudRunV2InstanceVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunV2InstanceVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunV2InstanceVolumesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cloudRunV2Instance.CloudRunV2InstanceVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudRunV2InstanceVolumesOutputReference_Override(c CloudRunV2InstanceVolumesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cloudRunV2Instance.CloudRunV2InstanceVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2InstanceVolumesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) PutCloudSqlInstance(value *CloudRunV2InstanceVolumesCloudSqlInstance) {
	if err := c.validatePutCloudSqlInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudSqlInstance",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) PutEmptyDir(value *CloudRunV2InstanceVolumesEmptyDir) {
	if err := c.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) PutGcs(value *CloudRunV2InstanceVolumesGcs) {
	if err := c.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) PutNfs(value *CloudRunV2InstanceVolumesNfs) {
	if err := c.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNfs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) PutSecret(value *CloudRunV2InstanceVolumesSecret) {
	if err := c.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecret",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ResetCloudSqlInstance() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudSqlInstance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		c,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		c,
		"resetGcs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		c,
		"resetNfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudRunV2InstanceVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

