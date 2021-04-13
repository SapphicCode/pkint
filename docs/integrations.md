# Integrations

## PluralKit

The heart and soul of this entire project.

- `pluralkit_token`: The PluralKit token to access the API with. It'll shout at you otherwise.

## Mastodon

- `mastodon_instance`: URL to your Mastodon instance, without trailing slash.
- `mastodon_auth`: Bearer token, can be created at `{mastodon_instance}/settings/applications`. Needs `write:accounts` scope.
- `mastodon_name`: Optional, the name to set the account to. Supports templating.
- `mastodon_avatar`: Optional, the avatar to set the account to. Supports templating.

## Pushover

- `pushover_application`: The Pushover Application token to be used. The specific instance being ran may have a fallback, but this is more reliable.
- `pushover_user`: The user or group ID to send the notification to.
- `pushover_priority`: Optional, the priority of the notification being sent. Defaults to `0`.
- `pushover_title`: Optional, the notification title to be sent. Supports templating.
- `pushover_message`: Optional, the message to be sent. Supports templating.
- `pushover_image_url`: Optional, the image to attach. Supports templating.
