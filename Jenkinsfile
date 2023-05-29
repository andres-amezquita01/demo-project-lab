def ECR_URL 
def HASH_COMMIT
def STAGING_USER = "ec2-user@ec2-3-238-105-179.compute-1.amazonaws.com"
def DEPLOYMENT_USER =  " ec2-user@ec2-3-91-151-183.compute-1.amazonaws.com"
pipeline {
    agent any

     stages {

        stage('Run unit test/coverage') {
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
                sh 'go test -v -coverprofile cover.out'
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
                    sh "cat sonar-project.properties"
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
                     script {
                        HASH_COMMIT = sh (
                          script: "git log -1 --pretty=format:'%H'",
                          returnStdout: true
                        )
                      }
                    sh "echo ${ECR_URL}"
                    sh "echo ${HASH_COMMIT}"

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
                   docker tag  ${ECR_URL}:latest ${ECR_URL}:${HASH_COMMIT}
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
                    docker push ${ECR_URL}:${HASH_COMMIT}
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
