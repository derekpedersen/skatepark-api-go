pipeline {
    agent {
        label 'build-golang-stable'
    }
    options {
    skipDefaultCheckout true
    }
    stages {
        stage('Environment') {
            steps {
                sh 'go version'
                sh 'GO111MODULE=on go get github.com/golang/mock/mockgen'
            }
        }
        stage('Checkout') {
            steps{
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    checkout scm
                }
            }
        }
        stage('Build') {
            steps{
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make build'
                }
            }
        }
        stage('Swagger') {
            steps {
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make swagger'
                }
            }
        }
        stage('Test') {
            steps {
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make test'
                }
            }
        }
        stage('Docker') {
            steps {    
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make docker'
                }
            }
        }
        stage('Publish') {
            when {
                expression { env.BRANCH_NAME == 'master' }
            }
            steps {
                withCredentials([[$class: 'StringBinding', credentialsId: 'GCLOUD_PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
                    dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                        sh 'make publish'
                    }
                }
            }
        }
        stage('Deploy') {
            when {
                expression { env.BRANCH_NAME == 'master' }
            }
            steps {    
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make deploy'
                }
            }
        }
    }
}