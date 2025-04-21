#!/bin/bash


if [ -f "/etc/systemd/system/arwos.service" ]; then
    systemctl stop arwos
    systemctl disable arwos
    systemctl daemon-reload
fi

if [ -f "/etc/systemd/system/simple.service" ]; then
    systemctl stop simple
    systemctl disable simple
    systemctl daemon-reload
fi
