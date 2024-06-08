# AI Language Tutor

## Roadmap

- [x] Authorization
- [x] DB integration
- [x] Storage integration
- [x] Websocket room's communication
- [ ] Call mode
    - [x] Room creation
    - [x] Basic call mode UI
    - [x] Voice recording & saving to bucket
    - [ ] Voice to text transformation
    - [ ] Text request to GPT model
    - [ ] Text to voice transformation
    - [ ] Room history extraction
    - [ ] Context configurator
    - [ ] Call mode extended UI
- [ ] Chat mode
    - [ ] Basic chat mode UI
    - [ ] Text request to GPT model
- [ ] Text to speech converter


## Build & Development

> **Important**: Please keep in mind the project is built to have only one executable & transferable file. All the assets/credentials are embeded into the binary.

> **Important**: The build can be run in two modes: `release` and `devonly`. `devonly` build mode does not serve embeded UI assets to support quick development and to not re-compile backend side each time. Please use `make run` for compiling project in `devonly` mode, and `make build` to get a production-ready binary.

### Requirements

1. Google OAuth
1. Google Cloud Storage (service key)
1. MySQL/MariaDB 5.7 or higher
1. NodeJS v18

### Init Project

1. Run `make init` and add your GCP service key (with Read-Write-Create access to Cloud Storage) to the `./credentials/gcp.json` file
1. Ensure you set all the variables in `./credentials/env` file. Here's the example:

```env
APP_KEY=A02E..F== # base64 encoded bytes, used for JWT generating
GOOGLE_OAUTH_CLIENT_ID={cient id}
GOOGLE_OAUTH_SECRET={oauth secret}
GOOGLE_OAUTH_CALLBACK_URL=http://localhost:3000/auth/callback/google
STORAGE_BUCKET_NAME=my-bucket-name
DB_DSN=luke:secret@tcp(127.0.0.1:3306)/tutor
```

### Development

1. Use `make ui-watch`, or run `npm run watch` from the `./ui` dir
1. Use `make run` to compile project in `devonly` mode. Please keep in mind server will be serving assets from the actual `./ui/public` directory
