#
# Copyright (c) 2010, Pismo Labs Limited
#
# All rights reserved.
#
SHELL := /bin/bash

TOP=.

export BASEDIR=$(PWD)

target:
	@echo -e "\nGenerating tiny-ros..."
	@mkdir -p $(BASEDIR)/build/CMake
	@mkdir -p $(BASEDIR)/build/output/bin
	@python $(BASEDIR)/scripts/make_library_gcc.py $(BASEDIR) $(BASEDIR)/build/output/client_library/gcc $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_rtthread.py $(BASEDIR) $(BASEDIR)/build/output/client_library/rtthread $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_java.py $(BASEDIR) $(BASEDIR)/build/output/client_library/java $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_python.py $(BASEDIR) $(BASEDIR)/build/output/client_library/python $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_go.py $(BASEDIR) $(BASEDIR)/build/output/client_library/go $(BASEDIR)
	@rm -rf ${BASEDIR}/build/output/bin/python-packages
	@cp -rf $(BASEDIR)/build/output/client_library/python ${BASEDIR}/build/output/bin/python-packages
	@cd $(BASEDIR)/build/CMake; cmake $(BASEDIR); make
	@rm -rf $(BASEDIR)/dist/client_library
	@cp -a $(BASEDIR)/build/output/client_library $(BASEDIR)/dist
	@cp -a $(BASEDIR)/build/output/bin/*.exe $(BASEDIR)/dist/cygwin32
	
clean:
	@if [ -d $(BASEDIR)/build ]; then \
		rm -rf $(BASEDIR)/build; \
	fi
