---
subcategory: "Secrets Manager"
layout: "aws"
page_title: "AWS: aws_secretsmanager_secret"
description: |-
  Retrieve metadata information about a Secrets Manager secret
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_secretsmanager_secret

Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the [`aws_secretsmanager_secret_version` data source](/docs/providers/aws/d/secretsmanager_secret_version.html).

## Example Usage

### ARN

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_secretsmanager_secret import DataAwsSecretsmanagerSecret
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSecretsmanagerSecret(self, "by-arn",
            arn="arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456"
        )
```

### Name

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_secretsmanager_secret import DataAwsSecretsmanagerSecret
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSecretsmanagerSecret(self, "by-name",
            name="example"
        )
```

## Argument Reference

* `arn` - (Optional) ARN of the secret to retrieve.
* `name` - (Optional) Name of the secret to retrieve.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the secret.
* `description` - Description of the secret.
* `kms_key_id` - Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.
* `id` - ARN of the secret.
* `tags` - Tags of the secret.
* `policy` - Resource-based policy document that's attached to the secret.

<!-- cache-key: cdktf-0.20.0 input-0fca6cfb25caf0f32cdf4fba62b9433e266444e38aa59620febc991b1723f0c4 -->