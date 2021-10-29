#!/usr/bin/env bash

VERSION=$(date +%m%d%y)

docker build . -t golist:${VERSION} && \
	docker run --rm -i -t golist:${VERSION}
