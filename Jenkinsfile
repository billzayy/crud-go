pipeline {
  agent any

  tools { go '1.21.3' }
    
  stages {
    // stage('Docker Compose Down Previous'){
    //     steps{
    //         sh 'docker-compose down'
    //     }
    // }
    stage('Git Checkout') {
        steps {
            checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/billzayy/crud-go']])
        }
    }
    stage('Test Version') {
      steps {
        sh 'go version'
      }
    }
    stage('Check Docker'){
        steps{
            sh 'docker version'
        }
    }
    stage('Build & Push Images'){
        steps{
            script{
                withDockerRegistry(credentialsId: 'billzay-docker', toolName: 'Docker', url: 'https://index.docker.io/v1/') {
                    sh 'docker build -t coderbillzay/crud-go .'
                    sh 'docker push coderbillzay/crud-go'
                }
            }
        }
    }
    stage('Docker Run'){
        steps{
            sh 'docker-compose up -d'
        }
    }
  }
}