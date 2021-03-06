#!/bin/bash

DEV_NAMESPACE="${DEV_NAMESPACE:-dev}"

#INSTALL_FROM_ALTINITY_RELEASE_DOCKERHUB="${INSTALL_FROM_ALTINITY_RELEASE_DOCKERHUB:-yes}"
INSTALL_FROM_ALTINITY_RELEASE_DOCKERHUB="${INSTALL_FROM_ALTINITY_RELEASE_DOCKERHUB:-no}"

INSTALL_FROM_PERSONAL_DEV_MANIFEST="${INSTALL_FROM_PERSONAL_DEV_MANIFEST:-yes}"
#INSTALL_FROM_PERSONAL_DEV_MANIFEST="${INSTALL_FROM_PERSONAL_DEV_MANIFEST:-no}"
PERSONAL_DEV_INSTALL_MANIFEST="${PERSONAL_DEV_INSTALL_MANIFEST:-./deployment-dockerhub-dev.yaml}"
PERSONAL_DEV_INSTALL_IMAGE="${PERSONAL_DEV_INSTALL_IMAGE:-sunsingerus/clickhouse-operator:dev}"

# In case both INSTALL_* options are "no" we are going to run operator manually, not from dockerhub
