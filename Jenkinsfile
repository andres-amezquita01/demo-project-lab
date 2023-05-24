def ECR_URL = "282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo"
def STAGING_USER = "ec2-user@ec2-3-238-154-128.compute-1.amazonaws.com"
def DEPLOYMENT_USER =  "ec2-user@ec2-107-21-72-235.compute-1.amazonaws.com"
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
                sh 'echo ${ECR_URL}'
                //sh "docker build -t  ${ECR_URL} . --no-cache"
            }
        }
       stage('Get ecr'){
            agent {
                label "terraform"
            }
            steps{
                dir("terraform/remote_backend"){
                    sh 'echo ${ECR_URL}'
                    sh 'terraform output'
                }
                //sh "docker build -t  ${ECR_URL} . --no-cache"
            }
        }
//         stage('Tag image'){
//             agent {
//                 label "docker"
//             }
//             steps{
//                 sh """
//                    docker tag  ${ECR_URL}:latest ${ECR_URL}:${env.BUILD_NUMBER}
//                 """
//             }
//         }
//         stage('Push image'){
//             agent {
//                 label "docker"
//             }
//             steps{
//                 sh """
//                     docker push ${ECR_URL}:latest
//                     docker push ${ECR_URL}:${env.BUILD_NUMBER}
//                 """
//             }
//         }
//         stage('Deploy to stage'){
//             steps{
//                 sh """
//                 scp docker.sh ${STAGING_USER}:~/stage
//                 """
//             }
//         }
//         stage('Deploy to production'){
//             steps{
//                 input(message: '¿Do you want to deploy to production?', ok: 'yes')
//
//                 sh """
//                 scp docker.sh ${DEPLOYMENT_USER}:~/production
//                 """
//             }
//         }
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
