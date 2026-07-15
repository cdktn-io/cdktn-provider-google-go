// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containernodepool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerNodePoolMaintenancePolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerNodePoolMaintenancePolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerNodePoolMaintenancePolicyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolMaintenancePolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolMaintenancePolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolMaintenancePolicyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolMaintenancePolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerNodePoolMaintenancePolicyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

