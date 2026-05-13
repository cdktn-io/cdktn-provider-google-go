// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeCacheParametersCacheConfig struct {
	// Optional. Flag indicating whether a CIFS change notification is enabled for the FlexCache volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/netapp_volume#cifs_change_notify_enabled NetappVolume#cifs_change_notify_enabled}
	CifsChangeNotifyEnabled interface{} `field:"optional" json:"cifsChangeNotifyEnabled" yaml:"cifsChangeNotifyEnabled"`
}

