// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package chronicleparser

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChronicleParserCreatorList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ChronicleParserCreatorList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChronicleParserCreatorList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChronicleParserCreatorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChronicleParserCreatorList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChronicleParserCreatorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChronicleParserCreatorListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

