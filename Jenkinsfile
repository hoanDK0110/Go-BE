pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
        def dockerImage = 'hoandk0110/golang-web:v1.0'
    }
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

        
        stage('Build docker image') {
            steps {  
                script {
                    docker.build('$dockerImage:$BUILD_NUMBER')
                }
            }
        }

    }
}
