pipeline {
    agent any

    stages {
        stage ("Build") {
            environment {
                    SERVICE_PORT='9005'
                    DB_KEY=credentials('arangoMatchesProd')
            }
            steps{
                sh 'export SERVICE_PORT=9006'
                sh 'printenv'
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