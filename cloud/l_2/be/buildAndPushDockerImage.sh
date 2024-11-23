#!/bin/bash

TAG_VERSION="0.1.0-rc2"

docker build -t "sasha192bunin/uu_edu_monorepo_cloud_l2:$TAG_VERSION" -t sasha192bunin/uu_edu_monorepo_cloud_l2:latest . && \
docker push "sasha192bunin/uu_edu_monorepo_cloud_l2:$TAG_VERSION" && \
docker push "sasha192bunin/uu_edu_monorepo_cloud_l2:latest"
