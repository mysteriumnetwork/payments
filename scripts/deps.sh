#!/bin/bash

depcmd=`which dep`

if [ -z "$depcmd" ]; then echo "You need to install dep - go dependency manager"; exit -1; fi;

$depcmd $@