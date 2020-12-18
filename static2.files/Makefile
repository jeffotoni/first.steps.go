# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on

static:
	@echo "\033[0;33m################################## build statik ##################################\033[0m"
	@statik -src=web/
	@echo "\033[0;32mstatik generated\033[0m  \033[0;33m[ok]\033[0m "
	@echo "--------------------------------------------------------------------------------------"
	GARCH=386 GOOS=linux go build -ldflags="-s -w" -o main main.go
	#upx main
	@echo "\033[0;32mserverworkshop generated\033[0m \033[0;33m[ok]\033[0m \n"

dinamic:
	@echo "--------------------------------------------------------------------------------------"
	GARCH=386  GOOS=linux go build -ldflags="-s -w" -o main main.go
	upx main
	#upx --brute serverworkshop
	@echo "\033[0;32mserverworkshop generated\033[0m \033[0;33m[ok]\033[0m \n"	