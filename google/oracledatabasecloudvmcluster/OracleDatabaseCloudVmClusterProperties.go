// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudvmcluster


type OracleDatabaseCloudVmClusterProperties struct {
	// Number of enabled CPU cores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#cpu_core_count OracleDatabaseCloudVmCluster#cpu_core_count}
	CpuCoreCount *float64 `field:"required" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// License type of VM Cluster.   Possible values:  LICENSE_TYPE_UNSPECIFIED LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#license_type OracleDatabaseCloudVmCluster#license_type}
	LicenseType *string `field:"required" json:"licenseType" yaml:"licenseType"`
	// OCI Cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#cluster_name OracleDatabaseCloudVmCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The data disk group size to be allocated in TBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#data_storage_size_tb OracleDatabaseCloudVmCluster#data_storage_size_tb}
	DataStorageSizeTb *float64 `field:"optional" json:"dataStorageSizeTb" yaml:"dataStorageSizeTb"`
	// Local storage per VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#db_node_storage_size_gb OracleDatabaseCloudVmCluster#db_node_storage_size_gb}
	DbNodeStorageSizeGb *float64 `field:"optional" json:"dbNodeStorageSizeGb" yaml:"dbNodeStorageSizeGb"`
	// OCID of database servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#db_server_ocids OracleDatabaseCloudVmCluster#db_server_ocids}
	DbServerOcids *[]*string `field:"optional" json:"dbServerOcids" yaml:"dbServerOcids"`
	// diagnostics_data_collection_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#diagnostics_data_collection_options OracleDatabaseCloudVmCluster#diagnostics_data_collection_options}
	DiagnosticsDataCollectionOptions *OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions `field:"optional" json:"diagnosticsDataCollectionOptions" yaml:"diagnosticsDataCollectionOptions"`
	// The type of redundancy.   Possible values:  DISK_REDUNDANCY_UNSPECIFIED HIGH NORMAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#disk_redundancy OracleDatabaseCloudVmCluster#disk_redundancy}
	DiskRedundancy *string `field:"optional" json:"diskRedundancy" yaml:"diskRedundancy"`
	// Grid Infrastructure Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#gi_version OracleDatabaseCloudVmCluster#gi_version}
	GiVersion *string `field:"optional" json:"giVersion" yaml:"giVersion"`
	// Prefix for VM cluster host names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#hostname_prefix OracleDatabaseCloudVmCluster#hostname_prefix}
	HostnamePrefix *string `field:"optional" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// Use local backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#local_backup_enabled OracleDatabaseCloudVmCluster#local_backup_enabled}
	LocalBackupEnabled interface{} `field:"optional" json:"localBackupEnabled" yaml:"localBackupEnabled"`
	// Memory allocated in GBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#memory_size_gb OracleDatabaseCloudVmCluster#memory_size_gb}
	MemorySizeGb *float64 `field:"optional" json:"memorySizeGb" yaml:"memorySizeGb"`
	// Number of database servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#node_count OracleDatabaseCloudVmCluster#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// OCPU count per VM. Minimum is 0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#ocpu_count OracleDatabaseCloudVmCluster#ocpu_count}
	OcpuCount *float64 `field:"optional" json:"ocpuCount" yaml:"ocpuCount"`
	// SCAN listener port - TCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#scan_listener_port_tcp OracleDatabaseCloudVmCluster#scan_listener_port_tcp}
	ScanListenerPortTcp *float64 `field:"optional" json:"scanListenerPortTcp" yaml:"scanListenerPortTcp"`
	// SCAN listener port - TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#scan_listener_port_tcp_ssl OracleDatabaseCloudVmCluster#scan_listener_port_tcp_ssl}
	ScanListenerPortTcpSsl *float64 `field:"optional" json:"scanListenerPortTcpSsl" yaml:"scanListenerPortTcpSsl"`
	// Use exadata sparse snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#sparse_diskgroup_enabled OracleDatabaseCloudVmCluster#sparse_diskgroup_enabled}
	SparseDiskgroupEnabled interface{} `field:"optional" json:"sparseDiskgroupEnabled" yaml:"sparseDiskgroupEnabled"`
	// SSH public keys to be stored with cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#ssh_public_keys OracleDatabaseCloudVmCluster#ssh_public_keys}
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#time_zone OracleDatabaseCloudVmCluster#time_zone}
	TimeZone *OracleDatabaseCloudVmClusterPropertiesTimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
	// Specifies whether VM backups are stored on local DB server storage or Exascale storage. Possible values: VM_BACKUP_STORAGE_TYPE_UNSPECIFIED VM_BACKUP_STORAGE_TYPE_LOCAL VM_BACKUP_STORAGE_TYPE_EXASCALE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#vm_backup_storage_type OracleDatabaseCloudVmCluster#vm_backup_storage_type}
	VmBackupStorageType *string `field:"optional" json:"vmBackupStorageType" yaml:"vmBackupStorageType"`
	// Specifies whether VM file system storage / VM images are stored on local DB server storage or Exascale storage.
	//
	// Possible values:
	// VM_FILE_SYSTEM_STORAGE_TYPE_UNSPECIFIED
	// VM_FILE_SYSTEM_STORAGE_TYPE_LOCAL
	// VM_FILE_SYSTEM_STORAGE_TYPE_EXASCALE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/oracle_database_cloud_vm_cluster#vm_file_system_storage_type OracleDatabaseCloudVmCluster#vm_file_system_storage_type}
	VmFileSystemStorageType *string `field:"optional" json:"vmFileSystemStorageType" yaml:"vmFileSystemStorageType"`
}

