---
page_title: "btp_directory_roles Data Source - terraform-provider-btp"
subcategory: ""
description: |-
  List all roles.
  Further documentation
  https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/0039cf082d3d43eba9200fe15647922a.html
---

# btp_directory_roles (Data Source)

List all roles.

__Further documentation__
https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/0039cf082d3d43eba9200fe15647922a.html

## Example Usage

```terraform
data "btp_directory_roles" "all" {
  directory_id = "dd005d8b-1fee-4e6b-b6ff-cb9a197b7fe0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `directory_id` (String) The ID of the directory.

### Read-Only

- `id` (String, Deprecated) The ID of the directory.
- `values` (Attributes List) (see [below for nested schema](#nestedatt--values))

<a id="nestedatt--values"></a>
### Nested Schema for `values`

Read-Only:

- `app_id` (String) The ID of the xsuaa application.
- `description` (String) The description of the role.
- `name` (String) The name of the role.
- `read_only` (Boolean) Whether the role can be modified or not.
- `role_template_name` (String) The name of the role template.
- `scopes` (Attributes List) Scopes available with this role. (see [below for nested schema](#nestedatt--values--scopes))

<a id="nestedatt--values--scopes"></a>
### Nested Schema for `values.scopes`

Read-Only:

- `custom_grant_as_authority_to_apps` (Set of String)
- `custom_granted_apps` (Set of String)
- `description` (String) The description of the scope.
- `grant_as_authority_to_apps` (Set of String)
- `granted_apps` (Set of String)
- `name` (String) The name of the scope.