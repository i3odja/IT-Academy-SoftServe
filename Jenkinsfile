pipeline {
    agent { docker { image 'golang:1.14' } }
    environment {
        GOCACHE = '/tmp/'
    }
    stages {
        stage('Testing') {
            steps {
                sh 'go test'
            }
        }
    }
}
