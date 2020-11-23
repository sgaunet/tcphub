#!/bin/sh

if [ "$SERVER" = "" ]
then
  echo "\$SERVER not initialized"
  exit 1
fi

if [ "$PORT" = "" ]
then
  echo "\$PORT not initialized"
  exit 1
fi


exec $@ -s $SERVER -p $PORT