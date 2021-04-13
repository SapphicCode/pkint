# API surface

## Authentication

Authentication is achieved via the `Authorization` header.

## Routes

### `GET /`

Will redirect you to [pk-webs](https://spectralitree.github.io/pk-webs/dash), our PluralKit interface of choice.

### `GET /members`

Returns an easily-parsed list of members. [This endpoint requires authorization.](integrations.md)

Query parameters, with defaults:

- `sort`: `alpha` (default), `front_history`: The sorting method to use.
- `template`: `{{ Member.Name }} ({{ Member.ID }})`: Generates output like `Cassandra (bulpn)`.

### `POST /switch`

Query parameters:

- `regex`: `\((\w{5})\)$`: Regex to match the PluralKit ID. Default matches the default return of `GET /members`.

[Body parameters](integrations.md)
