variable "bucket_name" {
  description = "Name of the S3 bucket to be used to store state"
}

variable "dynamodb_table_name" {
  description = "Name of the dynamodb table to be used for locking"
}
