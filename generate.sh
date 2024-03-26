`# Runs project generation`
npm run generate \
&& rm -rf ../apm `# Removes apm folder from parent directory` \
&& cp -r dist/apm ../apm `# Copies newly generated apm folder to parent directory` \
&& grep -v "DRAFT DeploymentStatusType = \"DRAFT\"" ../apm/model_deployment_status_type.go > temp `# Creates temp file from model_deployment_status_type.go removing one line` `# This line is removed because it creates a conflict with another global variable named DRAFT` \
&& mv temp ../apm/model_deployment_status_type.go `# Replaces model_deployment_status_type.go with temp file` \
&& echo "Finished creating new go-client"

