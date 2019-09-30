SHELL:=/usr/bin/env bash

BUILD:=$(shell date -u +%s)
PROJECT_SINCE:=$(shell git log --pretty=format:"%ad" --date=unix|tail -1)
AUTO_COUNT_SINCE:=$(shell echo $$((($(BUILD)-$(PROJECT_SINCE))/(24*60*60))))
AUTO_COUNT_YEAR:=$(shell echo $$(($(AUTO_COUNT_SINCE)/365)))
AUTO_COUNT_DAY:=$(shell echo $$(($(AUTO_COUNT_SINCE)%365)))
AUTO_COUNT_LOG:=$(shell git log --since=midnight --oneline|wc -l|tr -d " ")
CODEBASE_NUMBER:=0
SERIAL:=$(CODEBASE_NUMBER).$(AUTO_COUNT_YEAR).$(AUTO_COUNT_DAY)-$(AUTO_COUNT_LOG)
TAG:=$(shell git describe --tags || echo NO-TAG)
HASH:=$(shell git describe --always --dirty)
BRANCH:=$(shell git symbolic-ref --short HEAD)
LDFLAGS:=-ldflags=" \
-X \"main.branch=$(BRANCH)\" \
-X \"main.build=$(BUILD)\" \
-X \"main.tag=$(TAG)\" \
-X main.hash=$(HASH) \
-X main.serial=$(SERIAL) \
"

COMMIT:=4b825dc
REVIEWDOG:=| reviewdog -efm='%f:%l:%c: %m' -diff="git diff $(COMMIT) HEAD"

GO:=go
GOM:=GO111MODULE=on $(GO)
GOPATH:=$(shell $(GO) env GOPATH)
PKG:=$(shell $(GO) list)
GOBIN:=$(notdir $(PKG))
GOLIST:=$(shell $(GO) list ./...)
GODIR:=$(patsubst $(PKG)/%,%,$(wordlist 2,$(words $(GOLIST)),$(GOLIST)))

PROTO:=$(shell find . -type d -name ".?*" -prune -or -type d -name vendor -prune -or -type f -name "*.proto" -print)
PB_GO:=$(PROTO:.proto=.pb.go)
VALIDATOR_PB_GO:=$(PROTO:.proto=.validator.pb.go)
IF_GO:=$(shell find . -type d -name vendor -prune\
 -or -type f -name "if-*.go" -print\
 -or -type f -name "entity-*.go" -print\
 -or -type f -name "vo-*.go" -print\
)

GOFILE:=$(filter-out %.pb.go $(IF_GO),$(shell find . -type d -name vendor -prune -or -type f -name "*.go" -print))
GOSRC:=$(GOFILE) $(PB_GO)
CEL:=$(shell find . -type d -name vendor -prune -or -type f -name "*.cel.txt" -print)
GOCEL:=$(patsubst %.cel.txt,%_gen.go,$(CEL))
#GOGEN:=$(shell find . -type d -name vendor -prune -or -type f -name "*.go" -print|xargs grep "^//go:generate " -l)

APP_DIR_PATH:=app
GOPHERJS:=$(APP_DIR_PATH)/web/gopher.js

