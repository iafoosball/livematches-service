pipeline {
    agent any

    stages {
        stage ("Build") {
            environment {
                    PORT = 9004
                    DB_KEY = credentials('arangoMatchesProd')
            }
            steps{
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