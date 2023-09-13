pipeline {
    agent any

    stages {
        stage('Clone code') {
            steps {
                // Sao chép mã nguồn từ GitHub
                git branch: 'main', url: 'https://github.com/hoanDK0110/Go-BE.git'
            }
        }
        
        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv('SonarQubeServer') {
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
    
}
