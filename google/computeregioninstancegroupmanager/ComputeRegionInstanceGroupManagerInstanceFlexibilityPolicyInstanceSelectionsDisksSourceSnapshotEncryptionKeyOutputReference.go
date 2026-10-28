// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/computeregioninstancegroupmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference interface {
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
	InternalValue() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey
	SetInternalValue(val *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey)
	KmsKeySelfLink() *string
	SetKmsKeySelfLink(val *string)
	KmsKeySelfLinkInput() *string
	KmsKeyServiceAccount() *string
	SetKmsKeyServiceAccount(val *string)
	KmsKeyServiceAccountInput() *string
	RawKey() *string
	SetRawKey(val *string)
	RawKeyInput() *string
	RsaEncryptedKey() *string
	SetRsaEncryptedKey(val *string)
	RsaEncryptedKeyInput() *string
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
	ResetKmsKeySelfLink()
	ResetKmsKeyServiceAccount()
	ResetRawKey()
	ResetRsaEncryptedKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference
type jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) InternalValue() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey {
	var returns *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) KmsKeySelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) KmsKeySelfLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) KmsKeyServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) KmsKeyServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) RawKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) RawKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) RsaEncryptedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaEncryptedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) RsaEncryptedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaEncryptedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference_Override(c ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetInternalValue(val *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetKmsKeySelfLink(val *string) {
	if err := j.validateSetKmsKeySelfLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeySelfLink",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetKmsKeyServiceAccount(val *string) {
	if err := j.validateSetKmsKeyServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyServiceAccount",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetRawKey(val *string) {
	if err := j.validateSetRawKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rawKey",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetRsaEncryptedKey(val *string) {
	if err := j.validateSetRsaEncryptedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaEncryptedKey",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ResetKmsKeySelfLink() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeySelfLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ResetKmsKeyServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ResetRawKey() {
	_jsii_.InvokeVoid(
		c,
		"resetRawKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ResetRsaEncryptedKey() {
	_jsii_.InvokeVoid(
		c,
		"resetRsaEncryptedKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

