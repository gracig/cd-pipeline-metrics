export VERSION=1.0.0
export CUSTOMER=vivo

#Global Environment
#export JENKINS_API_TOKEN=53b18010f6a085953cbd7c4572465c11
export JENKINS_API_TOKEN=password
export JENKINS_API_USER=admin
export INFLUXDB_URL=http://influxdb.framework-prod.svc.cluster.local:8086
export INFLUXDB_DB=pipeline
export INFLUXDB_USER=admin
export INFLUXDB_PWD=admin

#Inside Pipeline
export JENKINS_URL=http://jenkins.local-test.svc.cluster.local:8080/
export BUILD_URL=http://jenkins.local-test.svc.cluster.local:8080/job/hello_01_commit/65
export BUILD_TIMESTAMP=2016-11-29T09:00:19-05:00
export BUILD_NUMBER=65
export JOB_NAME=hello_01_commit

#Inside SONAR
export SONAR_URL=http://10.146.80.204:9000
export SONAR_USER=admin
export SONAR_PWD=admin
export SONAR_PROJECT_KEY=br.com.tim:devops-tim-rede-inge-cdr-gprs_ericsson
export SONAR_METRICS=tests,coverage,ncloc,lines,violations,bugs,reliability_rating,security_rating,vulnerabilities,complexity,function_complexity,code_smells,sqale_rating,sqale_index,sqale_debt_ratio



make build

./devops_metrics-${VERSION} -jkwfapi
