---
page_title: "btp_subaccount_trust_configuration Data Source - terraform-provider-btp"
subcategory: ""
description: |-
  Gets details about a trust configuration.
  Tip:
  You must be viewer or administrator of the subaccount.
  Further documentation:
  https://help.sap.com/docs/btp/sap-btp-neo-environment/platform-identity-provider
---

# btp_subaccount_trust_configuration (Data Source)

Gets details about a trust configuration.

__Tip:__
You must be viewer or administrator of the subaccount.

__Further documentation:__
<https://help.sap.com/docs/btp/sap-btp-neo-environment/platform-identity-provider>

## Example Usage

```terraform
# default identity provider
data "btp_subaccount_trust_configuration" "default" {
  subaccount_id = "6aa64c2f-38c1-49a9-b2e8-cf9fea769b7f"
  origin        = "sap.default"
}

# custom identity provider
data "btp_subaccount_trust_configuration" "custom" {
  subaccount_id = "6aa64c2f-38c1-49a9-b2e8-cf9fea769b7f"
  origin        = "terraformint-platform"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `origin` (String) The origin of the identity provider.
- `subaccount_id` (String) The ID of the subaccount.

### Read-Only

- `description` (String) The description of the trust configuration.
- `id` (String) The ID of the trust configuration.
- `identity_provider` (String) The name of the identity provider.
- `name` (String) The name of the trust configuration.
- `protocol` (String) The protocol used to establish trust with the identity provider.
- `read_only` (Boolean) Shows whether the trust configuration can be modified.
- `status` (String) Shows whether the identity provider is currently active or not.
- `type` (String) The trust type.