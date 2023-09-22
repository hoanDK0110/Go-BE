pipeline {
    environment {
        APP_NAME = "golangweb"
        RELEASE = "1.0.0"
        DOCKER_USER = "hoandk0110"
        IMAGE_NAME = "${DOCKER_USER}" + "/" + "${APP_NAME}"
        IMAGE_TAG = "${RELEASE}-${BUILD_NUMBER}"
        registryCredential = 'dockerhub'   // Credential ID for Docker Hub
    }
    agent any
    stages {
        stage('Checkout Code') {
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

        stage('Building Docker image') {
            steps {
                script {
                    // Define the Dockerfile location (replace with your actual Dockerfile path)
                    def dockerfile = './Dockerfile'

                    // Build the Docker image and tag it with the version (BUILD_NUMBER)
                    docker_image = docker.build("${IMAGE_NAME}", "-f ${dockerfile} .")
                }
            }
        }
    }
    
    stage('Pushing Docker image') {
        steps {
            script {
                // Authenticate and push the Docker image to the registry
                docker.withRegistry('', registryCredential) {
                    docker_image.push("${IMAGE_TAG}")
                    docker_image.push('latest')
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
}
