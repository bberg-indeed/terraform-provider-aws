---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "AWS: aws_iam_group_policy_attachment"
description: |-
  Attaches a Managed IAM Policy to an IAM group
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_iam_group_policy_attachment

Attaches a Managed IAM Policy to an IAM group

~> **NOTE:** The usage of this resource conflicts with the `aws_iam_policy_attachment` resource and will permanently show a difference if both are defined.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { IamGroup } from "./.gen/providers/aws/iam-group";
import { IamGroupPolicyAttachment } from "./.gen/providers/aws/iam-group-policy-attachment";
import { IamPolicy } from "./.gen/providers/aws/iam-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const group = new IamGroup(this, "group", {
      name: "test-group",
    });
    const policy = new IamPolicy(this, "policy", {
      description: "A test policy",
      name: "test-policy",
      policy: "{ ... policy JSON ... }",
    });
    new IamGroupPolicyAttachment(this, "test-attach", {
      group: group.name,
      policyArn: policy.arn,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `group`  (Required) - The group the policy should be applied to
* `policyArn`  (Required) - The ARN of the policy you want to apply

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import IAM group policy attachments using the group name and policy arn separated by `/`. For example:

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

Using `terraform import`, import IAM group policy attachments using the group name and policy arn separated by `/`. For example:

```console
% terraform import aws_iam_group_policy_attachment.test-attach test-group/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
```

<!-- cache-key: cdktf-0.20.0 input-3386ee59f151e68f1f49aa0bbe51aa6530bcb898108b669fb7540481f63c523b -->