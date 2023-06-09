def ECR_URL 
def HASH_COMMIT

pipeline {
    agent any

     stages {

        stage('Run unit test/coverage') {
            when { anyOf { branch 'main'; branch 'dev' } }
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
            when { anyOf { branch 'main'; branch 'dev' } }       
            agent {
                label "docker"
            }
            steps {
                withSonarQubeEnv("sonarqube-9.9.1"){
                    sh "/home/ec2-user/install_scanner/sonar-scanner-4.8.0.2856-linux/bin/sonar-scanner"
                }
            }
        }
        stage("Quality Gate") {
            when { anyOf { branch 'main'; branch 'dev' } }
            steps{
                timeout(time: 1, unit: 'HOURS') {
                waitForQualityGate abortPipeline: true
                }
            }
        }
        stage('Docker login') {
            when { anyOf { branch 'dev' } }
            agent {
                label "docker"
            }
            steps {
                sh 'aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 282335569253.dkr.ecr.us-east-1.amazonaws.com'
            }
        }
        stage('Get ecr url and hash commit'){
            when { anyOf { branch 'dev' } }
            agent {
                label "terraform"
            }
            steps{
                dir("terraform/global"){
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
            when { anyOf { branch 'dev' } }
            agent {
                label "docker"
            }
            steps{
                sh "docker build -t  ${ECR_URL} . --no-cache"
            }
        }
        stage('Tag image'){
            when { anyOf { branch 'dev' } }
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
            when { anyOf { branch 'dev' } }
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
        stage('Deploy to staging'){
            when { anyOf { branch 'dev' } }
            agent {
                label "terraform"
            }
            steps{
               dir("terraform/staging/"){
                    sh """
                    terraform init
                    terraform apply -var='image_tag=latest' -auto-approve
                    aws ecs update-service --region us-east-1 --cluster staging-cluster --service staging-service --task-definition 'staging-td'  --force-new-deployment
                    """                   
                    script {
                        STAGING_DNS = sh (
                          script: "terraform output --raw staging_lb",
                          returnStdout: true
                        )
                    }
                    sh "echo ${STAGING_DNS}"
               }
            }
        }
        stage('Deploy to production'){
            when { anyOf { branch 'main' } }
            agent {
                label "terraform"
            }
            steps{
               dir("terraform/production/"){
                    sh """
                    terraform init
                    terraform apply -var='image_tag=latest' -auto-approve
                    aws ecs update-service --region us-east-1 --cluster production-cluster --service production-service --task-definition 'production-td'  --force-new-deployment
                    """                   
                    script {
                        PRODUCTION_DNS = sh (
                          script: "terraform output --raw production_lb",
                          returnStdout: true
                        )
                    }
                    sh "echo ${PRODUCTION_DNS}"
               }
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
