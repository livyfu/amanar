// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    amanar, err := UnmarshalAmanar(bytes)
//    bytes, err = amanar.Marshal()

package amanar

import "encoding/json"

func UnmarshalAmanar(data []byte) (Amanar, error) {
	var r Amanar
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Amanar) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Amanar struct {
	AmanarConfiguration []AmanarConfiguration `json:"amanar_configuration" yaml:"amanar_configuration"`
}

type AmanarConfiguration struct {
	Constant              *Constant            `json:"constant,omitempty" yaml:"constant,omitempty"`                // Amanar is able to render constants, which are pieces of information that do not depend on; Vault. This allows Amanar to be used to configure local static credentials in addition to; Vault-derived ones.
	VaultAddress          *string              `json:"vault_address,omitempty" yaml:"vault_address,omitempty"`           // The address to a particular vault. Vault addresses usually differ for different; environments. For example, we may have one vault address for production and another for; staging.
	VaultAuth             *VaultAuth           `json:"vault_auth,omitempty" yaml:"vault_auth,omitempty"`              // Vault auth method.
	VaultConfiguration    []VaultConfiguration `json:"vault_configuration,omitempty" yaml:"vault_configuration,omitempty"`     
	VaultDBUsernamePrefix *string              `json:"vault_db_username_prefix,omitempty" yaml:"vault_db_username_prefix,omitempty"`// This will be used as the DisplayName of the Vault token, then further be used in the; generation of the temporary DB username depend on how you defined the username_template.; Currently, this is only supported by AWS IAM auth method.
}

// Amanar is able to render constants, which are pieces of information that do not depend on
// Vault. This allows Amanar to be used to configure local static credentials in addition to
// Vault-derived ones.
type Constant struct {
	Template     *string `json:"template,omitempty" yaml:"template,omitempty"`     // A constant Go template string that will be rendered.
	TemplatePath *string `json:"template_path,omitempty" yaml:"template_path,omitempty"`// A path to a Go template file that will be rendered.
}

// A list of vault roles and paths and configuration options for output to data sources
// within a particular vault environment.
type VaultConfiguration struct {
	Configurables Configurables `json:"configurables" yaml:"configurables"`
	VaultPath     string        `json:"vault_path" yaml:"vault_path"`   // The path representing the datastore in the Vault. This is equivalent to $VAULT_PATH in; the CLI command `vault read $VAULT_PATH/creds/$VAULT_ROLE`.
	VaultRole     string        `json:"vault_role" yaml:"vault_role"`   // The role representing the permissions that are sought to the Vault datastore. This is; equivalent to $VAULT_ROLE in the CLI command `vault read $VAULT_PATH/creds/$VAULT_ROLE`.
}

type Configurables struct {
	IntellijDatasources       []IntellijDatasource       `json:"intellij_datasources,omitempty" yaml:"intellij_datasources,omitempty"`       // Allows IntelliJ datasource usernames and passwords to be changed. Most useful for; DataGrip and databases within IntelliJ Ultimate.
	IntellijRunConfigurations []IntellijRunConfiguration `json:"intellij_run_configurations,omitempty" yaml:"intellij_run_configurations,omitempty"`// Allows changes to database access credentials within IntelliJ run configurations.
	JSONDatasources           []JSONDatasource           `json:"json_datasources,omitempty" yaml:"json_datasources,omitempty"`           // Allows a JSON file to be generated containing usernames and passwords.
	PosticoDatasources        []PosticoDatasource        `json:"postico_datasources,omitempty" yaml:"postico_datasources,omitempty"`        // Allows changes to database access credentials stored in a Postico SQLite database.
	Querious2Datasources      []Querious2Datasource      `json:"querious2_datasources,omitempty" yaml:"querious2_datasources,omitempty"`      // Allows changes to database access credentials stored in a Querious 2 SQLite database.
	SequelAceDatasources      []SequelAceDatasource      `json:"sequel_ace_datasources,omitempty" yaml:"sequel_ace_datasources,omitempty"`     // Allows changes to database access credentials for Sequel Ace plists.
	SequelProDatasources      []SequelProDatasource      `json:"sequel_pro_datasources,omitempty" yaml:"sequel_pro_datasources,omitempty"`     // Allows changes to database access credentials for Sequel Pro plists.
	ShellDatasources          []ShellDatasource          `json:"shell_datasources,omitempty" yaml:"shell_datasources,omitempty"`          // Allows a file to be generated in a shell script that contains exports of environment; variables containing the new credentials.
	TemplateDatasources       []TemplateDatasource       `json:"template_datasources,omitempty" yaml:"template_datasources,omitempty"`       // Fills credentials into a provided Go template string or template file and prints the; result to stdout. The Credentials object is set as the dot context and its fields; (.Username and .Password) are available. Unlike the other datasources, templates are; generated anew each time and do not attempt to find previously existing keys to modify.; Go templates are documented here: https://golang.org/pkg/text/template/
}

