#!/bin/sh

cat <<EOT > docs.md
---
name: ntfy
description: plugin to send notifications to a ntfy.sh instance
author: l-x
tags: [trigger, notify]
containerImage: codeberg.org/l-x/woodpecker-ntfy
containerImageUrl: https://codeberg.org/l-x/-/packages/container/woodpecker-ntfy/latest
url: https://codeberg.org/l-x/woodpecker-ntfy
---

EOT

tail -n +7 README.md >> docs.md