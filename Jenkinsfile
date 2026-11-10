pipeline {
    agent {
        label 'build-jenkins-golang || build-golang || build-golang-stable'
    }

    options {
        skipDefaultCheckout true
        timeout(time: 30, unit: 'MINUTES')
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
    }
}
