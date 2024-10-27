#!/bin/bash
buildshop(){
    ps -ef | grep shop | grep -v grep  |grep root | cut -c 9-16  | xargs kill -s 9
    rm -rf shop
    go build -o shop ./
    rm -rf ./controllers
    rm -rf ./model
    rm -rf ./routers
    rm -rf main.go
    nohup ./shop >/dev/null 2>&1 &
}

strart(){
nohup ./shop >/dev/null 2>&1 &   
}

buildshop
#strart
