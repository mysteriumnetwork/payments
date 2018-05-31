#!/usr/bin/env bash
soljitsu combine --src-dir=./contracts --dest-dir=$1
rm -rf $1/deps.*