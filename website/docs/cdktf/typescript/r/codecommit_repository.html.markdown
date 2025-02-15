---
subcategory: "CodeCommit"
layout: "aws"
page_title: "AWS: aws_codecommit_repository"
description: |-
  Provides a CodeCommit Repository Resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_codecommit_repository

Provides a CodeCommit Repository Resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CodecommitRepository } from "./.gen/providers/aws/codecommit-repository";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CodecommitRepository(this, "test", {
      description: "This is the Sample App Repository",
      repositoryName: "MyTestRepository",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `repositoryName` - (Required) The name for the repository. This needs to be less than 100 characters.
* `description` - (Optional) The description of the repository. This needs to be less than 1000 characters
* `defaultBranch` - (Optional) The default branch of the repository. The branch specified here needs to exist.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `repositoryId` - The ID of the repository
* `arn` - The ARN of the repository
* `cloneUrlHttp` - The URL to use for cloning the repository over HTTPS.
* `cloneUrlSsh` - The URL to use for cloning the repository over SSH.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CodeCommit repository using repository name. For example:

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

Using `terraform import`, import CodeCommit repository using repository name. For example:

```console
% terraform import aws_codecommit_repository.imported ExistingRepo
```

<!-- cache-key: cdktf-0.20.0 input-50d294f897d743c5c8bc69d3979c0e0ea1f7abc83eb38554128f2468842056ab -->