pipeline {
    agent any
    environment {
        DOCKER_IMAGE_NAME = 'golang-web:1.0'
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

        
        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv(credentialsId: 'sonarqube') {
                        sh 'sonar:sonar'
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
