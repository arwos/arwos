#!/bin/bash


if [ -f "/etc/systemd/system/arwos.service" ]; then
    systemctl start arwos
    systemctl enable arwos
    systemctl daemon-reload
fi

if [ -f "/etc/systemd/system/simple.service" ]; then
    systemctl start simple
    systemctl enable simple
    systemctl daemon-reload
fi
