---
subcategory: "ELB Classic"
layout: "aws"
page_title: "AWS: aws_elb_hosted_zone_id"
description: |-
  Get AWS Elastic Load Balancing Hosted Zone Id
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_elb_hosted_zone_id

Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
in a given region for the purpose of using in an AWS Route53 Alias.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_elb_hosted_zone_id import DataAwsElbHostedZoneId
from imports.aws.route53_record import Route53Record
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        main = DataAwsElbHostedZoneId(self, "main")
        Route53Record(self, "www",
            alias=Route53RecordAlias(
                evaluate_target_health=True,
                name=Token.as_string(aws_elb_main.dns_name),
                zone_id=Token.as_string(main.id)
            ),
            name="example.com",
            type="A",
            zone_id=primary.zone_id
        )
```

## Argument Reference

* `region` - (Optional) Name of the region whose AWS ELB HostedZoneId is desired.
  Defaults to the region from the AWS provider configuration.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - ID of the AWS ELB HostedZoneId in the selected region.

<!-- cache-key: cdktf-0.20.0 input-c2a71897abaf8ebad957ebc450a6e6db3c22c6c5e61e688eccadc3e3e3d48143 -->