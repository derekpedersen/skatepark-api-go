pipeline {
    agent {
        label 'build-golang-stable'
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
                withCredentials([file(credentialsId: 'k8s_json', variable: 'k8s_json')]) {
                    dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                        sh "cp \$k8s_json k8s_json.json"
                        sh 'gcloud auth activate-service-account odin-1985 --key-file=k8s_json.json'
                        sh 'make deploy'
                    }
                }
            }
        }
    }
    // post {
    //     always {
    //         withCredentials([[$class: 'StringBinding', credentialsId: 'SKATEPARK_API_COVERALLS_TOKEN', variable: 'COVERALLS_TOKEN']]) {
    //             dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
    //                 step([$class: 'CoberturaPublisher', autoUpdateHealth: false, autoUpdateStability: false, coberturaReportFile: '**/cp.xml', failUnhealthy: false, failUnstable: false, maxNumberOfBuilds: 0, onlyStable: false, sourceEncoding: 'ASCII', zoomCoverageChart: false]) 
    //                 sh 'go get github.com/derekpedersen/goveralls'
    //                 sh 'goveralls -coverprofile=cp.out'
    //             }
    //         }
    //     }
    // }
}