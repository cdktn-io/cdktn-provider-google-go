// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package firebaseremoteconfigremoteconfig


type FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue struct {
	// If true, the parameter is omitted from the parameter values returned to a client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/firebase_remote_config_remote_config#use_in_app_default FirebaseRemoteConfigRemoteConfig#use_in_app_default}
	UseInAppDefault interface{} `field:"optional" json:"useInAppDefault" yaml:"useInAppDefault"`
	// The string value that the parameter is set to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/firebase_remote_config_remote_config#value FirebaseRemoteConfigRemoteConfig#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

