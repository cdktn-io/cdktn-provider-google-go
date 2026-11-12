// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataformconfig

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataformConfig) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validatePutTimeoutsParameters(value *DataformConfigTimeouts) error {
	return nil
}

func (d *jsiiProxy_DataformConfig) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateDataformConfig_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDataformConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataformConfig_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDataformConfig_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetDefaultKmsKeyNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DataformConfig) validateSetRegionParameters(val *string) error {
	return nil
}

func validateNewDataformConfigParameters(scope constructs.Construct, id *string, config *DataformConfigConfig) error {
	return nil
}

