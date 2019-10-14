#!/bin/bash
echo $(curl -s https://api.github.com/users/Ainur13 | jq '.id')
