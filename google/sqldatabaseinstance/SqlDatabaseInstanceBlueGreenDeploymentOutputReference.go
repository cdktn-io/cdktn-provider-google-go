// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/sqldatabaseinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SqlDatabaseInstanceBlueGreenDeploymentOutputReference interface {
	cdktn.ComplexObject
	ActiveEnvironment() *string
	SetActiveEnvironment(val *string)
	ActiveEnvironmentInput() *string
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
	DeleteOldSourceOnDestroy() interface{}
	SetDeleteOldSourceOnDestroy(val interface{})
	DeleteOldSourceOnDestroyInput() interface{}
	DeploymentName() *string
	SetDeploymentName(val *string)
	DeploymentNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SqlDatabaseInstanceBlueGreenDeployment
	SetInternalValue(val *SqlDatabaseInstanceBlueGreenDeployment)
	OldSourceInstanceConnectionName() *string
	OldSourceInstanceName() *string
	TargetDatabaseVersion() *string
	SetTargetDatabaseVersion(val *string)
	TargetDatabaseVersionInput() *string
	TargetInstanceConnectionName() *string
	TargetInstanceIpAddresses() SqlDatabaseInstanceBlueGreenDeploymentTargetInstanceIpAddressesList
	TargetInstanceName() *string
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
	ResetActiveEnvironment()
	ResetDeleteOldSourceOnDestroy()
	ResetTargetDatabaseVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SqlDatabaseInstanceBlueGreenDeploymentOutputReference
type jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ActiveEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ActiveEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeleteOldSourceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOldSourceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeleteOldSourceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOldSourceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeploymentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeploymentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) InternalValue() *SqlDatabaseInstanceBlueGreenDeployment {
	var returns *SqlDatabaseInstanceBlueGreenDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) OldSourceInstanceConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldSourceInstanceConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) OldSourceInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldSourceInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetDatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDatabaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetDatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDatabaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetInstanceConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstanceConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetInstanceIpAddresses() SqlDatabaseInstanceBlueGreenDeploymentTargetInstanceIpAddressesList {
	var returns SqlDatabaseInstanceBlueGreenDeploymentTargetInstanceIpAddressesList
	_jsii_.Get(
		j,
		"targetInstanceIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSqlDatabaseInstanceBlueGreenDeploymentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SqlDatabaseInstanceBlueGreenDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewSqlDatabaseInstanceBlueGreenDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstanceBlueGreenDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlDatabaseInstanceBlueGreenDeploymentOutputReference_Override(s SqlDatabaseInstanceBlueGreenDeploymentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstanceBlueGreenDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetActiveEnvironment(val *string) {
	if err := j.validateSetActiveEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeEnvironment",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetDeleteOldSourceOnDestroy(val interface{}) {
	if err := j.validateSetDeleteOldSourceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOldSourceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetDeploymentName(val *string) {
	if err := j.validateSetDeploymentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetInternalValue(val *SqlDatabaseInstanceBlueGreenDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetTargetDatabaseVersion(val *string) {
	if err := j.validateSetTargetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDatabaseVersion",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ResetActiveEnvironment() {
	_jsii_.InvokeVoid(
		s,
		"resetActiveEnvironment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ResetDeleteOldSourceOnDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteOldSourceOnDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ResetTargetDatabaseVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetDatabaseVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstanceBlueGreenDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

