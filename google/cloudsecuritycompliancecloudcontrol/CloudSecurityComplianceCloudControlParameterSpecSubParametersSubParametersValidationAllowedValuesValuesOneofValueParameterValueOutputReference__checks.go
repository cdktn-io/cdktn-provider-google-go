// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package cloudsecuritycompliancecloudcontrol

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validatePutStringListValueParameters(value *CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValue) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetBoolValueParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetInternalValueParameters(val *CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValue) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetNumberValueParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetStringValueParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCloudSecurityComplianceCloudControlParameterSpecSubParametersSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

