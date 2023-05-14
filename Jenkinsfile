pipeline {
    agent any

    stages {
        stage('Hello') {
            agent {
                label "docker"
            }
            steps {
                echo 'Hello World'
                sh 'pwd'
                sh 'whoami'
                sh 'ls -la'
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
                sh '''
                   docker rmi -f $(docker images -a -q)
                   docker build -t final-demo/${env.BUILD_NUMBER} .
                '''
            }
        }
        stage('Tag image'){
            agent {
                label "docker"
            }
            steps{
                sh '''
                   docker tag finaldemo/${env.BUILD_NUMBER} 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:latest
                   docker tag finaldemo/${env.BUILD_NUMBER} 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:${env.BUILD_NUMBER}
                '''                
            }
        }
        stage('Push image'){
            agent {
                label "docker"
            }
            steps{
                sh '''
                    docker push --all-tags 282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo
                '''
            }
        }
    }
    post{
        always {
            node('docker'){
                sh 'docker rmi -f $(docker images -a -q)'
                sh 'docker images'
                sh 'docker logout'
            }
        }
    }
}
