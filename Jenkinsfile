pipeline {
    agent { 
        label 'bazel-agent'
    }
    stages {
        stage('build') {
            steps {
                sh 'bazelisk run //services/app1:push'
            }
        }
    }
}
