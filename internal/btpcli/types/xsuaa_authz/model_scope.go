/*
 * Authorization
 *
 * Provides functions to administrate the Authorization and Trust Management service (XSUAA) of SAP BTP, Cloud Foundry environment. You can manage service instances of the Authorization and Trust Management service. You can also manage roles, role templates, and role collections of your subaccount.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package xsuaa_authz

type Scope struct {
	CustomGrantAsAuthorityToApps []string `json:"custom-grant-as-authority-to-apps,omitempty"`
	CustomGrantedApps            []string `json:"custom-granted-apps,omitempty"`
	Description                  string   `json:"description,omitempty"`
	GrantAsAuthorityToApps       []string `json:"grant-as-authority-to-apps,omitempty"`
	GrantedApps                  []string `json:"granted-apps,omitempty"`
	// The name of the application and scope as defined in the application security descriptor xs-security.json. The name has a maximum length of 193 characters, including the fully qualified application name. The fully qualified scope name starts with the application ID followed by an optional number of components and finally the scope, each separated by a period (.). For example: service-manager!b105.entitlement.notify. Only the following characters are allowed: alphanumeric characters (aA-zZ) and (0-9), hyphen (-), underscore (_), forward slash (/), backslash (\\), and colon (:).
	Name string `json:"name,omitempty"`
}