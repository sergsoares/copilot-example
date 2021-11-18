
# Variables used to create app2
export PROJECT_DOCKER_PATH=./app2/Dockerfile
export PROJECT_NAME='applications'
export SERVICE_NAME='app2'
export AWS_REGION=us-east-1

copilot app init $PROJECT_NAME --resource-tags "project=${PROJECT_NAME},source=copilot"

# Command for init app
copilot init --app $PROJECT_NAME \
  --name $SERVICE_NAME \
  --type 'Load Balanced Web Service' \
  --dockerfile $PROJECT_DOCKER_PATH \
  --port 80

---

# Variables used to create app2
export PROJECT_DOCKER_PATH=./app2/Dockerfile
export PROJECT_NAME='applications'
export SERVICE_NAME='app2'
export AWS_REGION=us-east-1

# Command for init app
copilot init --app $PROJECT_NAME \
  --name $SERVICE_NAME \
  --type 'Load Balanced Web Service' \
  --dockerfile $PROJECT_DOCKER_PATH \
  --port 80

## Initialize ENV (create )

export ENV_NAME=prod
export APP_NAME='app1'
export PUBLIC_SUBNETS='subnet-xxxxxxxxxxxxx,subnet-yyyyyyyyyyyy'
export PRIVATE_SUBNETS='subnet-xxxxxxxxxxxxx,subnet-yyyyyyyyyyyy'
export VPC_ID='vpc-zzzzzzzzzzz'
export PROFILE=default

copilot env init --name $ENV_NAME \
--profile $PROFILE \
--prod \
--import-vpc-id $VPC_ID \
--import-public-subnets $PUBLIC_SUBNETS \
--import-private-subnets $PRIVATE_SUBNETS

##
copilot app ls

##
copilot env ls

## Deploy app1 with tags
export ENV_NAME=prod
export SERVICE_NAME='app1'
export PROJECT_NAME='applications'

copilot svc deploy --env $ENV_NAME --name $SERVICE_NAME --resource-tags "project=${PROJECT_NAME},source=copilot"

## Deploy app2 with tags
export ENV_NAME=prod
export SERVICE_NAME='app2'
export PROJECT_NAME='applications'

copilot svc deploy --env $ENV_NAME --name $SERVICE_NAME --resource-tags "project=${PROJECT_NAME},source=copilot"

## Pipeline
export ENV_NAME=prod
export SERVICE_NAME='app1'
export PROJECT_NAME='applications'
export GITHUB_REPO_URL='https://github.com/sergsoares/copilot-example.git'
export GITHUB_BRANCH='https://github.com/sergsoares/copilot-example.git'

copilot pipeline init --app $PROJECT_NAME --environments $ENV_NAME --url $GITHUB_REPO_URL --git-branch $GITHUB_BRANCH

copilot pipeline update

copilot pipeline status

## Delete pipeline

copilot pipeline delete

## Delete app1
export ENV_NAME=prod
export SERVICE_NAME='app1'

copilot svc delete --env $ENV_NAME --name $SERVICE_NAME 


## Delete app2
export ENV_NAME=prod
export SERVICE_NAME='app2'

copilot svc delete --env $ENV_NAME --name $SERVICE_NAME

# Delete entire environment

export ENV_NAME=prod

copilot env delete --name $ENV_NAME

# Delete entire app

export ENV_NAME=prod

copilot env delete --name $ENV_NAME

## TEST Script

URL=<PUT_APP_URL>
while true; sleep 0.2; do curl $URL -i; echo''; done

## Alternative 

aws ecs update-service --cluster <cluster name> --service <service name> --force-new-deployment

## References

- https://stackoverflow.com/questions/34840137/how-do-i-deploy-updated-docker-images-to-amazon-ecs-tasks
- https://aws.github.io/copilot-cli/docs/commands/pipeline-init/
- https://github.com/paddycarey/go-echo/blob/master/main.go