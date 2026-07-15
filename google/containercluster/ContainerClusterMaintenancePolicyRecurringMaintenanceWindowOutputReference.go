// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/containercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference interface {
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
	DelayUntil() ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntilOutputReference
	DelayUntilInput() *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerClusterMaintenancePolicyRecurringMaintenanceWindow
	SetInternalValue(val *ContainerClusterMaintenancePolicyRecurringMaintenanceWindow)
	Recurrence() *string
	SetRecurrence(val *string)
	RecurrenceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WindowDuration() *string
	SetWindowDuration(val *string)
	WindowDurationInput() *string
	WindowStartTime() ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTimeOutputReference
	WindowStartTimeInput() *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime
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
	PutDelayUntil(value *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil)
	PutWindowStartTime(value *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime)
	ResetDelayUntil()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference
type jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) DelayUntil() ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntilOutputReference {
	var returns ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntilOutputReference
	_jsii_.Get(
		j,
		"delayUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) DelayUntilInput() *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil {
	var returns *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil
	_jsii_.Get(
		j,
		"delayUntilInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) InternalValue() *ContainerClusterMaintenancePolicyRecurringMaintenanceWindow {
	var returns *ContainerClusterMaintenancePolicyRecurringMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) Recurrence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) RecurrenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowStartTime() ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTimeOutputReference {
	var returns ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTimeOutputReference
	_jsii_.Get(
		j,
		"windowStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowStartTimeInput() *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime {
	var returns *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime
	_jsii_.Get(
		j,
		"windowStartTimeInput",
		&returns,
	)
	return returns
}


func NewContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference_Override(c ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetInternalValue(val *ContainerClusterMaintenancePolicyRecurringMaintenanceWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetRecurrence(val *string) {
	if err := j.validateSetRecurrenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrence",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetWindowDuration(val *string) {
	if err := j.validateSetWindowDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowDuration",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) PutDelayUntil(value *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil) {
	if err := c.validatePutDelayUntilParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDelayUntil",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) PutWindowStartTime(value *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime) {
	if err := c.validatePutWindowStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWindowStartTime",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ResetDelayUntil() {
	_jsii_.InvokeVoid(
		c,
		"resetDelayUntil",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

