pipeline {
    agent {
        label 'build-jenkins-golang || build-golang || build-golang-stable'
    }

    options {
        skipDefaultCheckout true
        timestamps()
    }

    environment {
        PROJECT_DIR = '.'
        DOCKERHUB_CREDENTIALS_ID = 'DOCKERHUB_CREDENTIALS'
        DOKS_KUBECONFIG_CREDENTIALS_ID = 'DOKS_KUBECONFIG'
    }

    stages {
        stage('Checkout') {
            steps {
                dir("${env.PROJECT_DIR}") {
                    checkout scm
                }
            }
        }

        stage('Dependencies') {
            steps {
                dir("${env.PROJECT_DIR}") {
                    sh 'make dependencies'
                }
            }
        }

        stage('Test') {
            steps {
                dir("${env.PROJECT_DIR}") {
                    sh 'make test'
                }
            }
        }

        stage('Build') {
            steps {
                dir("${env.PROJECT_DIR}") {
                    sh 'make build'
                }
            }
        }

        stage('Docker') {
            steps {
                dir("${env.PROJECT_DIR}") {
                    sh 'make docker'
                }
            }
        }

        stage('Publish') {
            when {
                expression { env.BRANCH_NAME == 'main' }
            }
            steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: "${env.DOCKERHUB_CREDENTIALS_ID}",
                        usernameVariable: 'DOCKERHUB_USERNAME',
                        passwordVariable: 'DOCKERHUB_TOKEN'
                    )
                ]) {
                    dir("${env.PROJECT_DIR}") {
                        sh 'echo "$DOCKERHUB_TOKEN" | docker login --username "$DOCKERHUB_USERNAME" --password-stdin'
                        sh 'make publish'
                        sh 'docker logout || true'
                    }
                }
            }
        }

        stage('Deploy') {
            when {
                expression { env.BRANCH_NAME == 'main' }
            }
            steps {
                withCredentials([
                    file(
                        credentialsId: "${env.DOKS_KUBECONFIG_CREDENTIALS_ID}",
                        variable: 'KUBECONFIG'
                    )
                ]) {
                    dir("${env.PROJECT_DIR}") {
                        sh 'make deploy'
                    }
                }
            }
        }
    }
}
