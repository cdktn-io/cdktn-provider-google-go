// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package agentregistryservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgentRegistryServiceInterfacesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AgentRegistryServiceInterfacesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AgentRegistryServiceInterfacesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AgentRegistryServiceInterfacesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AgentRegistryServiceInterfacesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AgentRegistryServiceInterfacesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AgentRegistryServiceInterfacesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAgentRegistryServiceInterfacesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

