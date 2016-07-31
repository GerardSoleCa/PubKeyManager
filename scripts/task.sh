#!/usr/bin/env bash

# Output colors
NORMAL='\033[0m' # No Color
RED='\033[0;31m'
BLUE='\033[0;34m'
DIR=`pwd`

# INTERNAL USAGE
log() {
    echo -e "${BLUE}${1}${NORMAL}"
}

# INTERNAL USAGE
error() {
    echo -e "${RED}ERROR - ${1}${NORMAL}"
    return -1
}

setup(){
    log "Setting up project"
    cd $DIR/public
    npm install
    bower install
    cd $DIR
    log "Setting up finished"
}

build(){
    log "Building go project"
    go build -v
    log "Building finished"
}

gruntfiles(){
    log "Grunting frontend files"
    cd $DIR/public
    grunt
    cd $DIR
    log "Grunting finished"
}

docker-build(){
    build
    gruntfiles

    log "Preparing docker"
    mkdir $DIR/tmp
    cp $DIR/scripts/Dockerfile $DIR/tmp/
    cp $DIR/PubKeyManager $DIR/tmp/
    cp $DIR/pubkeymanager.conf $DIR/tmp/pubkeymanager.conf
    cp -r $DIR/public $DIR/tmp/
    rm -r $DIR/tmp/public/node_modules
    rm -r $DIR/tmp/public/assets/components
    rm -r $DIR/tmp/public/assets/css
    rm -r $DIR/tmp/public/assets/js
    rm -r $DIR/tmp/public/assets/tpls
    log "Dockerizing"
    cp $DIR/scripts/Dockerfile $DIR/tmp/Dockerfile
    cd $DIR/tmp
    docker build -t gerardsoleca/pubkeymanager:latest .
    cd $DIR
    rm -r $DIR/tmp
    log "Dockerizing finished"
}

docker-push(){
    log "Pushing docker image"
    docker push gerardsoleca/pubkeymanager:latest
    log "Push finished"
}

help() {
  echo -e -n "$BLUE"
  echo "-----------------------------------------------------------------------"
  echo "-                     Available commands                              -"
  echo "-----------------------------------------------------------------------"
  echo "   > setup          - Resolve node and bower dependencies"
  echo "   > build          - Build golang project"
  echo "   > gruntfiles     - Gruntify files"
  echo "   > docker-build   - Build docker image"
  echo "   > docker-push    - Push docker image"
  echo "-----------------------------------------------------------------------"
  echo -e -n "$NORMAL"
}

if [ ! -f scripts/task.sh ]; then
    error "Script must be run from project root-dir"
    exit -1
fi

if [ -z "$*" ]; then
    help
else
    $*
fi