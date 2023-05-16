pipeline {
    agent any

    stages {
        stage('Test') {
            tools {
                go 'go-1.20.3'
            }
            environment {
                GO111MODULE = 'on'
            }
            agent {
                label "docker"
            }
            steps {
                sh 'go test'
            }
        }
        stage('Docker login') {
            agent {
                label "docker"
            }
            steps {
                sh 'ls -la'
                sh 'pwd'
                sh 'aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 282335569253.dkr.ecr.us-east-1.amazonaws.com'
            }
        }
        stage('Build image'){
            agent {
                label "docker"
            }
            steps{
                sh "docker build -t 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo . --no-cache"
            }
        }
        stage('Tag image'){
            agent {
                label "docker"
            }
            steps{
                sh """
                   docker tag 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:latest 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:${env.BUILD_NUMBER}
                """                
            }
        }
        stage('Push image'){
            agent {
                label "docker"
            }
            steps{
                sh """
                    docker push 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:latest
                    docker push 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:${env.BUILD_NUMBER}
                """
            }
        }
    }
    post{
        always {
            node('docker'){
                sh 'docker image prune -af'
                sh 'docker images'
                sh 'docker logout'
            }
        }
    }
}