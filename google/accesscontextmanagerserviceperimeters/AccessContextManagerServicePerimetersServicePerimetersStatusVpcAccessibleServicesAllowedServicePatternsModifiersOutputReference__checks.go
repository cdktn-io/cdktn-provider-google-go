// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package accesscontextmanagerserviceperimeters

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validatePutAddRequestHeaderParameters(value *AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersAddRequestHeader) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiers:
		val := val.(*AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiers)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiers:
		val_ := val.(AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiers)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiers; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

