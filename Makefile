check_env:
	@echo "Current dir:" $(CURDIR)
	@test -f $(CURDIR)/.env && echo ".env found" || echo ".env not found"

grm:
	@/bin/bash -c 'set -a && source $(CURDIR)/.env && set +a && go run main.go'