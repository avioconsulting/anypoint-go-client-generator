npm run generate \
&& rm -rf ../apm \
&& cp -r dist/apm ../apm \
&& grep -v "DRAFT DeploymentStatusType = \"DRAFT\"" ../apm/model_deployment_status_type.go > temp \
&& mv temp ../apm/model_deployment_status_type.go \
&& echo "Finished creating new go-client"