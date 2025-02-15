---
subcategory: "CloudWatch RUM"
layout: "aws"
page_title: "AWS: aws_rum_metrics_destination"
description: |-
  Provides a CloudWatch RUM Metrics Destination resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_rum_metrics_destination

Provides a CloudWatch RUM Metrics Destination resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RumMetricsDestination } from "./.gen/providers/aws/rum-metrics-destination";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new RumMetricsDestination(this, "example", {
      appMonitorName: Token.asString(awsRumAppMonitorExample.name),
      destination: "CloudWatch",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `appMonitorName` - (Required) The name of the CloudWatch RUM app monitor that will send the metrics.
* `destination` - (Required)  Defines the destination to send the metrics to. Valid values are `CloudWatch` and `Evidently`. If you specify `Evidently`, you must also specify the ARN of the CloudWatchEvidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
* `destinationArn` - (Optional) Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
* `iamRoleArn` - (Optional) This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The name of the CloudWatch RUM app monitor that will send the metrics.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cloudwatch RUM Metrics Destination using the `id`. For example:

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

Using `terraform import`, import Cloudwatch RUM Metrics Destination using the `id`. For example:

```console
% terraform import aws_rum_metrics_destination.example example
```

<!-- cache-key: cdktf-0.20.0 input-b0a8bc87eed073c64dab3eea6177045d7d17087af1a9092e3e8b592be02eb910 -->