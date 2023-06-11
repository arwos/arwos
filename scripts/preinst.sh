#!/bin/bash


if ! [ -d /var/lib/agent/ ]; then
    mkdir /var/lib/agent
fi

if [ -f "/etc/systemd/system/agent.service" ]; then
    systemctl stop agent
    systemctl disable agent
    systemctl daemon-reload
fi

if ! [ -d /var/lib/server/ ]; then
    mkdir /var/lib/server
fi

if [ -f "/etc/systemd/system/server.service" ]; then
    systemctl stop server
    systemctl disable server
    systemctl daemon-reload
fi
