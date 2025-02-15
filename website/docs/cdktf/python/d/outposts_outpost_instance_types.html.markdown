---
subcategory: "Outposts"
layout: "aws"
page_title: "AWS: aws_outposts_outpost_instance_types"
description: |-
  Information about Outpost Instance Types.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_outposts_outpost_instance_types

Information about Outposts Instance Types.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_outposts_outpost_instance_types import DataAwsOutpostsOutpostInstanceTypes
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsOutpostsOutpostInstanceTypes(self, "example",
            arn=Token.as_string(data_aws_outposts_outpost_example.arn)
        )
```

## Argument Reference

The following arguments are required:

* `arn` - (Required) Outpost ARN.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `instance_types` - Set of instance types.

<!-- cache-key: cdktf-0.20.0 input-3dcd4d42109b60dec29ef5ce3754b18b70307fc52baa1c511a6b236d274c3d89 -->