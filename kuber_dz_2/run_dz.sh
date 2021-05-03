#echo "installing Postgresql.."
#
#helm install glebpg \
#  --set persistenceEnabled=false,postgresqlUsername=gleb,postgresqlPassword=secretpassword,postgresqlDatabase=kuber-dz-2 \
#  bitnami/postgresql
#
#echo "installing App.."
#helm install glebio-users-crud chart

echo "Test  REST Api by curl.."
sleep 1
curl -s arch.homework/api/v1/
echo
sleep 1
user_id=$(curl -s arch.homework/api/v1/user -d'{"username":"1sFT0I0m67","firstName":"6M1B1dtUN","lastName":"3oEHpLF79","email":"76Hdp","phone":"b4PYX72rE2"}')
echo "inserted user #$user_id"
sleep 1
echo "find user #$user_id"
if command -v jq &>/dev/null; then
  curl -s arch.homework/api/v1/user/"$user_id" | jq
else
  curl -s arch.homework/api/v1/user/"$user_id"
  echo
fi
