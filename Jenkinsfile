def ECR_URL = "282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo"
def DEPLOYMENT_USER =  "ec2-user@ec2-54-160-222-29.compute-1.amazonaws.com"
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
                sh "docker build -t  ${ECR_URL} . --no-cache"
            }
        }
        stage('Tag image'){
            agent {
                label "docker"
            }
            steps{
                sh """
                   docker tag  ${ECR_URL}:latest ${ECR_URL}:${env.BUILD_NUMBER}
                """                
            }
        }
        stage('Push image'){
            agent {
                label "docker"
            }
            steps{
                sh """
                    docker push ${ECR_URL}:latest
                    docker push ${ECR_URL}:${env.BUILD_NUMBER}
                """
            }
        }
        stage('Deploy to production'){
            steps{
                sh """
                scp docker.sh ${DEPLOYMENT_USER}:~/production
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
