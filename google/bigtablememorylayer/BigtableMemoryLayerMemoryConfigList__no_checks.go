// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bigtablememorylayer

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BigtableMemoryLayerMemoryConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BigtableMemoryLayerMemoryConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BigtableMemoryLayerMemoryConfigList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BigtableMemoryLayerMemoryConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BigtableMemoryLayerMemoryConfigList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BigtableMemoryLayerMemoryConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBigtableMemoryLayerMemoryConfigListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

