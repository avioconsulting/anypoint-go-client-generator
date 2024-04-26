export "LOCAL_DIR_PATH=../terraform-provider-anypoint-partner-manager/anypoint-go-client-generator"
# Runs project generation and removes incorrect file that references DRAFT in model_deployment_status_type.go. Used to install locally and in this repository
npm run generate
# Removes apm folder from parent directory
rm -rf $LOCAL_DIR_PATH/apm
mkdir -p $LOCAL_DIR_PATH/apm
# Copies newly generated apm folder to parent directory for local testing
cp -r apm/ $LOCAL_DIR_PATH/apm/
# Creates temp file from model_deployment_status_type.go removing one line. This line is removed because it creates a conflict with another global variable named DRAFT
grep -v "DRAFT DeploymentStatusType = \"DRAFT\"" ../apm/model_deployment_status_type.go > temp
# Remove generated files using DRAFT line
rm $LOCAL_DIR_PATH/apm/model_deployment_status_type.go
rm apm/model_deployment_status_type.go
# Copy temp file to model_deployment_status_type.go w/o the DRAFT line then remove 
cp temp $LOCAL_DIR_PATH/apm/model_deployment_status_type.go
cp temp apm/model_deployment_status_type.go
rm temp
cp apm/go.mod $LOCAL_DIR_PATH/go.mod
mv apm/go.mod go.mod
rm apm/go.sum
go mod tidy
echo "Finished creating new go-client"

