pipeline {
    agent any

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
        

        stage ('Build Application') {
            steps {
                sh "go build"
            }
        }

        stage ('Test Application') {
            steps {
                sh "mvn test"
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
        
        stage("Push Image"){
            withDockerRegistry(credentialsId: 'docker-hub', url: 'https://index.docker.io/v1/') {
                // some block
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
