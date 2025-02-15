---
subcategory: "FSx"
layout: "aws"
page_title: "AWS: aws_fsx_ontap_volume"
description: |-
  Manages a FSx ONTAP Volume.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_fsx_ontap_volume

Manages a FSx ONTAP Volume.
See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-volumes.html) for more information.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { FsxOntapVolume } from "./.gen/providers/aws/fsx-ontap-volume";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new FsxOntapVolume(this, "test", {
      junctionPath: "/test",
      name: "test",
      sizeInMegabytes: 1024,
      storageEfficiencyEnabled: true,
      storageVirtualMachineId: Token.asString(
        awsFsxOntapStorageVirtualMachineTest.id
      ),
    });
  }
}

```

### Using Tiering Policy

Additional information on tiering policy with ONTAP Volumes can be found in the [FSx ONTAP Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-volumes.html).

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { FsxOntapVolume } from "./.gen/providers/aws/fsx-ontap-volume";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new FsxOntapVolume(this, "test", {
      junctionPath: "/test",
      name: "test",
      sizeInMegabytes: 1024,
      storageEfficiencyEnabled: true,
      storageVirtualMachineId: Token.asString(
        awsFsxOntapStorageVirtualMachineTest.id
      ),
      tieringPolicy: {
        coolingPeriod: 31,
        name: "AUTO",
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name of the Volume. You can use a maximum of 203 alphanumeric characters, plus the underscore (_) special character.
* `bypassSnaplockEnterpriseRetention` - (Optional) Setting this to `true` allows a SnapLock administrator to delete an FSx for ONTAP SnapLock Enterprise volume with unexpired write once, read many (WORM) files. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
* `copyTagsToBackups` - (Optional) A boolean flag indicating whether tags for the volume should be copied to backups. This value defaults to `false`.
* `junctionPath` - (Optional) Specifies the location in the storage virtual machine's namespace where the volume is mounted. The junction_path must have a leading forward slash, such as `/vol3`
* `ontapVolumeType` - (Optional) Specifies the type of volume, valid values are `RW`, `DP`. Default value is `RW`. These can be set by the ONTAP CLI or API. This setting is used as part of migration and replication [Migrating to Amazon FSx for NetApp ONTAP](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/migrating-fsx-ontap.html)
* `securityStyle` - (Optional) Specifies the volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`.
* `sizeInMegabytes` - (Required) Specifies the size of the volume, in megabytes (MB), that you are creating.
* `skipFinalBackup` - (Optional) When enabled, will skip the default final backup taken when the volume is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
* `snaplockConfiguration` - (Optional) The SnapLock configuration for an FSx for ONTAP volume. See [SnapLock Configuration](#snaplock-configuration) below.
* `snapshotPolicy` - (Optional) Specifies the snapshot policy for the volume. See [snapshot policies](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snapshots-ontap.html#snapshot-policies) in the Amazon FSx ONTAP User Guide
* `storageEfficiencyEnabled` - (Optional) Set to true to enable deduplication, compression, and compaction storage efficiency features on the volume.
* `storageVirtualMachineId` - (Required) Specifies the storage virtual machine in which to create the volume.
* `tags` - (Optional) A map of tags to assign to the volume. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `tieringPolicy` - (Optional) The data tiering policy for an FSx for ONTAP volume. See [Tiering Policy](#tiering-policy) below.

### SnapLock Configuration

* `auditLogVolume` - (Optional) Enables or disables the audit log volume for an FSx for ONTAP SnapLock volume. The default value is `false`.
* `autocommitPeriod` - (Optional) The configuration object for setting the autocommit period of files in an FSx for ONTAP SnapLock volume. See [Autocommit Period](#autocommit-period) below.
* `privilegedDelete` - (Optional) Enables, disables, or permanently disables privileged delete on an FSx for ONTAP SnapLock Enterprise volume. Valid values: `DISABLED`, `ENABLED`, `PERMANENTLY_DISABLED`. The default value is `DISABLED`.
* `retentionPeriod` - (Optional) The retention period of an FSx for ONTAP SnapLock volume. See [SnapLock Retention Period](#snaplock-retention-period) below.
* `snaplockType` - (Required) Specifies the retention mode of an FSx for ONTAP SnapLock volume. After it is set, it can't be changed. Valid values: `COMPLIANCE`, `ENTERPRISE`.
* `volumeAppendModeEnabled` - (Optional) Enables or disables volume-append mode on an FSx for ONTAP SnapLock volume. The default value is `false`.

### Autocommit Period

* `type` - (Required) The type of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume. Setting this value to `NONE` disables autocommit. Valid values: `MINUTES`, `HOURS`, `DAYS`, `MONTHS`, `YEARS`, `NONE`.
* `value` - (Optional) The amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.

### SnapLock Retention Period

* `defaultRetention` - (Required) The retention period assigned to a write once, read many (WORM) file by default if an explicit retention period is not set for an FSx for ONTAP SnapLock volume. The default retention period must be greater than or equal to the minimum retention period and less than or equal to the maximum retention period. See [Retention Period](#retention-period) below.
* `maximumRetention` - (Required) The longest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume. See [Retention Period](#retention-period) below.
* `minimumRetention` - (Required) The shortest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume. See [Retention Period](#retention-period) below.

### Retention Period

* `type` - (Required) The type of time for the retention period of an FSx for ONTAP SnapLock volume. Set it to one of the valid types. If you set it to `INFINITE`, the files are retained forever. If you set it to `UNSPECIFIED`, the files are retained until you set an explicit retention period. Valid values: `SECONDS`, `MINUTES`, `HOURS`, `DAYS`, `MONTHS`, `YEARS`, `INFINITE`, `UNSPECIFIED`.
* `value` - (Optional) The amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.

### Tiering Policy

* `name` - (Required) Specifies the tiering policy for the ONTAP volume for moving data to the capacity pool storage. Valid values are `SNAPSHOT_ONLY`, `AUTO`, `ALL`, `NONE`. Default value is `SNAPSHOT_ONLY`.
* `coolingPeriod` - (Optional) Specifies the number of days that user data in a volume must remain inactive before it is considered "cold" and moved to the capacity pool. Used with `AUTO` and `SNAPSHOT_ONLY` tiering policies only. Valid values are whole numbers between 2 and 183. Default values are 31 days for `AUTO` and 2 days for `SNAPSHOT_ONLY`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name of the volune.
* `id` - Identifier of the volume, e.g., `fsvol-12345678`
* `fileSystemId` - Describes the file system for the volume, e.g. `fs-12345679`
* `flexcacheEndpointType` - Specifies the FlexCache endpoint type of the volume, Valid values are `NONE`, `ORIGIN`, `CACHE`. Default value is `NONE`. These can be set by the ONTAP CLI or API and are use with FlexCache feature.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `uuid` - The Volume's UUID (universally unique identifier).
* `volumeType` - The type of volume, currently the only valid value is `ONTAP`.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `delete` - (Default `30m`)
* `update` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import FSx ONTAP volume using the `id`. For example:

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

Using `terraform import`, import FSx ONTAP volume using the `id`. For example:

```console
% terraform import aws_fsx_ontap_volume.example fsvol-12345678abcdef123
```

<!-- cache-key: cdktf-0.20.0 input-4150b6016178afdca4f0be9139d92abeffcd41f64e3fc04737a2e1147ba2af7e -->