type IntellijDatasource struct {
	DatabaseUUID       string `json:"database_uuid" yaml:"database_uuid"`       // The IntelliJ UUID for the database you want to update. You can find this by examining the; dataSources.local.xml file.
	DatasourceFilePath string `json:"datasource_file_path" yaml:"datasource_file_path"`// The path to IntelliJ data sources file. The file is typically called; dataSources.local.xml.
}

type IntellijRunConfiguration struct {
	DatabaseHost                string `json:"database_host" yaml:"database_host"`                 // The username and password for the URL will only be updated if the host of URL in the; environment variable matches this string.
	EnvironmentVariable         string `json:"environment_variable" yaml:"environment_variable"`          // The environment variable in the run configuration under which the database connection
	RunConfigurationsFolderPath string `json:"run_configurations_folder_path" yaml:"run_configurations_folder_path"`// A directory containing all IntelliJ run configurations to be examined. Usually located in; .idea/runConfigurations. Run configurations may need to be shared before becoming visible; in this folder.
}

type JSONDatasource struct {
	Filepath   string `json:"filepath" yaml:"filepath"`  // The path the JSON file should be generated to.
	Identifier string `json:"identifier" yaml:"identifier"`// The name of this vault role and vault path pair to be used as an identifier for this JSON; object.
}

type PosticoDatasource struct {
	DatabaseUUID      string `json:"database_uuid" yaml:"database_uuid"`      // The unique identifier for the Postico database to update. Can be found by looking in the; SQLite database.
	PosticoSqlitePath string `json:"postico_sqlite_path" yaml:"postico_sqlite_path"`// Path to the SQLite database in which Postico stores its data. The file is typically; called ConnectionFavorites.db
}

type Querious2Datasource struct {
	DatabaseUUID        string `json:"database_uuid" yaml:"database_uuid"`        // The unique identifier for the Querious database to update. Can be found by looking in the; SQLite database.
	Querious2SqlitePath string `json:"querious2_sqlite_path" yaml:"querious2_sqlite_path"`// Path to the SQLite database in which Querious 2 stores its data. The file is typically; called Connections.sqlite.
}

type SequelAceDatasource struct {
	DatabaseUUID       string `json:"database_uuid" yaml:"database_uuid"`        // The unique identifier for the Sequel Ace database to update. Can be found by looking in; the plist.
	SequelAcePlistPath string `json:"sequel_ace_plist_path" yaml:"sequel_ace_plist_path"`// Path to the plist in which Sequel Ace stores its data. The file is typically called; Favorites.plist
}

type SequelProDatasource struct {
	DatabaseUUID       string `json:"database_uuid" yaml:"database_uuid"`        // The unique identifier for the Sequel Pro database to update. Can be found by looking in; the plist.
	SequelProPlistPath string `json:"sequel_pro_plist_path" yaml:"sequel_pro_plist_path"`// Path to the plist in which Sequel Pro stores its data. The file is typically called; Favorites.plist
}

type ShellDatasource struct {
	Filepath         string `json:"filepath" yaml:"filepath"`         // The path the shell script should be generated to.
	PasswordVariable string `json:"password_variable" yaml:"password_variable"`// The name of the environment variable that should contain the password
	UsernameVariable string `json:"username_variable" yaml:"username_variable"`// The name of the environment variable that should contain the username
}

type TemplateDatasource struct {
	Template     *string `json:"template,omitempty" yaml:"template,omitempty"`     // A Go template string that will be filled in with credentials.
	TemplatePath *string `json:"template_path,omitempty" yaml:"template_path,omitempty"`// The path to the Go template file that will be filled in with credentials.
}

// Vault auth method.
type VaultAuth string
const (
	AwsIam VaultAuth = "aws_iam"
	Github VaultAuth = "github"
	Token VaultAuth = "token"
)
