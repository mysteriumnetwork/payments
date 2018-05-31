#!/usr/bin/env bash
soljitsu combine --src-dir=$1 --dest-dir=$2
rm -rf $2/deps.*