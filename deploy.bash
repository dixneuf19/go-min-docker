PROJECT_NAME="go-min-docker"
GOOGLE_PROJECT="********"
REGISTRY_URL="eu.gcr.io/********"
COMMIT_HASH=$(git rev-parse --short HEAD)
DOCKER_IMAGE="$REGISTRY_URL/$PROJECT_NAME-$COMMIT_HASH"
KUBERNETES_CONFIG_SECRETNAME="$PROJECT_NAME-secret"

cd "$GOPATH/src/github.com/dixneuf19/$PROJECT_NAME"
dep ensure -vendor-only
go test -v ./...

docker build -t "$DOCKER_IMAGE" .
docker push "$DOCKER_IMAGE"

DOCKER_IMAGE=$DOCKER_IMAGE COMMIT_HASH=$COMMIT_HASH envsubst  < kubernetes.template.yml > kubernetes.yml

kubectl create secret generic "$KUBERNETES_CONFIG_SECRETNAME" --from-file=secret/secret.yml --dry-run -o yaml | kubectl apply -f -
kubectl apply -f kubernetes.yml


