// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedkafkacluster

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedKafkaClusterPublicClusterDetailsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedKafkaClusterPublicClusterDetailsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedKafkaClusterPublicClusterDetailsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedKafkaClusterPublicClusterDetailsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedKafkaClusterPublicClusterDetailsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedKafkaClusterPublicClusterDetailsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedKafkaClusterPublicClusterDetailsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

