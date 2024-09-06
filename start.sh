#!/bin/bash
function start_user() {
    go run user/user.go -f user/etc/user.yaml
}

function start_gateway() {
    o run gateway/gateway.go -f gateway/etc/gateway.yaml
}

(start_user) &
(start_gateway)