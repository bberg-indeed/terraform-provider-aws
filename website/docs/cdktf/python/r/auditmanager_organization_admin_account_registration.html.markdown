---
subcategory: "Audit Manager"
layout: "aws"
page_title: "AWS: aws_auditmanager_organization_admin_account_registration"
description: |-
  Terraform resource for managing AWS Audit Manager Organization Admin Account Registration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_auditmanager_organization_admin_account_registration

Terraform resource for managing AWS Audit Manager Organization Admin Account Registration.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.auditmanager_organization_admin_account_registration import AuditmanagerOrganizationAdminAccountRegistration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AuditmanagerOrganizationAdminAccountRegistration(self, "example",
            admin_account_id="012345678901"
        )
```

## Argument Reference

The following arguments are required:

* `admin_account_id` - (Required) Identifier for the organization administrator account.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Identifier for the organization administrator account.
* `organization_id` - Identifier for the organization.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Audit Manager Organization Admin Account Registration using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
```

Using `terraform import`, import Audit Manager Organization Admin Account Registration using the `id`. For example:

```console
% terraform import aws_auditmanager_organization_admin_account_registration.example 012345678901 
```

<!-- cache-key: cdktf-0.20.0 input-73a8d01ece65d6322ccb932f7e3936fdefadced313d0732bd4dd6da49b4ca658 -->