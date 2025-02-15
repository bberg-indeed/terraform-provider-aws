---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_document"
description: |-
  Provides an SSM Document resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssm_document

Provides an SSM Document resource

~> **NOTE on updating SSM documents:** Only documents with a schema version of 2.0
or greater can update their content once created, see [SSM Schema Features][1]. To update a document with an older schema version you must recreate the resource. Not all document types support a schema version of 2.0 or greater. Refer to [SSM document schema features and examples][2] for information about which schema versions are supported for the respective `documentType`.

## Example Usage

### Create an ssm document in JSON format

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmDocument } from "./.gen/providers/aws/ssm-document";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmDocument(this, "foo", {
      content:
        '  {\n    "schemaVersion": "1.2",\n    "description": "Check ip configuration of a Linux instance.",\n    "parameters": {\n\n    },\n    "runtimeConfig": {\n      "aws:runShellScript": {\n        "properties": [\n          {\n            "id": "0.aws:runShellScript",\n            "runCommand": ["ifconfig"]\n          }\n        ]\n      }\n    }\n  }\n\n',
      documentType: "Command",
      name: "test_document",
    });
  }
}

```

### Create an ssm document in YAML format

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmDocument } from "./.gen/providers/aws/ssm-document";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmDocument(this, "foo", {
      content:
        "schemaVersion: '1.2'\ndescription: Check ip configuration of a Linux instance.\nparameters: {}\nruntimeConfig:\n  'aws:runShellScript':\n    properties:\n      - id: '0.aws:runShellScript'\n        runCommand:\n          - ifconfig\n\n",
      documentFormat: "YAML",
      documentType: "Command",
      name: "test_document",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name of the document.
* `attachmentsSource` - (Optional) One or more configuration blocks describing attachments sources to a version of a document. Defined below.
* `content` - (Required) The JSON or YAML content of the document.
* `documentFormat` - (Optional, defaults to JSON) The format of the document. Valid document types include: `JSON` and `YAML`
* `documentType` - (Required) The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
* `permissions` - (Optional) Additional Permissions to attach to the document. See [Permissions](#permissions) below for details.
* `targetType` - (Optional) The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
* `tags` - (Optional) A map of tags to assign to the object. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `versionName` - (Optional) A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.

## attachments_source

The `attachmentsSource` block supports the following:

* `key` - (Required) The key describing the location of an attachment to a document. Valid key types include: `SourceUrl` and `S3FileUrl`
* `values` - (Required) The value describing the location of an attachment to a document
* `name` - (Optional) The name of the document attachment file

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `createdDate` - The date the document was created.
* `description` - The description of the document.
* `schemaVersion` - The schema version of the document.
* `defaultVersion` - The default version of the document.
* `documentVersion` - The document version.
* `hash` - The sha1 or sha256 of the document content
* `hashType` - "Sha1" "Sha256". The hashing algorithm used when hashing the content.
* `latestVersion` - The latest version of the document.
* `owner` - The AWS user account of the person who created the document.
* `status` - "Creating", "Active" or "Deleting". The current status of the document.
* `parameter` - The parameters that are available to this document.
* `platformTypes` - A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

[1]: http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-ssm-docs.html#document-schemas-features
[2]: https://docs.aws.amazon.com/systems-manager/latest/userguide/document-schemas-features.html

## Permissions

The permissions attribute specifies how you want to share the document. If you share a document privately,
you must specify the AWS user account IDs for those people who can use the document. If you share a document
publicly, you must specify All as the account ID.

The permissions mapping supports the following:

* `type` - The permission type for the document. The permission type can be `Share`.
* `accountIds` - The AWS user accounts that should have access to the document. The account IDs can either be a group of account IDs or `All`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SSM Documents using the name. For example:

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

Using `terraform import`, import SSM Documents using the name. For example:

```console
% terraform import aws_ssm_document.example example
```

The `attachmentsSource` argument does not have an SSM API method for reading the attachment information detail after creation. If the argument is set in the Terraform configuration on an imported resource, Terraform will always show a difference. To workaround this behavior, either omit the argument from the Terraform configuration or use [`ignore_changes`](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#ignore_changes) to hide the difference. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmDocument } from "./.gen/providers/aws/ssm-document";
interface MyConfig {
  content: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new SsmDocument(this, "test", {
      attachmentsSource: [
        {
          key: "SourceUrl",
          values: ["s3://${" + objectBucket.bucket + "}/test.zip"],
        },
      ],
      documentType: "Package",
      lifecycle: {
        ignoreChanges: [attachmentsSource],
      },
      name: "test_document",
      content: config.content,
    });
  }
}

```

<!-- cache-key: cdktf-0.20.0 input-0f47ccad709627d6f83189f2b6d41b8032365327e11d5385345cf2f96e8c2bad -->