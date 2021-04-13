# Switching integrations

## PluralKit (`pluralkit`)

The heart and soul of this entire project.

- `token`: The PluralKit token to access the API with.
  May also be given under the `Authorization` header.

## Mastodon (`mastodon`)

- `instance`: URL to your Mastodon instance, without trailing slash.
- `auth`: Bearer token, can be created at `{mastodon_instance}/settings/applications`.
  Needs `write:accounts` scope.
- `name`: Optional, the name to set the account to. Supports templating.
- `avatar`: Optional, the avatar to set the account to. Supports templating.

## Pushover (`pushover`)

- `application`: The Pushover Application token to be used.
  The specific instance being ran may have a fallback, but this is more reliable.
- `user`: The user or group ID to send the notification to.
- `priority`: Optional, the priority of the notification being sent. Defaults to `0`.
- `title`: Optional, the notification title to be sent. Supports templating.
- `message`: Optional, the message to be sent. Supports templating.
- `image_url`: Optional, the image to attach. Supports templating.
