# Makefile for building Chaos Center
# Reference Guide - https://www.gnu.org/software/make/manual/make.html

#
# Internal variables or constants.
# NOTE - These will be executed when any make target is invoked.
#

backend-services-checks:
	@echo "------------------"
	@echo "--> checking code style [backend]"
	@echo "------------------"
	@fmtRes=$$(gofmt -d $$(find . -path ./vendor -prune -o -name '*.go' -print)); \
	if [ -n "$${fmtRes}" ]; then \
		echo "gofmt checking failed!" && echo "$${fmtRes}" \
		&& echo "Please ensure you are using $$($(GO) version) for formatting code." \
		&& exit 1;\
	fi