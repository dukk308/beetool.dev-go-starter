# API Structure

## Prefixes

| Prefix     | Audience           | Exposed | Use                                       |
| ---------- | ------------------ | ------- | ----------------------------------------- |
| `/private` | Internal only      | No      | Internal calls; do not expose to clients. |
| `/admin`   | Admin or Editor    | Yes     | Admin/editor-only API routes.             |
| `/public`  | Anyone             | Yes     | Public API; no auth required.             |
| `/`        | Authenticated user | Yes     | User API; requires auth (user context).   |

## Path pattern

Full path format:

```
prefix/{rest}
```

- **prefix**: One of `/private`, `/admin`, `/public`, or `/` (root).
- **{rest}**: Resource path and optional IDs (e.g. `v1/notes`, `v1/account/profile`).

## Examples

| Prefix     | Example path          | Description                       |
| ---------- | --------------------- | --------------------------------- |
| `/private` | `/private/health`     | Internal health; not exposed.     |
| `/admin`   | `/admin/v1/users`     | List/manage users (admin/editor). |
| `/public`  | `/public/v1/blogs`    | Public blogs; no auth.            |
| `/`        | `/v1/notes`           | User notes (authenticated).       |
| `/`        | `/v1/account/profile` | User profile (authenticated).     |

## Summary

- **Private**: internal only, never exposed.
- **Admin**: admin/editor routes; protect by role.
- **Public**: no auth; safe to expose.
- **Root (`/`)**: user-facing API; require auth.
- All routes follow `prefix/{rest}`.
