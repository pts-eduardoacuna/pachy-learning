CMDIR := cmd
COMMANDS := image infer parse stats train

.PHONE: all

all compile docker push clean:
	@for dir in $(COMMANDS); do \
		echo '# Running make' $@ 'in' $(CMDIR)/$$dir; \
		$(MAKE) -C $(CMDIR)/$$dir $@; \
	done
