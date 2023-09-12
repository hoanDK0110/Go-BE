pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Sao chép mã nguồn từ GitHub
                checkout scm
            }
        }

        stage('Build') {
            steps {
                // Xây dựng mã nguồn Golang
                sh 'go build'
            }
        }

        stage('Unit Test') {
            steps {
                // Thực hiện kiểm tra đơn vị (unit test)
                sh 'go test ./...'
            }
        }
    }

    post {
        success {
            // Hành động sau khi pipeline chạy thành công
            echo 'Pipeline ran successfully'
        }
        failure {
            // Hành động sau khi pipeline thất bại
            echo 'Pipeline failed'
        }
    }
}
