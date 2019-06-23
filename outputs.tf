output "dynamodb_table_arn" {
  value = "${aws_dynamodb_table.locking_table.arn}"
}

output "s3_bucket_arn" {
  value = "${aws_s3_bucket.backend_bucket.arn}"
}
