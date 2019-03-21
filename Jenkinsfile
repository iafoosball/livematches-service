pipeline {
    agent any

    stages {
    /*
        stage ("Build Stag") {
                steps{
                    sh "docker-compose -f docker-compose.stag.yml build --pull"
                }
            }
            stage ("Remove old Stag") {
                steps {
                   sh "docker stop livematches-service-stag &"
                   sh "docker rm livematches-service-stag &"
                   sh "sleep 15s"
                }
            }
            stage ("Deploy Stag") {
                steps {
                    sh "docker-compose -p livematches-stag -f docker-compose.stag.yml up -d"
                }
            }
*/
        stage ("Build Production") {
            steps{
                sh "docker-compose -f docker-compose.prod.yml build --pull"
            }
        }
        stage ("Remove old Production") {
            steps {
               sh "docker stop livematches-service-prod &"
               sh "docker rm livematches-service-prod &"
            }
        }
        stage ("Production") {
            steps {
                /* sh "docker-compose  -p livematches-prod -f docker-compose.prod.yml up" */
            }
        }
    }
    post {
        always {
            sh "docker-compose -f docker-compose.prod.yml  down -v --rmi 'all'"
        }
    }
}
