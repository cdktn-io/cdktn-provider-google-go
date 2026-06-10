// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cestool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CesToolConnectorToolList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CesToolConnectorToolList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CesToolConnectorToolList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CesToolConnectorToolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CesToolConnectorToolList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CesToolConnectorToolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCesToolConnectorToolListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

