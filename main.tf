resource "aws_s3_bucket" "backend_bucket" {
  bucket = "${var.bucket_name}"
}

resource "aws_dynamodb_table" "locking_table" {
  name           = "${var.dynamodb_table_name}"
  hash_key       = "LockID"
  read_capacity  = 20
  write_capacity = 20

  attribute {
    name = "LockID"
    type = "S"
  }
}
