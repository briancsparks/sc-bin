#!/bin/bash -e

# See: https://claude.ai/chat/e8ee936f-92ba-468b-a83e-b147e009725e
#
# When changing this file:
# systemctl --user restart clipboard-monitor

# Use in a script
while clipnotify; do
    # Clipboard changed!
    content=$(xclip -selection clipboard -o)
    echo "New clipboard content: $content"

    # Do something with it
    notify-send "Clipboard" "$content"
done
