---
subcategory: "Storage Gateway"
layout: "aws"
page_title: "AWS: aws_storagegateway_cache"
description: |-
  Manages an AWS Storage Gateway cache
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_storagegateway_cache

Manages an AWS Storage Gateway cache.

~> **NOTE:** The Storage Gateway API provides no method to remove a cache disk. Destroying this Terraform resource does not perform any Storage Gateway actions.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { StoragegatewayCache } from "./.gen/providers/aws/storagegateway-cache";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new StoragegatewayCache(this, "example", {
      diskId: Token.asString(dataAwsStoragegatewayLocalDiskExample.id),
      gatewayArn: Token.asString(awsStoragegatewayGatewayExample.arn),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `diskId` - (Required) Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
* `gatewayArn` - (Required) The Amazon Resource Name (ARN) of the gateway.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Combined gateway Amazon Resource Name (ARN) and local disk identifier.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_storagegateway_cache` using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`). For example:

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

Using `terraform import`, import `aws_storagegateway_cache` using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`). For example:

```console
% terraform import aws_storagegateway_cache.example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
```

<!-- cache-key: cdktf-0.20.0 input-71c425b426e8be6b2795e6720d1217aa0af1c682498abb9c938909327c75ac67 -->