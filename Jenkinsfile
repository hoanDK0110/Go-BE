pipeline {
    agent any

    stages {
        stage('Checkout') {
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
    }
    
}
