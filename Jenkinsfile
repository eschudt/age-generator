node {
    def app

    stage('Clone repository') {
        checkout scm
    }

    stage('Test and Build image') {
        steps {
          sh 'make test'
        }
        app = docker.build("eschudt/age-generator")
    }

    stage('Integration Test') {
        app.inside {
            sh 'echo "Tests passed"'
        }
    }

    stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("0.0.${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }

    stage('Deploy - Dev') {
        steps {
            echo 'Deployed to Dev'
        }
    }
}