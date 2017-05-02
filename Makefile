DIRS = project1 project2

MAKECMDGOALS ?= all

$(MAKECMDGOALS): $(DIRS)
$(DIRS):
	@$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(DIRS) $(MAKECMDGOALS)
