#!/bin/bash

# Get a list of modified files comparing current branch to 'generated'
modified_files=$(git diff --name-only generated)

for file in $modified_files; do
    # Generate a patch file for each file
    git diff generated -- "$file" > "custom/${file//\//_}.patch"
done
