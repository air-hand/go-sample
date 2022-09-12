# Go samples repository

# Requirements

- Windows 10
- Docker
- Docker compose
- VSCode
- VSCode Remote Container Extension

# How to develop?

This repository can be done with VSCode remote container extension, so you don't need to install Go lang on Host Machine.

1. Rebuild and open Container (if the slimmed container exists, rebuild is needed.)
1. make
1. /go/bin/app

## SSH-key sharing with container

Windows Open-ssh-agent (Named pipe) is required. (WSL ssh-agent won't work.)
