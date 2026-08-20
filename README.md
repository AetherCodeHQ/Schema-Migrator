# Schema Migrator

![CI](https://github.com/Qyroxen/Schema-Migrator/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Schema-Migrator/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Schema-Migrator?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Schema-Migrator)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Schema-Migrator)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Schema-Migrator?style=social)](https://github.com/Qyroxen/Schema-Migrator/stargazers)

## What is it?

Schema Migrator is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Schema-Migrator.git
cd Schema-Migrator
go build -o schemamigrator .

# Run
./schemamigrator --help
```

## CLI Usage

```bash
# Basic usage
./schemamigrator

# With flags
./schemamigrator --verbose --output json

# Get help
./schemamigrator --help
```

## Examples

```bash
# Example 1
./schemamigrator example1

# Example 2
./schemamigrator example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o schemamigrator .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Schema-Migrator/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Schema-Migrator?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Schema-Migrator/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Schema-Migrator?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Schema-Migrator/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Schema-Migrator" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Schema-Migrator/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Schema-Migrator" alt="Pull Requests">
  </a>
</p>
