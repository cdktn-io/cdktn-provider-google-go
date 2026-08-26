// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsThreatConnectIocV3Settings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsThreatConnectIocV3SettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#fields ChronicleFeed#fields}
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Owners.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#owners ChronicleFeed#owners}
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
	// Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#schedule ChronicleFeed#schedule}
	Schedule *float64 `field:"optional" json:"schedule" yaml:"schedule"`
	// ThreatConnect Query Language filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#tql_query ChronicleFeed#tql_query}
	TqlQuery *string `field:"optional" json:"tqlQuery" yaml:"tqlQuery"`
}

