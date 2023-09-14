pipeline {
    agent any
    environment {
        DOCKER_IMAGE_NAME = 'golang-web:1.0' 
    }
    stages {
        
        stage('Cleanup Workspace') {
            steps {
                cleanWs()
            }
        }
        
        stage('Clone code') {
            steps {
                // Sao chép mã nguồn từ GitHub
                git branch: 'main', url: 'https://github.com/hoanDK0110/Go-BE.git'
            }
        }
        

        stage ('Build and Test Application') {
            steps {
                sh "go build"
                sh 'go test ./...'
            }
        }
        
        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv(credentialsId: 'jenkins-sonar-token') {
                        sh 'sonar-scanner'
                    }
                }
            }
        }
        
        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${env.DOCKER_IMAGE_NAME} ."
            }
        }
        stage("Push Image"){
            withDockerRegistry(credentialsId: 'docker-hub', url: 'https://index.docker.io/v1/') {
                sh "docker push ${env.DOCKER_IMAGE_NAME}"
            }
        }
    }

    post {
        success {
            echo 'Kiểm tra chất lượng mã nguồn thành công!'
        }
        failure {
            echo 'Kiểm tra chất lượng mã nguồn thất bại. Hãy kiểm tra lại.'
        }
    }
    
}
