pipeline {
    agent any

    stages {
        stage ("Build") {
            steps{
                sh "docker-compose -f docker-compose.prod.yml build --pull"
            }
        }
        stage ("Remove old") {
            steps {
               sh "docker stop livematches-service-prod &"
               sh "docker rm livematches-service-prod &"
               sh "sleep 15s"
            }
        }
        stage ("Production") {
            steps {
                sh "docker-compose -f docker-compose.prod.yml up"
            }
        }
    }
    post {
        always {
            sh "docker-compose -f docker-compose.prod.yml  down -v --rmi 'all'"
        }
    }
}
