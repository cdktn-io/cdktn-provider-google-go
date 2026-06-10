// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cestool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CesToolRemoteAgentToolList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CesToolRemoteAgentToolList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CesToolRemoteAgentToolList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CesToolRemoteAgentToolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CesToolRemoteAgentToolList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CesToolRemoteAgentToolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCesToolRemoteAgentToolListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

