pipeline {
    agent {
        label 'build-golang-stable'
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
                withCredentials([[$class: 'StringBinding', credentialsId: 'GCLOUD_PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
                    dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                        sh 'make deploy'
                    }
                }
            }
        }
    }
    post {
        always {
            step([$class: 'CoberturaPublisher', autoUpdateHealth: false, autoUpdateStability: false, coberturaReportFile: '/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go/cp.xml', failUnhealthy: false, failUnstable: false, maxNumberOfBuilds: 0, onlyStable: false, sourceEncoding: 'ASCII', zoomCoverageChart: false])
        }
    }
}