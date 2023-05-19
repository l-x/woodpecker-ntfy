---
name: ntfy
description: plugin to send notifications to a ntfy.sh instance
authors: l-x
tags: [trigger, notify]
containerImage: codeberg.org/l-x/woodpecker-ntfy
containerImageUrl: https://codeberg.org/l-x/-/packages/container/woodpecker-ntfy
url: https://codeberg.org/l-x/woodpecker-ntfy
---

# woodpecker-ntfy

A [Woodpecker] plugin to send notifications to a [ntfy.sh] instance.

## Configuration

| Name       | Description                                                            | Default                           |
| ---------- | ---------------------------------------------------------------------- | --------------------------------- |
| `url`      | Url (including the topic) to send the notification to                  | `https://ntfy.sh/woodpecker-ntfy` |
| `token`    | Authentication token for write-protected topics [^bearer-auth]         | none                              |
| `title`    | Notification Title [^message-title]                                    | none                              |
| `priority` | Notification Priority [^message-priority]                              | none                              |
| `actions`  | Action Buttons [^defining-actions]                                     | none                              |
| `click`    | Click Action [^click-action]                                           | `CI_BUILD_LINK`                   |
| `icon`     | Message Icon [^icons]                                                  | `CI_COMMIT_AUTHOR_AVATAR`         |
| `tags`     | Tags and Emojis [^tags-emojis]                                         | none                              |
| `message`  | Notification Body                                                      | none                              |
| `email`    | E-mail to which the message is to be forwarded [^e-mail-notifications] | none                              |

## Example

```yaml
pipeline:
    ntfy:
        image: codeberg.org/l-x/woodpecker-ntfy
        settings:
            url: https://custom.ntfy.instance/topic-to-notify
            token:
                from_secret: your-super-secret-ntfy-access-token
            title: notification title
            priority: urgent
            actions: "view, Open portal, https://home.nest.com/, clear=true; http, Turn down, https://api.nest.com/, body='{\"temperature\": 65}'"
            click: https://where.to.go
            icon: https://woodpecker-ci.org/img/logo.svg
            tags: robot,${CI_BUILD_EVENT},${CI_REPO_NAME}
            message: >
                📝 Commit by ${CI_COMMIT_AUTHOR} on ${CI_COMMIT_BRANCH}:

                ${CI_COMMIT_MESSAGE}
```

[Woodpecker]: https://woodpecker-ci.org/
[ntfy.sh]: http://ntfy.sh/

[^bearer-auth]: https://docs.ntfy.sh/publish/#bearer-auth
[^message-title]: https://docs.ntfy.sh/publish/#message-title
[^message-priority]: https://docs.ntfy.sh/publish/#message-priority
[^defining-actions]: https://docs.ntfy.sh/publish/#defining-actions
[^click-action]: https://docs.ntfy.sh/publish/#click-action
[^icons]: https://docs.ntfy.sh/publish/#icons
[^tags-emojis]: https://docs.ntfy.sh/publish/#tags-emojis
