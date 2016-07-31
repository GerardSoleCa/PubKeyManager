FROM ubuntu:16.04
MAINTAINER GerardSoleCa <g.sole.ca@gmail.com>

RUN apt-get update && apt-get install -y libssl-dev

RUN mkdir -p /opt/pubkeymanager

COPY public /opt/pubkeymanager/public
COPY pubkeymanager.conf /opt/pubkeymanager/pubkeymanager.conf
COPY PubKeyManager /opt/pubkeymanager/PubKeyManager

WORKDIR /opt/pubkeymanager
CMD ./PubKeyManager