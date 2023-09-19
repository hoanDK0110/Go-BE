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

        //Docker registry
        stage('Build and Push code by Docker') {
            steps {
                withDockerRegistry(credentialsId: 'docker-hub', url: 'https://index.docker.io/v1/') {
                    sh 'docker build -t hoandk0110/golang-web:v1.0 .'
                    sh 'docker push hoandk0110/golang-web:v1.0'
                }
            }
        }
    }
}
