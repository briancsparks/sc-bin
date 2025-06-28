#!/bin/bash

# Script to update commit messages from "WIP" to descriptive messages
# This uses git filter-branch to rewrite commit history

# Update multiple commit messages in one go
# The --msg-filter option processes commit messages through a shell script
git filter-branch -f --msg-filter '
if [ "$GIT_COMMIT" = "49c40c1bf3d2a6a7afed9cc9bdcaee65d6d58e12" ]; then
    echo "docs: Add initial structure and content to how-to guides"
elif [ "$GIT_COMMIT" = "30faf2f0137013348e825dee031af6eee55b386a" ]; then
    echo "docs: Create comprehensive how-to guides for Zig-C++ integration"
elif [ "$GIT_COMMIT" = "796444c68ebd8c371a9f3ca10d622bd163f36a53" ]; then
    echo "docs: Move adding-new-types guide to docs/howto directory"
elif [ "$GIT_COMMIT" = "a26141109c76b10be4fde924ada408023cb17b97" ]; then
    echo "docs: Create how-to guide for adding new types to stack math"
else
    cat  # Keep original message for all other commits
fi
' HEAD~4..HEAD

# Notes:
# - The -f flag forces filter-branch to run even if there's a backup
# - HEAD~4..HEAD means "rewrite the last 4 commits"
# - $GIT_COMMIT contains the SHA-1 of each commit being processed
# - 'cat' passes through unchanged messages for commits we don't want to modify

# Alternative approaches that could also work:

# 1. Using git rebase with exec (requires manual editing in interactive mode):
# git rebase -i HEAD~4
# Then change each "pick" to "reword" and save

# 2. Using git commit --amend for the most recent commit only:
# git commit --amend -m "New commit message"

# 3. Using git filter-repo (newer, recommended replacement for filter-branch):
# git filter-repo --message-callback '
#   if commit.original_id == b"49c40c1bf3d2a6a7afed9cc9bdcaee65d6d58e12":
#     return b"docs: Add initial structure and content to how-to guides"
#   return message
# '