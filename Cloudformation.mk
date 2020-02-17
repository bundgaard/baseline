## VALIDATE TEMPLATE
# $1 TEMPLATE FILE
cf-validate=aws cloudformation validate-template \
	--template-body file://$1

## PACKAGE STACK
# $1 TEMPLATE_FILE
# $2 S3_BUCKET
cf-package=aws cloudformation package \
	--template-file $1 \
	--s3-bucket $2 \
	--output json \
	--output-template-file $(subst .yml,,$(2))_packaged.yml

## DEPLOY STACK
# $1 STACK_NAME
# $2 TEMPLATE_FILE
# $3 PARAMETERS Key=Value space separated list
cf-deploy=aws cloudformation deploy \
	--capabilities CAPABILITY_IAM \
	--template-file $(subst .yml,,$(2))_packaged.yml \
	--stack-name $1 \
	$(if $3,--parameters-override $(foreach parm,$(3), $(parm)))
	

## DELETE STACK
# $1 STACK_NAME
cf-delete=aws cloudformation delete-stack --stack-name $1


## EMPTY BUCKET
# $1 STACK_NAME
cf-empty-bucket=aws s3 rm s3://$1


ifneq ($(AWS_PROFILE),)
cf-validate += --profile $(AWS_PROFILE)
cf-package +=  --profile $(AWS_PROFILE)
cf-deploy +=   --profile $(AWS_PROFILE)
cf-delete +=   --profile $(AWS_PROFILE)
endif

ifneq ($(AWS_REGION),)
cf-validate += --region $(AWS_REGION)
cf-package +=  --region $(AWS_REGION)
cf-deploy +=   --region $(AWS_REGION)
cf-delete +=   --region $(AWS_REGION)
endif

