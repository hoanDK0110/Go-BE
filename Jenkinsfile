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
                script {
                    git branch: 'main', url: 'https://github.com/hoanDK0110/Go-BE.git'
                }
            }
        } 
    }
}
