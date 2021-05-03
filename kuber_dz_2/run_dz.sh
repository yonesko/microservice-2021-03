echo "installing Postgresql.."

helm install glebpg \
--set persistenceEnabled=false,postgresqlUsername=gleb,postgresqlPassword=secretpassword,postgresqlDatabase=kuber-dz-2 \
bitnami/postgresql

echo "installing App.."
helm install glebio-users-crud chart


#arch.homework curls
#curl -s -X POST -d'{"username":"1sFT0I0m67","firstName":"6M1B1dtUN","lastName":"3oEHpLF79","email":"76Hdp","phone":"b4PYX72rE2"}'