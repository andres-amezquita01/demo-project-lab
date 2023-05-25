def ECR_URL 
def STAGING_USER = "ec2-user@ec2-3-237-95-49.compute-1.amazonaws.com"
def DEPLOYMENT_USER =  " ec2-user@ec2-23-20-80-13.compute-1.amazonaws.com"
pipeline {
    agent any

     stages {
        stage('Run unit test') {
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
        stage('Run sonarqube') {
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
                withSonarQubeEnv("sonarqube-9.9.1"){
                    sh "/home/ec2-user/install_scanner/sonar-scanner-4.8.0.2856-linux/bin/sonar-scanner"
                }
            }
        }

        stage('Docker login') {
            agent {
                label "docker"
            }
            steps {
                sh 'aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 282335569253.dkr.ecr.us-east-1.amazonaws.com'
            }
        }
        stage('Get ecr url'){
            agent {
                label "terraform"
            }
            steps{
                dir("terraform/remote_backend"){
                    sh 'terraform init'
                     script {
                        ECR_URL = sh (
                          script: "terraform output --raw ecr_repository_url",
                          returnStdout: true
                        )
                      }
                    sh "echo ${ECR_URL}"
                }
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
        stage('Deploy to stage'){
            steps{
                sh """
                scp docker.sh ${STAGING_USER}:~/stage
                """
            }
        }
        stage('Deploy to production'){
            steps{
                input(message: '¿Do you want to deploy to production?', ok: 'yes')
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
