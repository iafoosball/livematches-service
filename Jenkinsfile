pipeline {
    agent any

    stages {
        stage ("Build") {
        environment {
                    PORT = 8000
                }
            steps{
                sh "echo $PORT"
                sh "docker-compose build --pull"
            }
        }
        stage ("Remove old") {
            steps {
               sh "docker stop livematches-service &"
               sh "docker rm livematches-service &"
               sh "sleep 15s"
            }
        }
        stage ("Production") {
            steps {
                sh "docker-compose up"
            }
        }
    }
    post {
        always {
            sh "docker-compose down -v --rmi 'all'"
        }
    }
}