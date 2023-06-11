#!/bin/bash


if [ -f "/etc/systemd/system/agent.service" ]; then
    systemctl start agent
    systemctl enable agent
    systemctl daemon-reload
fi

if [ -f "/etc/systemd/system/server.service" ]; then
    systemctl start server
    systemctl enable server
    systemctl daemon-reload
fi
