# Contributing to ReleaseDesk

Hey there! 👋 Thanks for your interest in contributing to my project. It means a lot!
If you need any help or have questions, don’t hesitate to ask.

To understand the tech stack, please refer to the [Tech Stack](TECH_STACK.md) document.

## Quick Start

This project supports live reload with the [Go Air](https://github.com/air-verse/air) CLI. To get started:

1. Install the Air CLI.
2. Run it in the project root to enable hot reloading.
3. Open [http://localhost:8090](http://localhost:8090) in your browser.

Note: you must the UI Development command at least one time to generate the web assets.

### UI Development

- **Bundler**: Uses `esbuild` for fast build times. Note that it doesn’t validate TypeScript; that’s left to your IDE or
  the TypeScript CLI if needed.
- **CSS**: Tailwind CLI generates the CSS file, pulling classes from the web components and Go template files.

To run the UI development environment with live reload, use:

```bash
npm run dev
```

## Project Architecture

This project is a multi-page application with a hybrid setup. It uses Go’s `html/template` for server-side rendering,
sending
mostly plain HTML over the wire. Web Components step in only where extra interactivity is needed, keeping things simple
and sticking to web standards.

### Dependency Injection

Uses `uber-dig` for managing dependencies. All components (services, repositories, controllers) are added to a container
that resolves the dependency graph.

### Database

A SQLite database is used, with migrations handled by `golang-migrate`. Migrations are written as simple SQL files
stored in the `/migrations` directory. Each schema change requires a `change.up.sql` file (for applying the
change) and a `change.down.sql` file (for rollbacks).

### HTTP Server

Built using `go-chi` to manage both API and page requests.

### Code

This project follows Domain-Driven Design (DDD) concepts and draws inspiration from both the MVC and Clean Architecture
patterns and composability is highly recommended whenever possible.
The code style used is the default formatting provided by the IntelliJ IDE.

## References

- https://go-chi.io
- https://jmoiron.github.io/sqlx/
- https://lit.dev/docs/
- https://go.uber.org/dig
- https://github.com/golang-migrate/migrate