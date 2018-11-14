ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECTNAME=gohound
DOCKERNAME=${PROJECTNAME}-mongoserver

mongoserver-start: mongoserver-stop
	docker create --rm --name ${DOCKERNAME} -p 47017:27017 -e MONGO_INITDB_DATABASE=${PROJECTNAME} mongo:4
	docker start ${DOCKERNAME}

mongoserver-stop:
	-docker stop ${DOCKERNAME}