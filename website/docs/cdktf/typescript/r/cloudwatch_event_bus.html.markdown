---
subcategory: "EventBridge"
layout: "aws"
page_title: "AWS: aws_cloudwatch_event_bus"
description: |-
  Provides an EventBridge event bus resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cloudwatch_event_bus

Provides an EventBridge event bus resource.

~> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventBus } from "./.gen/providers/aws/cloudwatch-event-bus";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CloudwatchEventBus(this, "messenger", {
      name: "chat-messages",
    });
  }
}

```

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventBus } from "./.gen/providers/aws/cloudwatch-event-bus";
import { DataAwsCloudwatchEventSource } from "./.gen/providers/aws/data-aws-cloudwatch-event-source";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const examplepartner = new DataAwsCloudwatchEventSource(
      this,
      "examplepartner",
      {
        namePrefix: "aws.partner/examplepartner.com",
      }
    );
    const awsCloudwatchEventBusExamplepartner = new CloudwatchEventBus(
      this,
      "examplepartner_1",
      {
        eventSourceName: Token.asString(examplepartner.name),
        name: Token.asString(examplepartner.name),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsCloudwatchEventBusExamplepartner.overrideLogicalId("examplepartner");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name of the new event bus. The names of custom event buses can't contain the / character. To create a partner event bus, ensure the `name` matches the `eventSourceName`.
* `eventSourceName` (Optional) The partner event source that the new event bus will be matched with. Must match `name`.
* `tags` - (Optional)  A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the event bus.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EventBridge event buses using the `name` (which can also be a partner event source name). For example:

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

Using `terraform import`, import EventBridge event buses using the `name` (which can also be a partner event source name). For example:

```console
% terraform import aws_cloudwatch_event_bus.messenger chat-messages
```

<!-- cache-key: cdktf-0.20.0 input-740b7ef9a57fce58b33a2b98397aad29a82f722b6025aed22c549b727c9580eb -->