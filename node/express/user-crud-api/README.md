# user-crud-api

Lightweight User CRUD API using Node.js, Express, TypeScript and Prisma (Postgres).

## What this repository is
- Simple REST API exposing CRUD for a `User` model using Prisma as the ORM.
- API root: `GET /` — health check.
- Users endpoints mounted at `/api/users`.

## Key files
- [prisma.config.ts](prisma.config.ts) — Prisma configuration.
- [src/config/database.ts](src/config/database.ts) — Prisma client initialization and adapter detection.
- [src/server.ts](src/server.ts) — server startup and database connect logic.
- [src/app.ts](src/app.ts) — Express app and route mounting.
- [src/routes/user.routes.ts](src/routes/user.routes.ts) — users routes.
- [src/controllers/user.controller.ts](src/controllers/user.controller.ts) — controller logic for endpoints.
- [src/services/user.service.ts](src/services/user.service.ts) — business logic using Prisma client.
- [src/middleware/error.middleware.ts](src/middleware/error.middleware.ts) — error logging/response (temporarily verbose for debugging).
- [package.json](package.json) — scripts and dependencies.
- [.env](.env) — example env file containing `DATABASE_URL`.

## What I changed (summary of troubleshooting and fixes)
- Made `prisma.config.ts` omit `datasource.url` when `DATABASE_URL` is undefined to satisfy strict TS options.
- Adjusted `src/config/database.ts` to:
  - Detect available Prisma adapters at runtime (`@prisma/adapter-ppg` preferred, fallback to `@prisma/adapter-pg`).
  - Lazily construct the `PrismaClient` and export `adapterAvailable` so the server can start even when an adapter is not yet installed.
  - Provide a helpful stub that throws a clear error if DB operations are attempted without an adapter.
- Updated `src/server.ts` to skip `prisma.$connect()` / `.$disconnect()` when an adapter is missing (so server can run for local development without adapter installed).
- Added temporary debug logs in `src/controllers/user.controller.ts` and `src/services/user.service.ts` to print incoming payloads, created records, and Prisma errors.
- Enhanced `src/middleware/error.middleware.ts` to log stack traces and return error details in development (revert before production).
- Updated `package.json` to include adapter packages requested during debugging (`@prisma/adapter-pg` and `@prisma/adapter-ppg`).

## Why these changes
- Prisma v7 requires a driver adapter to connect to Postgres. During development the adapter was not installed (npm cache permissions blocked installs), causing runtime errors. The changes above allowed the server to start and made the failure mode explicit and recoverable while debugging.

## Environment / prerequisites
- Node.js (the project uses ESM).
- PostgreSQL reachable by `DATABASE_URL` in `.env`. Example in repository:

```
DATABASE_URL="postgresql://hareesh:root@localhost:5432/user_crud_db?schema=public"
PORT=5000
```

## Install & run
1. If you encounter npm cache ownership errors, fix permissions (replace UID:GID if different):

```bash
sudo chown -R 501:20 "/Users/hareesh/.npm"
```

2. Install dependencies:

```bash
cd /path/to/user-crud-api
npm install
```

3. Build (TypeScript):

```bash
npm run build
```

4. Start development server (auto-reloads):

```bash
npm run dev
```

## Quick API curl examples
- Create
```bash
curl -X POST http://localhost:5000/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Hareesh","email":"hareesh@example.com","age":30}'
```

- List
```bash
curl http://localhost:5000/api/users
```

- Get
```bash
curl http://localhost:5000/api/users/1
```

- Update
```bash
curl -X PUT http://localhost:5000/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated","email":"updated@example.com"}'
```

- Delete
```bash
curl -X DELETE http://localhost:5000/api/users/1
```

- Health
```bash
curl http://localhost:5000/
```

## Verification
- After installing the adapter and setting `DATABASE_URL`, `POST /api/users` should return `201` and the created user. `GET /api/users` should return persisted rows.

## Recommended cleanup before production
- Remove or reduce verbose error output in `src/middleware/error.middleware.ts` to avoid leaking internals.
- Remove temporary debug `console.log` statements in controllers/services.
- Use a proper migration workflow: `npx prisma migrate dev` (development) or `npx prisma migrate deploy` (production) and run `npx prisma generate` when updating schema.

## If you want me to continue
- I can revert debug changes, add integration tests for create/update/delete, or create a small script that runs the curl tests and reports results. Tell me which you prefer.

---
Generated on 2026-08-09 — concise troubleshooting and fix summary for this workspace.
