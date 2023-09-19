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
                    withCredentials([usernamePassword(credentialsId: 'docker-hub', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh "docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD"
                    }

                    // Build và push image
                    sh "docker build -t $dockerImage ."
                    sh "docker push $dockerImage"
                }
            }
        }
    }
}
