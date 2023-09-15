pipeline {
    agent any
    tools {
        go 'golang'
    }
    environment {
        APP_NAME = 'app-pipeline'
        RELEASE = "1.0.0"
        DOCKER_USER = "hoanDK0110"
        DOCKER_PASS = 'Dokimhoan2001'
        IMAGE_NAME = "${DOCKER_USER}" + "/" + "${APP_NAME}"
        IMAGE_TAG = "${RELEASE}-${BUILD_NUMBER}"
        DOCKER_IMAGE_NAME = 'golang-web:1.0'
        JENKINS_API_TOKEN = credentials("jenkins-sonar")
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
                sh "go build"
            }
       }

        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'make unit-tests'
            }
        }

        stage("functional-test") {
            steps {
                echo 'FUNCTIONAL TEST EXECUTION STARTED'
                sh 'make functional-tests'
            }
        }
        
        
        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv(credentialsId: 'jenkins-sonar') {
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
