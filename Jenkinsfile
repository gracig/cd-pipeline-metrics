//Este job foi cadastrado no Jenkins através do arquivo de inicialização: 150-job-hello.groovy
pipeline {
	//O agente slave foi cadastrado no Jenkins através do arquivo de inicialização: 005-set-kubernetes.groovy
	agent {label "slave"}

	environment {
		//NEXUS_REPOSITORY Hardcoded. O Nexus inicia com esse repositório.//TODO: Criar uma variável de ambiente
		NEXUS_REPOSITORY   = "devops"
		NEXUS_PROTOCOL     = "http"
		//NEXUS_URL definida no arquivo de inicialização do Jenkins: 010-environment-variables.groovy
		NEXUS_URL          = "${NEXUS_REPO_MAVEN}"
		NEXUS_VERSION      = "nexus3"
		//NEXUS_CREDENTIAL definida no arquivo de inicialização do Jenkins: 003-set-credentials.groovy
		NEXUS_CREDENTIAL   = "nexus"
		//SONARQUBE_ENDPOINT Definida no arquivo de inicialização do Jenkins: 010-environment-variables.groovy
		SONARQUBE_ENDPOINT = "${SONARQUBE_HTTP_ENDPOINT}"

		GO_APP_PATH        = "${env.HOME}/go/src/bitbucket.org/gracig"
		APP_NAME           = "cdrc"
		APP_VERSION        = "1.0.0"
		GO_APP_FULLPATH    = "${GO_APP_PATH}/${APP_NAME}"
		APP_BIN            = "${APP_NAME}"
	}

	stages {
		stage ('Verificação Ambiente'){
			steps {
				sh 'printenv'
			}
		}
		stage ('Compilação') {
			steps {
				echo 'Prepara ambiente go e compila a ferramenta'
				sh '''
					mkdir -p $GO_APP_PATH
					ln -s `pwd` $GO_APP_FULLPATH
					cd $GO_APP_FULLPATH
					VERSION=$APP_VERSION TOOLNAME=$APP_BIN make
				'''
			}
		}
		stage('Publicação') {
			steps {
				echo " Publicação do artefato"
				echo "Publica os artefatos da aplicação ${APP_NAME} no repositório: ${NEXUS_REPOSITORY}"
				script {
					nexusArtifactUploader nexusUrl: NEXUS_URL,
						artifacts: [
							[ artifactId: APP_NAME, classifier: '', file: APP_BIN ]
						],
						credentialsId: NEXUS_CREDENTIAL,
						nexusVersion: NEXUS_VERSION,
						protocol: NEXUS_PROTOCOL,
						repository: NEXUS_REPOSITORY,
						version: APP_VERSION
				}
			}
		}
	}
}
