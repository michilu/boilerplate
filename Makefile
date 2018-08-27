PROJECT_SINCE:=$(shell git log --pretty=format:"%ad" --date=unix|tail -1)
AUTO_COUNT_SINCE:=$(shell echo $$(((`date -u +%s`-$(PROJECT_SINCE))/(24*60*60))))
AUTO_COUNT_LOG:=$(shell git log --since=midnight --oneline|wc -l|tr -d " ")
COMMIT:=4b825dc
REVIEWDOG:=| reviewdog -efm='%f:%l:%c: %m' -diff="git diff $(COMMIT) HEAD"

GO:=$(shell (type go1.11rc2 > /dev/null 2>&1) && echo go1.11rc2 || echo go)
GOM:=GO111MODULE=on $(GO)
GOPATH:=$(shell $(GO) env GOPATH)
PKG:=$(shell $(GO) list)
GOBIN:=$(notdir $(PKG))
GOLIST:=$(shell $(GO) list ./...)
GODIR:=$(patsubst $(PKG)/%,%,$(wordlist 2,$(words $(GOLIST)),$(GOLIST)))

GOSRC:=$(shell find . -type d -name vendor -prune -or -type f -name "*.go" -print)
CEL:=$(shell find . -type d -name vendor -prune -or -type f -name "*.cel.txt" -print)
GOGEN:=$(patsubst %.cel.txt,%_gen.go,$(CEL))

APP_DIR_PATH:=app
GOPHERJS:=$(APP_DIR_PATH)/web/gopher.js

LIBGO:=$(wildcard lib/*.go)
GOLIB:=$(LIBGO:.go=.so)
.SUFFIXES: .go .so
.go.so: vendor
	$(GO) build -buildmode=c-shared -o $@ $<


.PHONY: all
all: $(GOBIN) $(GOLIB) $(APP_DIR_PATH)/build
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


vendor: go.mod
	$(GOM) mod vendor

%_gen.go: %.go %.cel.txt
	go generate ./...

$(GOBIN): vendor $(GOSRC)
	$(GOM) build -ldflags=" \
-X main.serial=$(AUTO_COUNT_SINCE).$(AUTO_COUNT_LOG) \
-X main.hash=$(shell git describe --always --dirty=+) \
-X \"main.build=$(shell LANG=en date -u +'%b %d %T %Y')\" \
"

.PHONY: go-get
go-get: $(GOSRC)
	echo > go.mod
	rm -rf vendor
	$(GOM) build -ldflags=" \
-X main.serial=$(AUTO_COUNT_SINCE).$(AUTO_COUNT_LOG) \
-X main.hash=$(shell git describe --always --dirty=+) \
-X \"main.build=$(shell LANG=en date -u +'%b %d %T %Y')\" \
"

$(GOPHERJS): vendor $(GOSRC) $(GOGEN)
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
	rm -f $(GOBIN) $(GOLIB) $(wildcard lib/*.h)
	rm -rf vendor $(APP_DIR_PATH)/build
	find . -name .DS_Store -delete

.PHONY: test
test: vendor
	$(GO) test $(PKG)/...

.PHONY: bench
bench:
	$(GO) test -bench . -benchmem -count 5 -run none $(PKG)/... | tee bench/now.txt
	[ -f bench/before.txt ] && ( type benchcmp > /dev/null 2>&1 ) && benchcmp bench/before.txt bench/now.txt || :

.PHONY: lint
lint: vendor
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
	-$(GO) vet $(GOLIST) $(REVIEWDOG)
	@echo
	-$(GO) vet -shadow $(GOLIST) $(REVIEWDOG)
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
