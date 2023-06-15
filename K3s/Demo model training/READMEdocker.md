In order to utlize the docker file and pushing image to use , following commands can be followed:

# Pushing image to the global repository and pulling it from there to be used in cluster
1. docker build -t ml-train-example:1.0 .
2. docker tag ml-train-example:1.0 sarvagya23/ml-train-example:1.0
3. docker push sarvagya23/ml-train-example:1.0

# Use it for local testing
1. docker build -t ml-train-example:1.0 .
2. docker run --name new-app -d ml-train-example:1.0

#Once this is done, deployment can be applied with the following command
3. sudo k3s kubectl apply -f deployment tensorflow-deployment2.yaml

#Then we can get the pod and view its logs
4. sudo k3s kubectl get pods 
5. sudo k3s kubectl logs  <pod-name>
