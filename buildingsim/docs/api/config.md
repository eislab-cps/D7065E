# Config API

Server configuration endpoint.

## Get Config

```
GET /api/config
```

Returns server configuration flags.

```bash
curl http://localhost:9090/api/config
```

```json
{"edit_mode": false}
```

| Field | Type | Description |
|-------|------|-------------|
| `edit_mode` | bool | Whether floor plan editing tools are enabled (`--edit` flag) |
