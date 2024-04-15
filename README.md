# Envex

Envex (Env Example) is a command-line tool written in Go for generating sample dotenv files from existing `.env` files. It helps streamline the process of creating sample environment configuration files for your projects.

## Usage

```bash
go install github.com/wickes1/envex

envex generate
# Sample dotenv file has been saved to: .env.example

# Specify input and output files
envex generate -f .env -o .env.sample

# Retain comments
envex generate -c
```

## Example

Given the following `.env` file:

```env
# Example dotenv

# Database

DB_HOST=localhost
DB_USER=root
DB_PASS=secret
DB_NAME=database

# App

APP_NAME=MyApp
APP_ENV=development
APP_DEBUG=true
APP_URL=http://localhost:8000
```

Running `envex generate` will generate the following `.env.example` file:

```env
DB_HOST=
DB_USER=
DB_PASS=
DB_NAME=
APP_NAME=
APP_ENV=
APP_DEBUG=
APP_URL=
```

Running `envex generate -c` will retain the comments:

```env
# Example dotenv

# Database

DB_HOST=
DB_USER=
DB_PASS=
DB_NAME=

# App

APP_NAME=
APP_ENV=
APP_DEBUG=
APP_URL=
```
