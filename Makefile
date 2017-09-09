CMDIR := cmd
COMMANDS := image infer parse stats train

.PHONE: all

all compile docker push clean: test
	@echo ""
	@echo "********************************"
	@echo "* Building standalone commands *"
	@echo "********************************"
	@echo ""
	@for dir in $(COMMANDS); do \
		echo ""; \
		echo '* Running make' $@ 'for' $(CMDIR)/$$dir; \
		echo ""; \
		$(MAKE) -C $(CMDIR)/$$dir $@; \
	done

test: lint
	@echo ""
	@echo "*****************"
	@echo "* Running tests *"
	@echo "*****************"
	@echo ""
	go test ./...

lint:
	@echo ""
	@echo "***********************"
	@echo "* Checking code style *"
	@echo "***********************"
	@echo ""
	go fmt ./...
	@echo ""
	go vet ./...
	@echo ""
	go list ./... | xargs golint -set_exit_status=true
