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
	@python $(BASEDIR)/scripts/make_library_gcc.py $(BASEDIR) $(BASEDIR)/build/output/roslib/gcc $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_rtthread.py $(BASEDIR) $(BASEDIR)/build/output/roslib/rtthread $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_java.py $(BASEDIR) $(BASEDIR)/build/output/roslib/java $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_python.py $(BASEDIR) $(BASEDIR)/build/output/roslib/python $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_go.py $(BASEDIR) $(BASEDIR)/build/output/roslib/go $(BASEDIR)
	@rm -rf ${BASEDIR}/build/output/bin/python-packages
	@cp -rf $(BASEDIR)/build/output/roslib/python ${BASEDIR}/build/output/bin/python-packages
	@cd $(BASEDIR)/build/CMake; cmake $(BASEDIR); make
	@rm -rf $(BASEDIR)/dist/roslib; cp -a $(BASEDIR)/build/output/roslib $(BASEDIR)/dist
	
clean:
	@if [ -d $(BASEDIR)/build ]; then \
		rm -rf $(BASEDIR)/build; \
	fi
