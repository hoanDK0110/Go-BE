pipeline {
    agent any
    stages {
        stage('Cleanup Workspace') {
            steps {
                cleanWs()
            }
        }

        // Clone code SCM
        stage('Clone code') {
            steps {
                script {
                    git branch: 'main', url: 'https://github.com/hoanDK0110/Go-BE.git'
                }
            }
        }

        // Docker registry
        stage('Build and Push code by Docker') {
            steps {
                script {
                    def dockerImage = 'hoandk0110/golang-web:v1.0'

                    // Đăng nhập vào Docker Hub bằng credentials
                    withCredentials([usernamePassword(credentialsId: 'docker-hub-credentials', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh "echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin"
                    }

                    // Build và push image
                    sh "docker build -t $dockerImage ."
                    sh "docker push $dockerImage"
                }
            }
        }
    }
}
