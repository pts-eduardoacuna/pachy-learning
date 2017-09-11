REPO := github.com/pts-eduardoacuna/pachy-learning

CMDIR := cmd
COMMANDS := image infer parse stats train

DOCDIR := doc
PACKAGES := learning mnist csnv json

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
	go test -v ./...
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
	@cat doc/img/header.txt
	@echo ""
