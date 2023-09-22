pipeline {
    environment {
        registry = "hoandk0110/golangweb"  // Your Docker Hub account and repository
        registryCredential = 'dockerhub'   // Credential ID for Docker Hub
        dockerImage = ''
    }
    agent any
    stages {
        stage('Cloning our Git') {
            steps {
                // Clone your GitHub repository from the 'main' branch
                git branch: 'main', url: 'https://github.com/hoanDK0110/Go-BE.git'
            }
        }

        stage('SonarQube Analysis') {
            steps {
                script {
                    def scannerHome = tool 'SonarScanner'
                    withSonarQubeEnv('Sonar-Server') {
                        sh "${scannerHome}/bin/sonar-scanner -Dsonar.projectKey=golang-web"
                    }
                }
            }
        }

        stage("Quality Gate") {
            steps {
                script {
                    waitForQualityGate abortPipeline: false, credentialsId: 'sonar-token'
                }
            }
        }
        /*
        stage('Building and Pushing Docker image') {
            steps {
                script {
                    // Define the Dockerfile location (replace with your actual Dockerfile path)
                    def dockerfile = './Dockerfile'

                    // Build the Docker image and tag it with the version (BUILD_NUMBER)
                    dockerImage = docker.build("${registry}:${BUILD_NUMBER}", "-f ${dockerfile} .")

                    // Authenticate and push the Docker image to the registry
                    docker.withRegistry('', registryCredential) {
                        dockerImage.push()
                    }
                }
            }
        }

        stage('Cleaning up') {
            steps {
                // Clean up by removing the locally built Docker image
                sh "docker rmi ${registry}:${BUILD_NUMBER}"
            }
        }
        */
    }
}
