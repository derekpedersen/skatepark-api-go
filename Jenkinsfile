pipeline {
    agent {
        label 'k8s'
    }
    stages {
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
        stage('Test') {
            steps {
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    // sh 'go get github.com/golang/mock/gomock && \
                    //     go install github.com/golang/mock/mockgen && \
                    //     make test'
                    //sh 'make test'
                }
            }
        }
        stage('Publish') {
            // when {
            //     expression { env.BRANCH_NAME == 'master' }
            // }
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
                withCredentials([[$class: 'StringBinding', credentialsId: 'GCLOUD_PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
                    dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                        sh 'make deploy'
                    }
                }
            }
        }
    }
}