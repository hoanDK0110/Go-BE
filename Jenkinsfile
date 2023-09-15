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
                checkout SCM
            }
        }

        stage('Build Application') {
            steps {
                sh "${GO_PATH}/go build" // Sử dụng biến môi trường GO_PATH để tham chiếu đến lệnh Go
            }
        }
        
        stage('Test Application') {
            steps {
                sh "${GO_PATH}/go test ./..." // Sử dụng biến môi trường GO_PATH để tham chiếu đến lệnh Go
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
