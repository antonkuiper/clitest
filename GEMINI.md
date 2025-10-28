# Project Overview

This project is a Go CLI program to make a database connection towards a postgres instance. so it is able to create database via a file .

# Building and Running

To run the program, use the following command:

```bash
go run main.go -host <host> -port <port> -user <user> -password <password> -dbname <dbname> -file db.sql
```

Replace the placeholders with your actual database credentials.

# Development Conventions

This project uses the standard Go conventions.

# Release Process

This project uses [semantic-release](https://github.com/semantic-release/semantic-release) to automate the release process. When a commit is pushed to the `main` branch, a new release is automatically created based on the commit messages.

To trigger a release, use the following commit message format:

```
<type>(<scope>): <subject>
```

For example:

```
feat: Add new feature
fix: Fix bug
chore: Update documentation
```