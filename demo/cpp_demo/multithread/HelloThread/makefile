all:Thread

CC=g++
CPPFLAGS=-Wall -std=c++11 -ggdb

Thread:Thread.o
	$(CC)  -o $@ $^

Thread.o:Thread.cpp
	$(CC) $(CPPFLAGS) -o $@ -c $^

.PHONY:
	clean

clean:
	rm Thread.o Thread
