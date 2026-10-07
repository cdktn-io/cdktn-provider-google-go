// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dataprocworkflowtemplate

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetDiskSizeGbParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetDiskTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfig:
		val := val.(*DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfig)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfig:
		val_ := val.(DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfig)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfig; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetProvisionedIopsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetProvisionedThroughputParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigAttachedDiskConfigOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

