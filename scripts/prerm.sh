#!/bin/bash


if [ -f "/etc/systemd/system/agent.service" ]; then
    systemctl stop agent
    systemctl disable agent
    systemctl daemon-reload
fi

if [ -f "/etc/systemd/system/server.service" ]; then
    systemctl stop server
    systemctl disable server
    systemctl daemon-reload
fi
