// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsTrellixHxHostsSettingsAuthenticationMsso struct {
	// The login api endpoint url.
	//
	// This must be a valid URL with an http or https scheme. It has no default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/chronicle_feed#api_endpoint ChronicleFeed#api_endpoint}
	ApiEndpoint *string `field:"required" json:"apiEndpoint" yaml:"apiEndpoint"`
	// Password of the account identified by username.
	//
	// There are no restrictions on the format of the password. It has no default,
	// specifically enforced min / max length or character set. The password
	// will have been provided by an MSSO administrator and it is assumed that
	// they have provided a password that is internally consistent with MSSO
	// authentication requirements / validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/chronicle_feed#password ChronicleFeed#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for MSSO authentication.
	//
	// There are no restrictions on the format of the username. It has no default,
	// specifically enforced min / max length or character set. The username
	// will have been provided by an MSSO administrator and it is assumed that
	// they have provided a username that is internally consistent with MSSO
	// authentication requirements / validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/chronicle_feed#username ChronicleFeed#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

