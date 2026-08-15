// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionProperties struct {
	// The connection type. Possible values: GOLDENGATE KAFKA KAFKA_SCHEMA_REGISTRY MYSQL JAVA_MESSAGE_SERVICE MICROSOFT_SQLSERVER OCI_OBJECT_STORAGE ORACLE AZURE_DATA_LAKE_STORAGE POSTGRESQL AZURE_SYNAPSE_ANALYTICS SNOWFLAKE AMAZON_S3 HDFS ORACLE_AI_DATA_PLATFORM ORACLE_NOSQL MONGODB AMAZON_KINESIS AMAZON_REDSHIFT DB2 REDIS ELASTICSEARCH GENERIC GOOGLE_CLOUD_STORAGE GOOGLE_BIGQUERY DATABRICKS GOOGLE_PUBSUB MICROSOFT_FABRIC ICEBERG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#connection_type OracleDatabaseGoldengateConnection#connection_type}
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// An object's Display Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#display_name OracleDatabaseGoldengateConnection#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// amazon_kinesis_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#amazon_kinesis_connection_properties OracleDatabaseGoldengateConnection#amazon_kinesis_connection_properties}
	AmazonKinesisConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties `field:"optional" json:"amazonKinesisConnectionProperties" yaml:"amazonKinesisConnectionProperties"`
	// amazon_redshift_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#amazon_redshift_connection_properties OracleDatabaseGoldengateConnection#amazon_redshift_connection_properties}
	AmazonRedshiftConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties `field:"optional" json:"amazonRedshiftConnectionProperties" yaml:"amazonRedshiftConnectionProperties"`
	// amazon_s3_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#amazon_s3_connection_properties OracleDatabaseGoldengateConnection#amazon_s3_connection_properties}
	AmazonS3ConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties `field:"optional" json:"amazonS3ConnectionProperties" yaml:"amazonS3ConnectionProperties"`
	// azure_data_lake_storage_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#azure_data_lake_storage_connection_properties OracleDatabaseGoldengateConnection#azure_data_lake_storage_connection_properties}
	AzureDataLakeStorageConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties `field:"optional" json:"azureDataLakeStorageConnectionProperties" yaml:"azureDataLakeStorageConnectionProperties"`
	// azure_synapse_analytics_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#azure_synapse_analytics_connection_properties OracleDatabaseGoldengateConnection#azure_synapse_analytics_connection_properties}
	AzureSynapseAnalyticsConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties `field:"optional" json:"azureSynapseAnalyticsConnectionProperties" yaml:"azureSynapseAnalyticsConnectionProperties"`
	// databricks_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#databricks_connection_properties OracleDatabaseGoldengateConnection#databricks_connection_properties}
	DatabricksConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties `field:"optional" json:"databricksConnectionProperties" yaml:"databricksConnectionProperties"`
	// db2_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#db2_connection_properties OracleDatabaseGoldengateConnection#db2_connection_properties}
	Db2ConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties `field:"optional" json:"db2ConnectionProperties" yaml:"db2ConnectionProperties"`
	// Metadata about this specific object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#description OracleDatabaseGoldengateConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// elasticsearch_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#elasticsearch_connection_properties OracleDatabaseGoldengateConnection#elasticsearch_connection_properties}
	ElasticsearchConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties `field:"optional" json:"elasticsearchConnectionProperties" yaml:"elasticsearchConnectionProperties"`
	// generic_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#generic_connection_properties OracleDatabaseGoldengateConnection#generic_connection_properties}
	GenericConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties `field:"optional" json:"genericConnectionProperties" yaml:"genericConnectionProperties"`
	// goldengate_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#goldengate_connection_properties OracleDatabaseGoldengateConnection#goldengate_connection_properties}
	GoldengateConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties `field:"optional" json:"goldengateConnectionProperties" yaml:"goldengateConnectionProperties"`
	// google_big_query_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#google_big_query_connection_properties OracleDatabaseGoldengateConnection#google_big_query_connection_properties}
	GoogleBigQueryConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties `field:"optional" json:"googleBigQueryConnectionProperties" yaml:"googleBigQueryConnectionProperties"`
	// google_cloud_storage_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#google_cloud_storage_connection_properties OracleDatabaseGoldengateConnection#google_cloud_storage_connection_properties}
	GoogleCloudStorageConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties `field:"optional" json:"googleCloudStorageConnectionProperties" yaml:"googleCloudStorageConnectionProperties"`
	// google_pubsub_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#google_pubsub_connection_properties OracleDatabaseGoldengateConnection#google_pubsub_connection_properties}
	GooglePubsubConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties `field:"optional" json:"googlePubsubConnectionProperties" yaml:"googlePubsubConnectionProperties"`
	// hdfs_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#hdfs_connection_properties OracleDatabaseGoldengateConnection#hdfs_connection_properties}
	HdfsConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties `field:"optional" json:"hdfsConnectionProperties" yaml:"hdfsConnectionProperties"`
	// iceberg_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#iceberg_connection_properties OracleDatabaseGoldengateConnection#iceberg_connection_properties}
	IcebergConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties `field:"optional" json:"icebergConnectionProperties" yaml:"icebergConnectionProperties"`
	// java_message_service_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#java_message_service_connection_properties OracleDatabaseGoldengateConnection#java_message_service_connection_properties}
	JavaMessageServiceConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties `field:"optional" json:"javaMessageServiceConnectionProperties" yaml:"javaMessageServiceConnectionProperties"`
	// kafka_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#kafka_connection_properties OracleDatabaseGoldengateConnection#kafka_connection_properties}
	KafkaConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties `field:"optional" json:"kafkaConnectionProperties" yaml:"kafkaConnectionProperties"`
	// kafka_schema_registry_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#kafka_schema_registry_connection_properties OracleDatabaseGoldengateConnection#kafka_schema_registry_connection_properties}
	KafkaSchemaRegistryConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties `field:"optional" json:"kafkaSchemaRegistryConnectionProperties" yaml:"kafkaSchemaRegistryConnectionProperties"`
	// microsoft_fabric_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#microsoft_fabric_connection_properties OracleDatabaseGoldengateConnection#microsoft_fabric_connection_properties}
	MicrosoftFabricConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties `field:"optional" json:"microsoftFabricConnectionProperties" yaml:"microsoftFabricConnectionProperties"`
	// microsoft_sqlserver_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#microsoft_sqlserver_connection_properties OracleDatabaseGoldengateConnection#microsoft_sqlserver_connection_properties}
	MicrosoftSqlserverConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties `field:"optional" json:"microsoftSqlserverConnectionProperties" yaml:"microsoftSqlserverConnectionProperties"`
	// mongodb_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#mongodb_connection_properties OracleDatabaseGoldengateConnection#mongodb_connection_properties}
	MongodbConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties `field:"optional" json:"mongodbConnectionProperties" yaml:"mongodbConnectionProperties"`
	// mysql_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#mysql_connection_properties OracleDatabaseGoldengateConnection#mysql_connection_properties}
	MysqlConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties `field:"optional" json:"mysqlConnectionProperties" yaml:"mysqlConnectionProperties"`
	// oci_object_storage_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#oci_object_storage_connection_properties OracleDatabaseGoldengateConnection#oci_object_storage_connection_properties}
	OciObjectStorageConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties `field:"optional" json:"ociObjectStorageConnectionProperties" yaml:"ociObjectStorageConnectionProperties"`
	// oracle_ai_data_platform_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#oracle_ai_data_platform_connection_properties OracleDatabaseGoldengateConnection#oracle_ai_data_platform_connection_properties}
	OracleAiDataPlatformConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties `field:"optional" json:"oracleAiDataPlatformConnectionProperties" yaml:"oracleAiDataPlatformConnectionProperties"`
	// oracle_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#oracle_connection_properties OracleDatabaseGoldengateConnection#oracle_connection_properties}
	OracleConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties `field:"optional" json:"oracleConnectionProperties" yaml:"oracleConnectionProperties"`
	// oracle_nosql_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#oracle_nosql_connection_properties OracleDatabaseGoldengateConnection#oracle_nosql_connection_properties}
	OracleNosqlConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties `field:"optional" json:"oracleNosqlConnectionProperties" yaml:"oracleNosqlConnectionProperties"`
	// postgresql_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#postgresql_connection_properties OracleDatabaseGoldengateConnection#postgresql_connection_properties}
	PostgresqlConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties `field:"optional" json:"postgresqlConnectionProperties" yaml:"postgresqlConnectionProperties"`
	// redis_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#redis_connection_properties OracleDatabaseGoldengateConnection#redis_connection_properties}
	RedisConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties `field:"optional" json:"redisConnectionProperties" yaml:"redisConnectionProperties"`
	// The routing method for the GoldengateConnection. Possible values: SHARED_DEPLOYMENT_ENDPOINT DEDICATED_ENDPOINT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#routing_method OracleDatabaseGoldengateConnection#routing_method}
	RoutingMethod *string `field:"optional" json:"routingMethod" yaml:"routingMethod"`
	// snowflake_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/oracle_database_goldengate_connection#snowflake_connection_properties OracleDatabaseGoldengateConnection#snowflake_connection_properties}
	SnowflakeConnectionProperties *OracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties `field:"optional" json:"snowflakeConnectionProperties" yaml:"snowflakeConnectionProperties"`
}

