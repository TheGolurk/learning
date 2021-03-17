#!/bin/bash

if [ -z "$1" ]; then
	echo usage: $0 directory
fi

SRCD=$1
TGTD="/var/backups"
OF=home-$(date +%Y%m%d).tgz
tar -cZf $TGTD$OF $SRCD
