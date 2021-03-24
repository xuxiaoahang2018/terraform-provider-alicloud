---
subcategory: "Log Service (SLS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_log_etl"
sidebar_current: "docs-alicloud-resource-log-etl"
description: |-
  Provides a Alicloud log etl resource.
---

# alicloud\_log\_etl

The data processing function of the log service is a hosted, highly available, and scalable data processing service, 
which is widely applicable to scenarios such as data regularization, enrichment, distribution, aggregation, and index reconstruction.
[Refer to details](https://www.alibabacloud.com/help/zh/doc-detail/125384.htm).

## Example Usage

Basic Usage

```
resource "alicloud_log_project" "example" {
  name        = "tf-log"
  description = "created by terraform"

}

resource "alicloud_log_store" "example" {
  project               = alicloud_log_project.example.name
  name                  = "tf-test-logstore"
  retention_period      = 3650
  shard_count           = 3
  auto_split            = true
  max_split_shard_count = 60
  append_meta           = true
}

resource "alicloud_log_store" "example2" {
  project               = alicloud_log_project.example.name
  name                  = "tf-test-logstore-2"
  retention_period      = 3650
  shard_count           = 3
  auto_split            = true
  max_split_shard_count = 60
  append_meta           = true
}


resource "alicloud_log_store" "example3" {
  project               = alicloud_log_project.example.name
  name                  = "tf-test-logstore-3"
  retention_period      = 3650
  shard_count           = 3
  auto_split            = true
  max_split_shard_count = 60
  append_meta           = true
}

resource "alicloud_log_etl" "example" {
   etl_name = "etl_name"
   project = alicloud_log_project.example.name
   display_name = "display_name"
   description = "etl_description"
   access_key_id = "access_key_id"
   access_key_secret = "access_key_secret"
   script = "e_set('new','key')"
   logstore = alicloud_log_store.example.name
   etl_sinks {
        name = "target_name",
        access_key_id = "example2_access_key_id",
        access_key_secret = "example2_access_key_secret",
        endpoint = "cn-hangzhou.log.aliyuncs.com",
        project = alicloud_log_project.example.name
        logstore = alicloud_log_store.example2.name
   }
   etl_sinks {
           name = "target_name2",
           access_key_id = "example3_access_key_id",
           access_key_secret = "example3_access_key_secret",
           endpoint = "cn-hangzhou.log.aliyuncs.com",
           project = alicloud_log_project.example.name
           logstore = alicloud_log_store.example3.name
   }

}

```

Stop the task in progress

```
resource "alicloud_log_etl" "example" {
   status = STOPPED
   etl_name = "etl_name"
   project = alicloud_log_project.example.name
   display_name = "display_name"
   description = "etl_description"
   access_key_id = "access_key_id"
   access_key_secret = "access_key_secret"
   script = "e_set('new','key')"
   logstore = alicloud_log_store.example.name
   etl_sinks {
        name = "target_name",
        access_key_id = "example2_access_key_id",
        access_key_secret = "example2_access_key_secret",
        endpoint = "cn-hangzhou.log.aliyuncs.com",
        project = alicloud_log_project.example.name
        logstore = alicloud_log_store.example2.name
   }
   etl_sinks {
           name = "target_name2",
           access_key_id = "example3_access_key_id",
           access_key_secret = "example3_access_key_secret",
           endpoint = "cn-hangzhou.log.aliyuncs.com",
           project = alicloud_log_project.example.name
           logstore = alicloud_log_store.example3.name
   }

}

```

ReStart the stopped task

```

resource "alicloud_log_etl" "example" {
   status = RUNNING
   etl_name = "etl_name"
   project = alicloud_log_project.example.name
   display_name = "display_name"
   description = "etl_description"
   access_key_id = "access_key_id"
   access_key_secret = "access_key_secret"
   script = "e_set('new','key')"
   logstore = alicloud_log_store.example.name
   etl_sinks {
        name = "target_name",
        access_key_id = "example2_access_key_id",
        access_key_secret = "example2_access_key_secret",
        endpoint = "cn-hangzhou.log.aliyuncs.com",
        project = alicloud_log_project.example.name
        logstore = alicloud_log_store.example2.name
   }
   etl_sinks {
           name = "target_name2",
           access_key_id = "example3_access_key_id",
           access_key_secret = "example3_access_key_secret",
           endpoint = "cn-hangzhou.log.aliyuncs.com",
           project = alicloud_log_project.example.name
           logstore = alicloud_log_store.example3.name
   }

}

```



## Argument Reference

The following arguments are supported:

* `etl_name` - (Required, ForceNew) The name of the log etl job.
* `description` - (Optional) Description of the log etl job.
* `project` - (Required, ForceNew) The name of the project where the etl job is located.
* `display_name` - (Required) Log service etl job alias.
* `schedule` - (Optional) Job scheduling type, the default value is Resident.
* `etl_type` - (Optional) Log service etl type, the default value is ETL.
* `status` - (Optional, Computed) Log project tags. the default value is RUNNING, Only 4 values are supported: STARTING，RUNNING，STOPPING，STOPPED.
* `create_time` - (Optional, Computed) The etl job create time.
* `last_modified_time` - (Optional, Computed) ETL job last modified time.
* `access_key_id` - (Required) Source logstore access key id.
* `access_key_secret` - (Optional) Source logstore access key secret
* `from_time` - (Optional,Computed) The start time of the processing job, the default starts from the current time.
* `to_time` - (Optional) Deadline of processing job, the default value is None.
* `script` - (Required) Processing operation grammar.
* `version` - (Optional) Log etl job version. the default value is 2.
* `logstore` - (Required) The source logstore of the processing job.
* `parameters` - (Optional) Advanced parameter configuration of processing operations.
* `role_arn` - (Optional) Sts role info.
* `etl_sinks` - (Required) Target logstore configuration for delivery after data processing.
    * `access_key_id` - (Required) Delivery target logstore access key id.
    * `access_key_secret`- (Required) Delivery target logstore access key secret.
    * `endpoint` - (Required) Delivery target logstore region.
    * `name` - (Required) Delivery target name.
    * `project` - (Required) The project where the target logstore is delivered.
    * `logstore` - (Required) Delivery target logstore.
    * `role_arn` - (Required) Sts role info.
    * `type` - (Optional)  ETL sinks type, the default value is AliyunLOG.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the log etl. It sames as its name.

## Import

Log etl can be imported using the id, e.g.

```
$ terraform import alicloud_log_etl.example tf-log-project:tf-log-etl-name
```
