pipeline {
    agent any
    environment {
        APP_NAME = 'app-pipeline'
        RELEASE = "1.0.0"
        DOCKER_USER = "hoanDK0110"
        DOCKER_PASS = 'Dokimhoan2001'
        IMAGE_NAME = "${DOCKER_USER}" + "/" + "${APP_NAME}"
        IMAGE_TAG = "${RELEASE}-${BUILD_NUMBER}"
        DOCKER_IMAGE_NAME = 'golang-web:1.0'
        JENKINS_API_TOKEN = credentials("jenkins-sonar")
        GO_PATH = "/usr/local/go/bin" // Đường dẫn đầy đủ đến thư mục chứa lệnh Go
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
                script {
                    git branch: 'main', url: 'https://github.com/hoanDK0110/Go-BE.git'
                }
            }
        }

        stage("Build Application"){
            steps {
                sh "${GO_PATH}/go build"
            }
        }
        
        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv(installationName: 'sonarqube', credentialsId: 'jenkins-sonar') {
                        sh 'mvn clean package sonar:sonar'
                    }
                }
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
