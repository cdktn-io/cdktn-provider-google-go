// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleProviderFunctions) validateLocationFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateNameFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateProjectFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateRegionFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateRegionFromZoneParameters(zone *string) error {
	return nil
}

func (g *jsiiProxy_GoogleProviderFunctions) validateZoneFromIdParameters(id *string) error {
	return nil
}

func validateNewGoogleProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

