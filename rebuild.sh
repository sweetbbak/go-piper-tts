#!/bin/bash

[ -f piper.sh ] && rm piper.sh
size="$(wc -c squashfs-start.sh | cut -d ' ' -f1)"
file=./squashfs-start.sh
sed -i "s/scriptsize=.*/scriptsize=${size}/g" "${file}"