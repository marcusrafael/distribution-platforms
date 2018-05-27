#!/bin/bash

# Iniciar nmon
idNmon=`nmon -f -s 1 -c -1 -m "." -p`
sleep 10
#echo "Nmon está coletando!"

#Iniciar serviço de nomes em GO.
go run /home/xavier/Xavier/distribution-platforms/assignment-04/src/github.com/marcusrafael/distribution-platforms/app/name-server/names.go&
#echo "Servidor de nomes iniciado!"
sleep 2

#Iniciar servidor em GO.
go run /home/xavier/Xavier/distribution-platforms/assignment-04/src/github.com/marcusrafael/distribution-platforms/app/gcp-server/gcp.go&
#echo "Servidor iniciado!!"
sleep 2

#echo "Software está começando a processar!"
#Quantidade de experimentos.
go run /home/xavier/Xavier/distribution-platforms/assignment-04/src/github.com/marcusrafael/distribution-platforms/app/scenarios/client.go

killall names
killall gcp
killall nmon
