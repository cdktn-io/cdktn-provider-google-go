// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cesagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CesAgentTransferRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CesAgentTransferRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CesAgentTransferRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CesAgentTransferRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CesAgentTransferRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CesAgentTransferRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CesAgentTransferRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCesAgentTransferRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

