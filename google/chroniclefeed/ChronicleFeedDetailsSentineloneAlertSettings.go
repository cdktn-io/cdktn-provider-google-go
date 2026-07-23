// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsSentineloneAlertSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsSentineloneAlertSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Hostname of SentinelOne alert settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// initialStartTime from when to fetch the alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#initial_start_time ChronicleFeed#initial_start_time}
	InitialStartTime *string `field:"optional" json:"initialStartTime" yaml:"initialStartTime"`
	// Is the customer subscribed to Alerts Api.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#is_alert_api_subscribed ChronicleFeed#is_alert_api_subscribed}
	IsAlertApiSubscribed interface{} `field:"optional" json:"isAlertApiSubscribed" yaml:"isAlertApiSubscribed"`
}

