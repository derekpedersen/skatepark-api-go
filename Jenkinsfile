node {
    stage('Checkout') {
        dir('../../src/github.com/derekpedersen/skatepark-api-go') {
            checkout scm
        }
    }
    stage('Build') {
        withCredentials([[$class: 'StringBinding', credentialsId: 'PROJECT_ID', variable: 'GCLOUD_PROJECT_ID']]) {
            dir('../../src/github.com/derekpedersen/skatepark-api-go') {
                sh 'go get github.com/golang/mock/gomock && \
                    go install github.com/golang/mock/mockgen && \
                    make kubernetes'
            }
        }
    }
}