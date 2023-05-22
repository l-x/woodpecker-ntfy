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
| `message`  | Notification Body                                                      | none                              |
| `title`    | Notification Title [^message-title]                                    | none                              |
| `priority` | Notification Priority [^message-priority]                              | none                              |
| `url`      | Url (including the topic) to send the notification to                  | `https://ntfy.sh/woodpecker-ntfy` |
| `token`    | Authentication token for write-protected topics [^bearer-auth]         | none                              |
| `actions`  | Action Buttons [^defining-actions]                                     | none                              |
| `attach`   | Url for file to be attached [^attach-file-from-a-url]                  | none                              |
| `call`     | Phone number to send voice message to [^phone-calls] (ntfy >= 2.5.0)   | none                              |
| `click`    | Click Action [^click-action]                                           | `CI_BUILD_LINK`                   |
| `email`    | E-mail to which the message is to be forwarded [^e-mail-notifications] | none                              |
| `icon`     | Message Icon [^icons]                                                  | `CI_COMMIT_AUTHOR_AVATAR`         |
| `tags`     | Tags and Emojis [^tags-emojis]                                         | none                              |

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
[^e-mail-notifications]: https://docs.ntfy.sh/publish/#e-mail-notifications
[^attach-file-from-a-url]: https://docs.ntfy.sh/publish/#attach-file-from-a-url
[^phone-calls]: https://docs.ntfy.sh/publish/#phone-calls
