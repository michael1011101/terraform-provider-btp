---
page_title: "btp_globalaccount_resource_provider Resource - terraform-provider-btp"
subcategory: ""
description: |-
  Create a resource provider instance to allow your global account to connect to your provider account on a non-SAP cloud vendor. Through this channel, you can then consume remote service resources that you already own and which are supported by SAP BTP.
  For example, if you are subscribed to Amazon Web Services (AWS) and have already purchased services, such as PostgreSQL, you can register the vendor as a resource provider in SAP BTP and consume this service across your subaccounts together with other services offered by SAP.
  The use of this functionality is subject to the availability of the supported non-SAP cloud vendors in your country or region.
  Tips
  * You must be assigned to the global account admin role.
  * You can create more than one instance of a given resource provider, each with its unique configuration properties. In such cases, the display name and technical name should be descriptive enough so that you and developers can easily differentiate between each instance.
  * After you configure a new resource provider instance, its supported services are added as entitlements in your global account.
  Further documentation
  https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/e2c250dc5abd468a81f4f619206157a2.html
---

# btp_globalaccount_resource_provider (Resource)

Create a resource provider instance to allow your global account to connect to your provider account on a non-SAP cloud vendor. Through this channel, you can then consume remote service resources that you already own and which are supported by SAP BTP.
For example, if you are subscribed to Amazon Web Services (AWS) and have already purchased services, such as PostgreSQL, you can register the vendor as a resource provider in SAP BTP and consume this service across your subaccounts together with other services offered by SAP.

The use of this functionality is subject to the availability of the supported non-SAP cloud vendors in your country or region.

__Tips__
* You must be assigned to the global account admin role.
* You can create more than one instance of a given resource provider, each with its unique configuration properties. In such cases, the display name and technical name should be descriptive enough so that you and developers can easily differentiate between each instance.
* After you configure a new resource provider instance, its supported services are added as entitlements in your global account.

__Further documentation__
https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/e2c250dc5abd468a81f4f619206157a2.html

## Example Usage

```terraform
# register a AZURE project as resource provider
resource "btp_globalaccount_resource_provider" "azure" {
  id                = "my_azure_provider"
  resource_provider = "AZURE"
  parameters = jsonencode({
    region              = "westeurope"
    client_id           = "AZURECLIENTID"
    client_secret       = "AZURECLIENTSECRET"
    tenant_id           = "42x7676x-f455-423x-82x6-xx2d99791xx7"
    subscription_id     = "x1x9567x-8560-44xx-x4fx-741xx0x08x58"
    resource_group_name = "rg-landscape-azure-example"
  })
}

# register an AWS account as resource provider
resource "btp_globalaccount_resource_provider" "aws" {
  id                = "my_aws_provider"
  resource_provider = "AWS"
  parameters = jsonencode({
    access_key_id     = "AWSACCESSKEY"
    secret_access_key = "AWSSECRETKEY"
    vpc_id            = "vpc-test"
    region            = "eu-central-1"
  })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Unique technical name of the resource provider.
- `parameters` (String, Sensitive) Any relevant information about the resource provider that is not provided by other parameter values.
- `resource_provider` (String) Provider of the requested resource. For example: AWS, AZURE.

### Read-Only

- `description` (String) Description of the resource provider.
- `display_name` (String) Descriptive name of the resource provider.

## Import

Import is supported using the following syntax:

```terraform
# terraform import btp_globalaccount_resource_provider.<resource_name> <resource_provider>,<unique_technical_name>

terraform import btp_globalaccount_resource_provider.azure AZURE,my_azure_provider
```