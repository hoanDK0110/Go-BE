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

        
        stage('Build docker image') {
            steps {  
                sh 'docker build -t ylmt/flaskapp:$BUILD_NUMBER .'
            }
        }

        /* Docker registry
        stage('Build and Push code by Docker') {
            steps {
                script {
                    def dockerImage = 'hoandk0110/golang-web:v1.0'
                    def dockerToken = 'dckr_pat_wlxrClD2X0HDG88jRmVgtunf0nY'
                    
                    // Log in to Docker Hub using the access token
                    sh "docker login -u hoandk0110 -p $dockerToken"
                    
                    // Build and push the Docker image
                    sh "docker build -t $dockerImage ."
                    sh "docker push $dockerImage"
                }
            }
        }
        */
    }
}
