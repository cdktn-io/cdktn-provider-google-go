// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppTimeZoneSettings struct {
	// The time zone of the app from the time zone database, e.g., America/Los_Angeles, Europe/Paris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/ces_app#time_zone CesApp#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

