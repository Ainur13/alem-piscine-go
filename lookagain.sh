#!/bin/bash
find -name '*.sh' | sed 's/.sh//g; s/.\///g' 