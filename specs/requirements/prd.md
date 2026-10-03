 R-marker-1003c-presence.E2E marker 1003c-b5.

E2E marker p2-1003c.

# greeter — PRD

## Problem Statement

Teams building services on this platform need a minimal, known-good reference for a Go HTTP service — something small enough to stand up quickly, exercising the platform's conventions (per `app-factory-kaj/e2e-reference`) end to end without any real business complexity getting in the way.

## Solution

A small Go HTTP service, greeter, exposing a single `GET /hello` endpoint that returns a JSON greeting addressed to a name supplied as a query parameter, built to the conventions set by `app-factory-kaj/e2e-reference`.

## Actors

- **API Consumer** — any client (person or system) that calls the greeter service's HTTP endpoint to get a greeting.

## User Stories

1. As an API consumer, I want to send a GET request to `/hello` with a `name` query parameter, so that I receive a JSON greeting addressed to that name.
2. As an API consumer, I want to receive a sensible default greeting when I omit the `name` parameter, so that the endpoint still succeeds without it.

## Product Decisions

- Authentication: the `/hello` endpoint is publicly callable, with no API key, token or sign-in required. *assumed*
- Default behavior: a request with no `name` parameter returns a generic greeting (e.g. "Hello, World!") rather than an error. *assumed*
- Input handling: the service applies no extra validation or rate limiting on `name` — it is used as-is in the greeting, matching a small reference/demo service. *assumed*
- The service follows the structural and interface conventions demonstrated in `app-factory-kaj/e2e-reference`.

## Out of Scope

- User accounts, persistence, or any stored state.
- Rate limiting or abuse protection.
- Localization or multi-language greetings.
- Any endpoint other than `GET /hello`.

## Open Questions

None at this time.

## Further Notes

Persistence is out of scope for v1.