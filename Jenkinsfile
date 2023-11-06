pipeline {
    agent { 
        label 'bazel-agent'
    }
    stages {
        stage('build') {
            steps {
                sh 'bazelisk --version'
            }
        }
    }
}