LIBGO:=$(wildcard lib/*.go)
GOLIB:=$(LIBGO:.go=.so)
.SUFFIXES: .go .so
.go.so:
	$(GO) build -buildmode=c-shared -o $@ $<
%.pb.go: %.proto
	prototool all $<
%.validator.pb.go: %.proto
	( type protoc > /dev/null 2>&1 ) && protoc --govalidators_out=$(dir $<) -I $(dir $<) -I vendor $<


.PHONY: all
all: $(GOBIN) $(GOLIB) $(APP_DIR_PATH)/build
.PHONY: uml
uml:
	for i in domain application service usecase; do\
  [ -d $$i ] && gouml init --file $$i --out $$i/$$i.puml || echo no exists $$i;\
  done
$(VALIDATOR_PB_GO):
.PHONY: proto
proto: vendor $(PB_GO) $(VALIDATOR_PB_GO)
.PHONY: golang
golang: $(GOBIN) $(GOLIB)
.PHONY: gopherjs
gopherjs: $(GOPHERJS)
.PHONY: dart
dart: $(APP_DIR_PATH)/build


# do not use bundler in Docker container
ifeq ($(shell test -f /proc/self/cgroup && cat /proc/self/cgroup|grep -q docker;echo $$?),0)
BUNDLE_EXEC:=
else
BUNDLE_EXEC:=bundle exec
BUNDLE:=.bundle/bundle
Gemfile.lock: Gemfile
	bundle install --path $(BUNDLE)
$(BUNDLE): Gemfile.lock
	bundle install --path $(BUNDLE)
$(HTML): $(BUNDLE)
endif


deps: go.mod
	$(GOM) mod download
vendor: go.mod
	$(GOM) mod vendor
	$(GOM) mod tidy
	-git checkout -f vendor
.PHONY: generate
generate: vendor
	-go generate ./...
	make clean

$(IF_GO): $(filter-out $(IF_GO),$(GOSRC))
	-go generate ./...
%_gen.go: %.go %.cel.txt
	-go generate ./...

.PHONY: go-get
go-get: $(GOSRC)
	echo > go.mod
	rm -rf vendor
	time \
 $(GOM) build $(LDFLAGS)

$(GOBIN): deps $(GOSRC) $(GOCEL)
	time \
 $(GOM) build $(LDFLAGS)" -X \"main.semver=$(SERIAL)+$(HASH)\""

GOX_OPTION:=-osarch="darwin/amd64 linux/amd64 linux/arm"

.PHONY: channel
channel: deps $(GOSRC)
	time \
 GO111MODULE=on gox $(GOX_OPTION) \
 -output="assets/gox/$(BRANCH)/$(SERIAL)+$(HASH)/{{.OS}}-{{.Arch}}" \
 $(LDFLAGS)" -X \"main.semver=$(SERIAL)+$(HASH)\" -X \"main.channel=channel/$(BRANCH)\""
	mkdir -p docs/channel/$(BRANCH)/$(GOBIN)
	./assets/script/selfupdate.sh docs/channel/$(BRANCH)/$(GOBIN) 3
	time go-selfupdate -o docs/channel/$(BRANCH)/$(GOBIN) assets/gox/$(BRANCH)/$(SERIAL)+$(HASH) $(SERIAL)+$(HASH)

.PHONY: release
release: deps $(GOSRC) $(GOCEL)
	time \
 GO111MODULE=on gox $(GOX_OPTION) \
 -output="assets/gox/$(TAG)/{{.OS}}-{{.Arch}}" \
 $(LDFLAGS)" -X \"main.semver=$(TAG)\" -X \"main.channel=release\""
	mkdir -p docs/release/$(GOBIN)
	./assets/script/selfupdate.sh docs/release/$(GOBIN) 3
	time go-selfupdate -o docs/release/$(GOBIN) assets/gox/$(TAG) $(TAG)

.PHONY: package
package:
	./assets/script/package.sh

$(GOPHERJS): $(GOSRC) $(GOCEL)
	@# https://github.com/gopherjs/gopherjs/issues/598#issuecomment-282563634
	-find $(GOPATH)/pkg -depth 1 -type d -name "*_js" -exec rm -fr {} \;
	-find $(GOPATH)/pkg -depth 1 -type d -name "*_js_min" -exec rm -fr {} \;
	gopherjs build --tags gopherjs --minify $(PKG)/hackernews/gopherjs --output $@


NODE_MODULES_BASE:=node_modules
UGLIFYJS:=$(NODE_MODULES_BASE)/.bin/uglifyjs

$(NODE_MODULES_BASE): package.json
	npm install

.PHONY: uglifyjs
uglifyjs: $(GOPHERJS) $(NODE_MODULES_BASE)
	cd $(dir $<) && ../../$(UGLIFYJS) --compress --mangle --output $(notdir $<) $(notdir $<)


SLIM:=$(foreach dir,lib web,$(shell find $(APP_DIR_PATH)/$(dir) -type f -name "*.slim" -print))
HTML:=$(SLIM:.slim=.html)
.SUFFIXES: .slim .html
.slim.html:
	$(BUNDLE_EXEC) slimrb --pretty --option splat_prefix="'**'" --option code_attr_delims="{'{' => '}'}" --option attr_list_delims="{'{' => '}'}" $< > $@
html: $(HTML)


PUB_SPEC:=$(shell find . -type d -name build -prune -o -type f -name pubspec.yaml -print)
PUB_LOCK:=$(PUB_SPEC:.yaml=.lock)
.SUFFIXES: .yaml .lock
.yaml.lock:
	(cd $(dir $@) && pub get)

$(APP_DIR_PATH)/.packages: $(PUB_LOCK)
	(cd $(dir $@) && pub get)

DART:=hackernews/dart/lib/hackernews.dart
G_DART:=$(DART:.dart=.g.dart)
.SUFFIXES: .dart .g.dart
.dart.g.dart:
	(cd hackernews/dart && pub run build_runner build --delete-conflicting-outputs)

.PHONY: $(APP_DIR_PATH)/build
$(APP_DIR_PATH)/build: uglifyjs $(PUB_LOCK) $(APP_DIR_PATH)/.packages $(G_DART) $(HTML)
	(cd $(APP_DIR_PATH)\
	&& pub run build_runner build --delete-conflicting-outputs --release --fail-on-severe --output build\
	&& pub run pwa --exclude ".DS_Store,packages/**,.packages,*.dart,*.js.deps,*.js.info.json,*.js.map,*.js.tar.gz,*.module"\
	)
	rm -rf $(APP_DIR_PATH)/build/web/packages/test
	find $(APP_DIR_PATH)/build/web -type f -name "*.dart" -exec rm -f {} \;
	find $(APP_DIR_PATH)/build/web -type f -name "*.slim" -exec rm -f {} \;
	while [ `find $(APP_DIR_PATH)/build/web -type d -empty |wc -l` != 0 ] ; do find $(APP_DIR_PATH)/build/web -type d -empty -exec rm -rf {} + ; done

UNAME:=$(shell uname -s)
ifeq ($(UNAME),Darwin)
CONVERT_PREFIX:=docker-compose run
endif
CONVERT:=$(CONVERT_PREFIX) convert -verbose -colorspace sRGB -transparent white -density 256x256
$(APP_DIR_PATH)/web/images/icons/android-icon-%.png \
$(APP_DIR_PATH)/web/images/icons/favicon-%.png \
: $(APP_DIR_PATH)/web/images/logo.svg
	$(CONVERT) -filter Point $< -trim +filter -resize $* $@

$(APP_DIR_PATH)/web/favicon.ico: $(APP_DIR_PATH)/web/images/logo.svg
	$(CONVERT) -filter Point $< -trim +filter -resize 16x -define icon:auto-resize -colors 256 $@


RELEASE:=--release
.PHONY: serve
serve: $(GOPHERJS) $(PUB_LOCK) $(G_DART) $(HTML)
	(cd $(APP_DIR_PATH) && webdev serve $(RELEASE))

.PHONY: fixes-webdev
fixes-webdev:
	-# https://github.com/dart-lang/sdk/issues/33601#issuecomment-402469804
	( type pubglobalupdate > /dev/null 2>&1 ) || pub global activate pubglobalupdate ; pubglobalupdate

.PHONY: deploy
deploy: $(APP_DIR_PATH)/build
	firebase deploy

.PHONY: clean
clean:
	find . -type f -name coverprofile -delete
	rm -f $(GOBIN) $(GOLIB) $(wildcard lib/*.h)
	rm -rf package $(APP_DIR_PATH)/build
	find . -name .DS_Store -delete
	find assets -type d -name assets -delete
	for file in $$(find . -type d -name vendor -prune\
 -or -type f -name "entity-*.go" -print\
 -or -type f -name "gen-*.go" -print\
 -or -type f -name "if-*.go" -print\
 -or -type f -name "vo-*.go" -print\
); do\
  sed -i '' 's|"github.com/michilu/boilerplate/vendor/github.com/|"github.com/|g' $$file;\
  chmod -x $$file;\
  done

.PHONY: test
test: deps
	$(GOM) test $(PKG)/...

.PHONY: pprof
pprof:
	pprof -http=: localhost:8888/debug/pprof/profile

.PHONY: bench
bench:
	$(GOM) test -bench . -benchmem -count 5 -run none $(PKG)/... | tee bench/now.txt
	[ -f bench/before.txt ] && ( type benchcmp > /dev/null 2>&1 ) && benchcmp bench/before.txt bench/now.txt || :

.PHONY: coverage
coverage:
	@for pkg in $(GOLIST); do\
		echo start test for $$pkg;\
		$(GOM) test $$pkg -race -coverprofile=$${pkg#$(PKG)/}/coverprofile -covermode=atomic;\
	done
	rm -f coverage.txt && find . -type f -name coverprofile -exec tail -n+2 {} >>coverage.txt \; -delete
	$(GOM) tool cover -html=coverage.txt -o coverage.html

.PHONY: lint
lint:
	-gofmt -s -w .
	@echo
	-echo $(GOLIST) | xargs -L1 golint
	@echo
	-deadcode $(GODIR) 2>&1
	@echo
	-find $(GODIR) -type f -exec misspell {} \; $(REVIEWDOG)
	@echo
	-staticcheck $(GOLIST) $(REVIEWDOG)
	@echo
	-errcheck $(GOLIST) $(REVIEWDOG)
	@echo
	-safesql $(GOLIST)
	@echo
	-goconst $(GOLIST) $(REVIEWDOG)
	@echo
	-$(GOM) vet $(GOLIST) $(REVIEWDOG)
	@echo
	-$(GOM) vet -shadow $(GOLIST) $(REVIEWDOG)
	@echo
	-aligncheck $(GOLIST) $(REVIEWDOG)
	@echo
	-gosimple $(GOLIST) $(REVIEWDOG)
	@echo
	-unconvert $(GOLIST) $(REVIEWDOG)
	@echo
	-interfacer $(GOLIST) $(REVIEWDOG)

.PHONY: review
review:
	$(MAKE) lint COMMIT:=master

.PHONY: review-dupl
review-dupl:
	-git diff $(COMMIT) HEAD --name-only --diff-filter=AM|grep -e "\.go$$" | xargs dupl
