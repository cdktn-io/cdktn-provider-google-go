// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection


type OracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties struct {
	// Authentication type for Java Message Service. Possible values: NONE BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#authentication_type OracleDatabaseGoldengateConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The Java class implementing javax.jms.ConnectionFactory interface supplied by the JMS provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#connection_factory OracleDatabaseGoldengateConnection#connection_factory}
	ConnectionFactory *string `field:"optional" json:"connectionFactory" yaml:"connectionFactory"`
	// Connection URL of the Java Message Service, specifying the protocol, host, and port. e.g.: 'mq://myjms.host.domain:7676'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#connection_url OracleDatabaseGoldengateConnection#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// The Connection Factory can be looked up using this name. e.g.: 'ConnectionFactory'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#jndi_connection_factory OracleDatabaseGoldengateConnection#jndi_connection_factory}
	JndiConnectionFactory *string `field:"optional" json:"jndiConnectionFactory" yaml:"jndiConnectionFactory"`
	// The implementation of javax.naming.spi.InitialContextFactory interface used to obtain initial naming context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#jndi_initial_context_factory OracleDatabaseGoldengateConnection#jndi_initial_context_factory}
	JndiInitialContextFactory *string `field:"optional" json:"jndiInitialContextFactory" yaml:"jndiInitialContextFactory"`
	// The URL that Java Message Service will use to contact the JNDI provider. e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#jndi_provider_url OracleDatabaseGoldengateConnection#jndi_provider_url}
	JndiProviderUrl *string `field:"optional" json:"jndiProviderUrl" yaml:"jndiProviderUrl"`
	// The password associated to the principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#jndi_security_credentials_secret OracleDatabaseGoldengateConnection#jndi_security_credentials_secret}
	JndiSecurityCredentialsSecret *string `field:"optional" json:"jndiSecurityCredentialsSecret" yaml:"jndiSecurityCredentialsSecret"`
	// Specifies the identity of the principal (user) to be authenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#jndi_security_principal OracleDatabaseGoldengateConnection#jndi_security_principal}
	JndiSecurityPrincipal *string `field:"optional" json:"jndiSecurityPrincipal" yaml:"jndiSecurityPrincipal"`
	// The content of the KeyStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#key_store_file OracleDatabaseGoldengateConnection#key_store_file}
	KeyStoreFile *string `field:"optional" json:"keyStoreFile" yaml:"keyStoreFile"`
	// Input only. The KeyStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#key_store_password OracleDatabaseGoldengateConnection#key_store_password}
	KeyStorePassword *string `field:"optional" json:"keyStorePassword" yaml:"keyStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the KeyStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#key_store_password_secret_version OracleDatabaseGoldengateConnection#key_store_password_secret_version}
	KeyStorePasswordSecretVersion *string `field:"optional" json:"keyStorePasswordSecretVersion" yaml:"keyStorePasswordSecretVersion"`
	// Input only. The password Oracle Goldengate uses to connect the Java Message Service in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#password OracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses to connect the associated Java
	// Message Service.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#password_secret_version OracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// Security protocol for Java Message Service. Possible values: PLAIN TLS MTLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#security_protocol OracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// Input only. The password for the cert inside of the KeyStore in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#ssl_key_password OracleDatabaseGoldengateConnection#ssl_key_password}
	SslKeyPassword *string `field:"optional" json:"sslKeyPassword" yaml:"sslKeyPassword"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password for the cert inside of the KeyStore.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#ssl_key_password_secret_version OracleDatabaseGoldengateConnection#ssl_key_password_secret_version}
	SslKeyPasswordSecretVersion *string `field:"optional" json:"sslKeyPasswordSecretVersion" yaml:"sslKeyPasswordSecretVersion"`
	// The technology type of JavaMessageServiceConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#technology_type OracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The content of the TrustStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#trust_store_file OracleDatabaseGoldengateConnection#trust_store_file}
	TrustStoreFile *string `field:"optional" json:"trustStoreFile" yaml:"trustStoreFile"`
	// Input only. The TrustStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#trust_store_password OracleDatabaseGoldengateConnection#trust_store_password}
	TrustStorePassword *string `field:"optional" json:"trustStorePassword" yaml:"trustStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the TrustStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#trust_store_password_secret_version OracleDatabaseGoldengateConnection#trust_store_password_secret_version}
	TrustStorePasswordSecretVersion *string `field:"optional" json:"trustStorePasswordSecretVersion" yaml:"trustStorePasswordSecretVersion"`
	// If set to true, Java Naming and Directory Interface (JNDI) properties should be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#use_jndi OracleDatabaseGoldengateConnection#use_jndi}
	UseJndi interface{} `field:"optional" json:"useJndi" yaml:"useJndi"`
	// The username Oracle Goldengate uses to connect to the Java Message Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/oracle_database_goldengate_connection#username OracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

