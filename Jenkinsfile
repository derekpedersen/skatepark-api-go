node {
    stage('Checkout') {
        dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
            checkout scm
        }
    }
    stage('Build') {
        dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
            sh 'make build'
        }
    }
    stage('Test') {
        dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
            sh 'go get github.com/golang/mock/gomock && \
                go install github.com/golang/mock/mockgen && \
                make test'
        }
    }
    stage('Publish') {
        if (env.BRANCH_NAME == 'master') {
            withCredentials([[$class: 'StringBinding', credentialsId: 'PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make publish'
                }
            }
        }
    }
    stage('Deploy') {
        if (env.BRANCH_NAME == 'master') {
            withCredentials([[$class: 'StringBinding', credentialsId: 'PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
                dir('/root/workspace/go/src/github.com/derekpedersen/skatepark-api-go') {
                    sh 'make deploy'
                }
            }
        }
    }
}