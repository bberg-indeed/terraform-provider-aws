---
subcategory: "WAF Classic Regional"
layout: "aws"
page_title: "AWS: aws_wafregional_ipset"
description: |-
  Provides a AWS WAF Regional IPSet resource for use with ALB.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_wafregional_ipset

Provides a WAF Regional IPSet Resource for use with Application Load Balancer.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { WafregionalIpset } from "./.gen/providers/aws/wafregional-ipset";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new WafregionalIpset(this, "ipset", {
      ipSetDescriptor: [
        {
          type: "IPV4",
          value: "192.0.7.0/24",
        },
        {
          type: "IPV4",
          value: "10.16.16.0/16",
        },
      ],
      name: "tfIPSet",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name or description of the IPSet.
* `ipSetDescriptor` - (Optional) One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.

## Nested Blocks

### `ipSetDescriptor`

#### Arguments

* `type` - (Required) The string like IPV4 or IPV6.
* `value` - (Required) The CIDR notation.

## Remarks

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the WAF IPSet.
* `arn` - The ARN of the WAF IPSet.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import WAF Regional IPSets using their ID. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
  }
}

```

Using `terraform import`, import WAF Regional IPSets using their ID. For example:

```console
% terraform import aws_wafregional_ipset.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```

<!-- cache-key: cdktf-0.20.0 input-bd423cc7eae4bc7869d7e7f89697634be360c63540ee5c25595e89016d52dd6a -->