pipeline {
    agent any
    stages {
        stage('Cleanup Workspace') {
            steps {
                cleanWs()
            }
        }
        // Clone code from github repo
        stage('Clone code') {
            steps {
                script {
                    git branch: 'main', url: 'https://github.com/hoanDK0110/Go-BE.git'
                }
            }
        } 
    }
}
