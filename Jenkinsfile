pipeline {
    agent {
        label 'build-jenkins-golang || build-golang || build-golang-stable'
    }

    options {
        skipDefaultCheckout true
    }

    environment {
        PROJECT_DIR = '.'
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
                script {
                    def dockerhubCredentialsId = env.DOCKERHUB_CREDENTIALS_ID ?: 'docker-jenkins-pat'
                    withDockerRegistry([credentialsId: dockerhubCredentialsId, url: 'https://index.docker.io/v1/']) {
                        dir("${env.PROJECT_DIR}") {
                            sh 'make publish'
                        }
                    }
                }
            }
        }

        stage('Deploy') {
            when {
                expression { env.BRANCH_NAME == 'main' }
            }
            steps {
                script {
                    def doksKubeconfigCredentialsId = env.DOKS_KUBECONFIG_CREDENTIALS_ID ?: 'DOKS_KUBECONFIG'
                    withCredentials([
                        file(
                            credentialsId: doksKubeconfigCredentialsId,
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
}
