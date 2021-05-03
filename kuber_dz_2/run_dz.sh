helm install glebpg \
--set postgresqlUsername=gleb,postgresqlPassword=secretpassword,postgresqlDatabase=kuber-dz-2 \
bitnami/postgresql
