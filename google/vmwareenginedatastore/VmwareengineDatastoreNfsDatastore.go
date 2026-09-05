// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vmwareenginedatastore


type VmwareengineDatastoreNfsDatastore struct {
	// google_file_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vmwareengine_datastore#google_file_service VmwareengineDatastore#google_file_service}
	GoogleFileService *VmwareengineDatastoreNfsDatastoreGoogleFileService `field:"optional" json:"googleFileService" yaml:"googleFileService"`
	// third_party_file_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vmwareengine_datastore#third_party_file_service VmwareengineDatastore#third_party_file_service}
	ThirdPartyFileService *VmwareengineDatastoreNfsDatastoreThirdPartyFileService `field:"optional" json:"thirdPartyFileService" yaml:"thirdPartyFileService"`
}

