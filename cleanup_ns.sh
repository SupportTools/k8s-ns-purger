#!/bin/bash

if [[ -z $PURGE_AGE ]]
then
  PURGE_AGE="8 hour"
fi

echo "Purge Age: $PURGE_AGE"

kubectl get ns -l ns-purge=true -o go-template --template '{{range .items}}{{.metadata.name}},{{.metadata.creationTimestamp}}{{"\n"}}{{end}}' | while IFS=, read -r namespace timestamp
do
  echo "#################################################################"
  echo "Namespace: $namespace"
  echo "Timestamp: $timestamp"
  deletetime=`date -d "$PURGE_AGE ago" -Ins --utc | sed 's/+0000/Z/'`
  if [[ "$timestamp" < "$deletetime" ]]
  then
    echo "Need to delete"
    if [[ -z $namespace ]]
    then
      echo "Namespace is empty for some reason"
    else
      echo "Deleteing..."
      kubectl delete ns "$namespace"
    fi
  fi
  echo "#################################################################"
done
