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
	@python $(BASEDIR)/scripts/make_library_gcc.py $(BASEDIR) $(BASEDIR)/build/output/roslib/gcc/tiny_ros $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_rtthread.py $(BASEDIR) $(BASEDIR)/build/output/roslib/rtthread $(BASEDIR)
	@python $(BASEDIR)/scripts/make_library_java.py $(BASEDIR) $(BASEDIR)/build/output $(BASEDIR)
	@cd $(BASEDIR)/build/CMake; cmake $(BASEDIR); make
	@rm -rf $(BASEDIR)/dist/*; cp -a $(BASEDIR)/build/output/roslib $(BASEDIR)/dist
	
clean:
	@if [ -d $(BASEDIR)/build ]; then \
		rm -rf $(BASEDIR)/build; \
	fi
