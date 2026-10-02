# Overview

This is a simple chat server

#  Example Usage

curl -d '{ "message": {"text":"Hows it going", "from":"Elaine"}, "to":"Joe"}' http://localhost:8080/send

curl -d '{ "message": {"text":"hi", "from":"Jack"}, "to":"Joe"}' http://localhost:8080/send

curl http://localhost:8080/check?user=Joe