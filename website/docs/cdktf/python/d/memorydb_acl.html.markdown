---
subcategory: "MemoryDB for Redis"
layout: "aws"
page_title: "AWS: aws_memorydb_acl"
description: |-
  Provides information about a MemoryDB ACL.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_memorydb_acl

Provides information about a MemoryDB ACL.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_memorydb_acl import DataAwsMemorydbAcl
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsMemorydbAcl(self, "example",
            name="my-acl"
        )
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the ACL.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - Name of the ACL.
* `arn` - ARN of the ACL.
* `minimum_engine_version` - The minimum engine version supported by the ACL.
* `tags` - Map of tags assigned to the ACL.
* `user_names` - Set of MemoryDB user names included in this ACL.

<!-- cache-key: cdktf-0.20.0 input-dae65be976707896a0ccf5394e359974398040d0921bfad084e1f4c241a2a529 -->