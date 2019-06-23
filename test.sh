#!/bin/bash
set -ex
docker run -it \
-v $(pwd):/go/src/github.com/mikhailadvani/terraform-s3-backend \
-v ~/.aws/credentials:/root/.aws/credentials \
-e AWS_PROFILE=$AWS_PROFILE \
-w /go/src/github.com/mikhailadvani/terraform-s3-backend \
--entrypoint /bin/bash \
terraform-workshop:latest \
$@
