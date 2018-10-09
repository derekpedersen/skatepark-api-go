node {
    stage('Checkout') {
        dir('../../src/github.com/derekpedersen/skatepark-api-go') {
            checkout scm
        }
    }
    stage('Build') {
        dir('../../src/github.com/derekpedersen/skatepark-api-go') {
            sh 'make build'
        }
    }
    stage('Test') {
        dir('../../src/github.com/derekpedersen/skatepark-api-go') {
            sh 'go get github.com/golang/mock/gomock && \
                go install github.com/golang/mock/mockgen && \
                make test'
        }
    }
    stage('Publish') {
        if (env.BRANCH_NAME == 'master') {
            withCredentials([[$class: 'StringBinding', credentialsId: 'PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
                dir('../../src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make publish'
                }
            }
        }
    }
    stage('Deploy') {
        if (env.BRANCH_NAME == 'master') {
            withCredentials([[$class: 'StringBinding', credentialsId: 'PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
                dir('../../src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make deploy'
                }
            }
        }
    }
}