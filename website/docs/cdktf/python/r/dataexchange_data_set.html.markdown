---
subcategory: "Data Exchange"
layout: "aws"
page_title: "AWS: aws_dataexchange_data_set"
description: |-
  Provides a DataExchange DataSet
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dataexchange_data_set

Provides a resource to manage AWS Data Exchange DataSets.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.dataexchange_data_set import DataexchangeDataSet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataexchangeDataSet(self, "example",
            asset_type="S3_SNAPSHOT",
            description="example",
            name="example"
        )
```

## Argument Reference

* `asset_type` - (Required) The type of asset that is added to a data set. Valid values are: `S3_SNAPSHOT`, `REDSHIFT_DATA_SHARE`, and `API_GATEWAY_API`.
* `description` - (Required) A description for the data set.
* `name` - (Required) The name of the data set.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Id of the data set.
* `arn` - The Amazon Resource Name of this data set.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DataExchange DataSets using their ARN. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
```

Using `terraform import`, import DataExchange DataSets using their ARN. For example:

```console
% terraform import aws_dataexchange_data_set.example arn:aws:dataexchange:us-west-2:123456789012:data-sets/4fa784c7-ccb4-4dbf-ba4f-02198320daa1
```

<!-- cache-key: cdktf-0.20.0 input-bcb32869378838571a8ae492463bc57c036b3f3e8528a58d44a0847e4f76852a -->