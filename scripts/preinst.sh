#!/bin/bash


if ! [ -d /var/lib/arwos/ ]; then
    mkdir /var/lib/arwos
fi

if [ -f "/etc/systemd/system/arwos.service" ]; then
    systemctl stop arwos
    systemctl disable arwos
    systemctl daemon-reload
fi

if ! [ -d /var/lib/simple/ ]; then
    mkdir /var/lib/simple
fi

if [ -f "/etc/systemd/system/simple.service" ]; then
    systemctl stop simple
    systemctl disable simple
    systemctl daemon-reload
fi
