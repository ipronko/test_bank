#!/usr/bin/env bash
#Create new account
echo "create first account"
curl --request POST --data '{"balance":1000}' http://localhost:8080/
echo -e "\ncreate second account"
curl --request POST --data '{"balance":0}' http://localhost:8080/

#Show balance of accounts
echo -e "\n\nshow balance in first account"
curl --request GET --data '{"id":1}' http://localhost:8080/
echo -e "\nshow balance in second account"
curl --request GET --data '{"id":2}' http://localhost:8080/


#Transfer
echo -e "\n\ntransfer 10 from first to second"
curl --request PATCH --data '{"from":1, "to":2, "sum":10}'   http://localhost:8080/ 

#Show balance of accounts
echo -e "\n\nshow balance in first account"
curl --request GET --data '{"id":1}' http://localhost:8080/
echo -e "\nshow balance in second account"
curl --request GET --data '{"id":2}' http://localhost:8080/
