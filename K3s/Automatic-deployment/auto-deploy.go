package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func main() {
	var dockerImage, namespace, fileLocation string

	// Prompt for Docker image name
	fmt.Print("Enter Docker image name: ")
	fmt.Scanln(&dockerImage)

	// Prompt for namespace
	fmt.Print("Enter namespace: ")
	fmt.Scanln(&namespace)

	fmt.Print("Enter file location: ")
	fmt.Scanln(&fileLocation)

	// Generate a base name for metadata and app labels
	baseName := "test"

	// Append a timestamp to the base name for uniqueness
	timestamp := time.Now().Format("20060102-150405")

	// Construct unique metadata.name and spec.template.metadata.labels.app names
	deploymentName := fmt.Sprintf("%s-deployment-%s", baseName, timestamp)
	appLabel := fmt.Sprintf("%s-app-%s", baseName, timestamp)

	// Generate Deployment YAML with unique names
	deploymentYAML := generateDeploymentYAML(deploymentName, dockerImage, namespace, appLabel, fileLocation)

	// Write Deployment YAML to a generated filename
	deploymentFileName := fmt.Sprintf("deployment-%s.yaml", timestamp)
	err := ioutil.WriteFile(deploymentFileName, []byte(deploymentYAML), 0644)
	if err != nil {
		fmt.Printf("Error writing to %s: %v\n", deploymentFileName, err)
		os.Exit(1)
	}

	fmt.Printf("Deployment YAML written to %s\n", deploymentFileName)

	// Use kubectl to apply the Deployment to the cluster
	applyCmd := exec.Command("sudo", "k3s", "kubectl", "apply", "-f", deploymentFileName)
	applyCmd.Stdout = os.Stdout
	applyCmd.Stderr = os.Stderr

	fmt.Println("Applying Deployment to the Kubernetes cluster...")
	err = applyCmd.Run()
	if err != nil {
		fmt.Printf("Error applying Deployment: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Deployment applied successfully.")
}

func generateDeploymentYAML(deploymentName, image, namespace, appLabel, fileLocation string) string {
	return fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
  namespace: %s
spec:
  replicas: 1
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
      - name: %s-container
        image: %s
        command: ["python", %s]
`, deploymentName, namespace, appLabel, appLabel, deploymentName, image, fileLocation)
}
