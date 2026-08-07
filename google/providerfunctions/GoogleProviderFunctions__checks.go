// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package providerfunctions

import (
	"fmt"
)

func (g *jsiiProxy_GoogleProviderFunctions) validateLocationFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateNameFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateProjectFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateRegionFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateRegionFromZoneParameters(zone *string) error {
	if zone == nil {
		return fmt.Errorf("parameter zone is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateZoneFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleProviderFunctionsParameters(providerLocalName *string) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	return nil
}

