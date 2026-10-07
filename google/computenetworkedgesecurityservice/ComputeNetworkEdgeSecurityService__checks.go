// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package computenetworkedgesecurityservice

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateAddMoveTargetParameters(moveTarget *string) error {
	if moveTarget == nil {
		return fmt.Errorf("parameter moveTarget is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateImportFromParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateMoveFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateMoveToParameters(moveTarget *string, index interface{}) error {
	if moveTarget == nil {
		return fmt.Errorf("parameter moveTarget is required, but nil was provided")
	}

	switch index.(type) {
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
		return fmt.Errorf("parameter index must be one of the allowed types: *string, *float64; received %#v (a %T)", index, index)
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateMoveToIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validatePutTimeoutsParameters(value *ComputeNetworkEdgeSecurityServiceTimeouts) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ComputeNetworkEdgeSecurityService) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	if feature == "" {
		return fmt.Errorf("parameter feature is required, but nil was provided")
	}

	return nil
}

func validateComputeNetworkEdgeSecurityService_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if importToId == nil {
		return fmt.Errorf("parameter importToId is required, but nil was provided")
	}

	if importFromId == nil {
		return fmt.Errorf("parameter importFromId is required, but nil was provided")
	}

	return nil
}

func validateComputeNetworkEdgeSecurityService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateComputeNetworkEdgeSecurityService_IsTerraformElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateComputeNetworkEdgeSecurityService_IsTerraformResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetConnectionParameters(val interface{}) error {
	switch val.(type) {
	case *cdktn.SSHProvisionerConnection:
		val := val.(*cdktn.SSHProvisionerConnection)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktn.SSHProvisionerConnection:
		val_ := val.(cdktn.SSHProvisionerConnection)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case *cdktn.WinrmProvisionerConnection:
		val := val.(*cdktn.WinrmProvisionerConnection)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktn.WinrmProvisionerConnection:
		val_ := val.(cdktn.WinrmProvisionerConnection)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *cdktn.SSHProvisionerConnection, *cdktn.WinrmProvisionerConnection; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetCountParameters(val interface{}) error {
	switch val.(type) {
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
	case cdktn.TerraformCount:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *float64, cdktn.TerraformCount; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetDeletionPolicyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetDescriptionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetProjectParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetProvisionersParameters(val *[]interface{}) error {
	for idx_97dfc6, v := range *val {
		switch v.(type) {
		case *cdktn.FileProvisioner:
			v := v.(*cdktn.FileProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktn.FileProvisioner:
			v_ := v.(cdktn.FileProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case *cdktn.LocalExecProvisioner:
			v := v.(*cdktn.LocalExecProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktn.LocalExecProvisioner:
			v_ := v.(cdktn.LocalExecProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case *cdktn.RemoteExecProvisioner:
			v := v.(*cdktn.RemoteExecProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktn.RemoteExecProvisioner:
			v_ := v.(cdktn.RemoteExecProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *cdktn.FileProvisioner, *cdktn.LocalExecProvisioner, *cdktn.RemoteExecProvisioner; received %#v (a %T)", idx_97dfc6, v, v)
			}
		}
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetRegionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeNetworkEdgeSecurityService) validateSetSecurityPolicyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewComputeNetworkEdgeSecurityServiceParameters(scope constructs.Construct, id *string, config *ComputeNetworkEdgeSecurityServiceConfig) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

