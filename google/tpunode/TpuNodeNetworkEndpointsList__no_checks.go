// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package tpunode

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TpuNodeNetworkEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TpuNodeNetworkEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TpuNodeNetworkEndpointsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TpuNodeNetworkEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TpuNodeNetworkEndpointsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TpuNodeNetworkEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTpuNodeNetworkEndpointsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

