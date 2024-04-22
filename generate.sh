# Runs project generation and removes incorrect file that references DRAFT in model_deployment_status_type.go. Used to install locally and in this repository
npm run generate
# Removes apm folder from parent directory
rm -rf ../apm 
# Copies newly generated apm folder to parent directory for local testing
cp -r apm ../apm
# Creates temp file from model_deployment_status_type.go removing one line. This line is removed because it creates a conflict with another global variable named DRAFT
grep -v "DRAFT DeploymentStatusType = \"DRAFT\"" ../apm/model_deployment_status_type.go > temp
# Remove generated files using DRAFT line
rm ../apm/model_deployment_status_type.go
rm apm/model_deployment_status_type.go
# Copy temp file to model_deployment_status_type.go w/o the DRAFT line then remove 
cp temp ../apm/model_deployment_status_type.go
cp temp apm/model_deployment_status_type.go
rm temp
mv apm/go.mod go.mod
go mod tidy
echo "Finished creating new go-client"

