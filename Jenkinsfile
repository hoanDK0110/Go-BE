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
                    
                    // Run SonarQube analysis
                    sh """${scannerHome}/bin/sonar-scanner \
                        -D sonar.projectVersion=1.0-SNAPSHOT \
                        -D sonar.login=admin \
                        -D sonar.password=1 \
                        -D sonar.projectBaseDir=/var/lib/jenkins/workspace/test1/ \
                        -D sonar.projectKey=golang-web \
                        -D sonar.sourceEncoding=UTF-8 \
                        -D sonar.language=golang \
                        -D sonar.sources=project/src/main \
                        -D sonar.tests=project/src/test \
                        -D sonar.host.url=http://192.168.1.25:9000/"""
                }
            }
        }

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
    }
}
