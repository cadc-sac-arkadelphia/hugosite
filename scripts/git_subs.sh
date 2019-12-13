#!/usr/bin/env bash

#git_subs.sh

# by David Skinner
# on October 1, 2019
# for CADC SAC Arkadelphia

# Copyright (c) 2019 Davsk Ltd Co. All Rights Reserved.
# MIT License.

git submodule init
git pull --recurse-submodules

echo "\\033[0;32mDeploying updates to GitHub...\\033[0m"
cd website/cadc-sac-arkadelphia || exit

# Build the project.
hugo -t cadc-sac

# Go To Public folder
cd public || exit
# Add changes to git.
git add -A

# Commit changes.
msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin master

# Come Back
cd ..
cd ../..