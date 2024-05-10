# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_servicecatalog_portfolio" "test" {
  name          = var.rName
  description   = "test-b"
  provider_name = "test-c"

  tags = var.tags
}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}

variable "tags" {
  description = "Tags to set on resource. To specify no tags, set to `null`"
  # Not setting a default, so that this must explicitly be set to `null` to specify no tags
  type     = map(string)
  nullable = true
}

