---
subcategory: "Elastic Beanstalk"
layout: "aws"
page_title: "AWS: aws_elastic_beanstalk_application"
description: |-
  Retrieve information about an Elastic Beanstalk Application
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_elastic_beanstalk_application

Retrieve information about an Elastic Beanstalk Application.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformOutput, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_elastic_beanstalk_application import DataAwsElasticBeanstalkApplication
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsElasticBeanstalkApplication(self, "example",
            name="example"
        )
        TerraformOutput(self, "arn",
            value=example.arn
        )
        TerraformOutput(self, "description",
            value=example.description
        )
```

## Argument Reference

* `name` - (Required) Name of the application

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - Name of the application
* `arn` - ARN of the application.
* `description` - Short description of the application

Application version lifecycle (`appversion_lifecycle`) supports the nested attribute containing.

* `service_role` - ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
* `max_count` - Maximum number of application versions to retain.
* `max_age_in_days` - Number of days to retain an application version.
* `delete_source_from_s3` - Specifies whether delete a version's source bundle from S3 when the application version is deleted.

<!-- cache-key: cdktf-0.20.0 input-91e8e35e02809b1291798f560a469b0ee44ac2ca2072d3eb5b99110f092216bf -->