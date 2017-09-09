REPO := github.com/pts-eduardoacuna/pachy-learning

CMDIR := cmd
COMMANDS := image infer parse stats train

DOCDIR := doc
PACKAGES := learning

.PHONE: all

all compile docker push clean: test doc
	@echo "********************************"
	@echo "* Building standalone commands *"
	@echo "********************************"
	@for dir in $(COMMANDS); do \
		echo ""; \
		echo "* Running make" $@ "for" $(CMDIR)/$$dir; \
		echo ""; \
		$(MAKE) -C $(CMDIR)/$$dir $@; \
	done

test: lint
	@echo "*****************"
	@echo "* Running tests *"
	@echo "*****************"
	@echo ""
	go test ./...
	@echo ""

doc: lint
	@echo "***************************"
	@echo "* Producing documentation *"
	@echo "***************************"
	@for pkg in $(PACKAGES); do \
		echo ""; \
		echo "* Making docs for" $(REPO)/$$pkg; \
		mkdir -p $(DOCDIR)/$$pkg; \
		godoc2md -v $(REPO)/$$pkg > $(DOCDIR)/$$pkg/README.md; \
		echo ""; \
	done

lint: header
	@echo "***********************"
	@echo "* Checking code style *"
	@echo "***********************"
	@echo ""
	go fmt ./...
	@echo ""
	go vet ./...
	@echo ""
	go list ./... | xargs golint -set_exit_status=true
	@echo ""

header:
	@echo "                      __                __                      _            "
	@echo "    ____  ____ ______/ /_  __  __      / /__  ____ __________  (_)___  ____ _"
	@echo "   / __ \\/ __ \`/ ___/ __ \\/ / / /_____/ / _ \\/ __ \`/ ___/ __ \\/ / __ \\/ __ \`/"
	@echo "  / /_/ / /_/ / /__/ / / / /_/ /_____/ /  __/ /_/ / /  / / / / / / / / /_/ / "
	@echo " / .___/\\__,_/\\___/_/ /_/\\__, /     /_/\\___/\\__,_/_/  /_/ /_/_/_/ /_/\\__, /  "
	@echo "/_/                     /____/                                      /____/   "
	@echo ""
