pipeline {
    agent any
    environment {
        COMPOSE_FILE = "docker-compose.yml"
    }

    stages {
        stage ("Build") {
            steps{
                sh "docker stop livematches-service &"
                sh "docker rm livematches-service &"
                sh "sleep 15s"
                sh "docker-compose build --pull"

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