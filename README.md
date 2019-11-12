# boilerplate

# Install
Use `go get -u github.com/michilu/boilerplate`.

OR

```console
$ git clone --depth=1 https://github.com/michilu/boilerplate.git
$ cd boilerplate
$ make
```

# develop

## dependencies

### protobuf
- [Protocol Buffers Version 3 Language Specification  \|  Protocol Buffers  \|  Google Developers](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec)
  - [Package google\.protobuf  \|  Protocol Buffers  \|  Google Developers](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)
- [mwitkow/go-proto-validators: Generate message validators from .proto annotations.](https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto)
- [protocolbuffers/protobuf: Protocol Buffers \- Google's data interchange format](https://github.com/protocolbuffers/protobuf) `$ brew install protobuf`
- [uber/prototool: Your Swiss Army Knife for Protocol Buffers](https://github.com/uber/prototool) `$ brew install prototool`

- [mitchellh/gox: A dead simple, no frills Go cross compile tool](https://github.com/mitchellh/gox)
- [rjeczalik/interfaces: Code generation tools for Go.](https://github.com/rjeczalik/interfaces)
- [sanbornm/go-selfupdate: Enable your Go applications to self update](https://github.com/sanbornm/go-selfupdate)

### go generate

require go1.12.x

```console
$ brew install go@1.12
$ echo "PATH_add /usr/local/opt/go@1.12/bin" >> .envrc
$ direnv allow
```

```console
$ go get -u github.com/cheekybits/genny
$ GO111MODULE=on go get -u github.com/rjeczalik/interfaces/cmd/interfacer
```

## build

- `make`

- `make golang`
- `make proto`
- `make generate`
- `make gox`:       cross compile
- `make channel`:   push out self update files to develop channel
- `make release`:   push out self update files to release channel
- `make uml`:       generate PlantUML from Go codes

- `make gopherjs`
- `make serve`

- `make clean`:     clean up

## launch

```console
$ ./assets/daemon/loop ./boilerplate --debug --verbose --trace --profiler --update run
```

## Architecture

### DDD

Each directory contains:
- `domain`: Domain Service, Domain Rule, Entity, Value Object, Repository Interface
  - Layered architecture: Domain
  - Hexagonal architecture: Domain Model
  - Onion architecture: Domain Model, Domain Service
  - Clean architecture: Entities
- `application`: Application, Handler of Entity/Value Object, Dataflow
  - Layered architecture: Application
  - Hexagonal architecture: Application
  - Onion architecture: Application Service
  - Clean architecture: Use Cases, Controllers
- `infra`: Repository
  - Layered architecture: Infrastructure
  - Hexagonal architecture: Adapter
  - Onion architecture: Infrastructure
  - Clean architecture: Gateways
- `presentation`: Endpoint
  - Layered architecture: User Interface, Presentation
  - Hexagonal architecture: Adapter
  - Onion architecture: User Interface
  - Clean architecture: Presentaters
- `usecase`: Use Case, Requirements
  - The context in Agile software development, not DDD.

### Event-driven architecture
NOT the Event Sourcing System.

# deploy

```console
$ make deploy
```

# environment variables for Wercker

see: https://app.wercker.com/<organization\>/<application\>/environment

- `FIREBASE_PROJECT`: needs by the deploy step. see: https://console.firebase.google.com/
- `FIREBASE_TOKEN`: needs by the deploy step, via `$(firebase login:ci)`. see: https://github.com/firebase/firebase-tools#using-with-ci-systems
- `NETLIFY_BRANCH_DEPLOY_SITE_ID`: `API ID(UUID4)` on https://app.netlify.com/sites/<site-name\>/settings/general
- `NETLIFY_TOKEN`: https://app.netlify.com/account/applications/personal
- `SLACK_TOKEN`: needs by the ['slackcli'](https://github.com/cixtor/slackcli) command. see: https://api.slack.com/custom-integrations/legacy-tokens
- `SLACK_URL`: needs by the 'slack-notifier' step. see: https://slack.com/apps/A0F7XDUAZ-incoming-webhooks

# setup the workflows on Wercker

see: https://app.wercker.com/<organization>/<application>/workflows

![](assets/wercker-pipeline.png)
![](assets/wercker-workflow.png)

# tools
- [mitchellh/gox: A dead simple, no frills Go cross compile tool](https://github.com/mitchellh/gox)
- [sanbornm/go-selfupdate: Enable your Go applications to self update](https://github.com/sanbornm/go-selfupdate)
- [PlantUML - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml)
  - [kazukousen/gouml: Automatically generate PlantUML from Go Code\.](https://github.com/kazukousen/gouml) `go get -u github.com/kazukousen/gouml/cmd/gouml`
- [XML Pretty Print](https://jsonformatter.org/xml-pretty-print)

# optimize

- https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas<Paste>
  - https://github.com/cathalgarvey/fmtless

1. minify: https://github.com/gopherjs/gopherjs#performance-tips

:

    7,196,874 bytes:  100%:   $ gopherjs build
    4,661,791 bytes:   65%:   $ gopherjs build --minify

2. UglifyJS 3: https://github.com/mishoo/UglifyJS2

:

    4,547,810 bytes:   63%:   $ gopherjs build --minify && uglifyjs
    4,274,152 bytes:   59%:   $ gopherjs build --minify && uglifyjs --compress
    3,843,890 bytes:   53%:   $ gopherjs build --minify && uglifyjs --compress --mangle

# sync subtree

Add upstream:
```console
$ git clone git@github.com:michilu/boilerplate.git .
$ git remote add upstream https://github.com/dart-lang/angular.git
$ git checkout upstream/master
$ git subtree split --prefix=examples/hacker_news_pwa -b examples/hacker_news_pwa
```

Sync to upstream:
```console
$ git fetch upstream master
$ git checkout upstream/master
$ git subtree push --prefix=examples/hacker_news_pwa origin examples/hacker_news_pwa
$ git checkout upstream
$ git subtree pull --prefix=app origin examples/hacker_news_pwa
```

ref:
- https://github.com/dart-lang/angular/tree/master/examples/hacker_news_pwa

## for proto

Add upstream:
```console
$ git remote add googleapis https://github.com/googleapis/googleapis
$ git fetch --depth=1 --no-tags googleapis
$ git checkout googleapis/master
$ git subtree split --prefix=google/type -b google/type
```

Add subtree:
```console
$ git checkout dev
$ git subtree add --prefix=vendor/github.com/googleapis/googleapis/google/type google/type
```

Sync to upstream:
```console
$ git fetch --no-tags googleapis
(TBD)
$ git subtree pull --prefix=vendor/github.com/googleapis/googleapis/google/type origin google/type
```
