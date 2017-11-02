PATH_PRJ_ROOT:=$(if $(PATH_PRJ_ROOT), $(PATH_PRJ_ROOT), .)
all:
	$(MAKE) -f makefile.32 clean
	$(MAKE) -f makefile.32 all
	$(MAKE) -f makefile.64 clean
	$(MAKE) -f makefile.64 all
clean:
	$(MAKE) -f makefile.32 clean
	$(MAKE) -f makefile.64 clean 
