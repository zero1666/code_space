#PATH_PRJ_ROOT是项目的根目录相对于当前目录的相对路径，必须保证正确
PATH_PRJ_ROOT:=$(if $(PATH_PRJ_ROOT), $(PATH_PRJ_ROOT), ./)

#规则是每个.cpp文件会生成一个对应名字的可执行文件

#需要在这里添加依赖的头文件和库
INC := 
LIB := 

#如果不是在本目录下，则需要定义以下变量
#SRC_DIR=./src/
#OBJS_DIR=./src/
#BIN_DIR=./lib/

#默认编译出来的可执行文件是strip掉符号表的，如果不想strip,请打开QZ_CGI_STRIP=nostrip
STRIP=nostrip

CC   = gcc
CXX  = g++
AR   = ar

BITS := 64
C_FLAGS := -ggdb -m$(BITS) -Wall -std=c++11 -ggdb 

SRCS?=$(wildcard $(SRC_DIR)*.cpp)
OBJS:=$(patsubst $(SRC_DIR)%.cpp,$(OBJS_DIR)%.o,$(SRCS))
DEPS:=$(patsubst $(SRC_DIR)%.cpp,$(OBJS_DIR)%.d,$(SRCS))

TARGET:=$(patsubst $(SRC_DIR)%.cpp,$(BIN_DIR)%,$(SRCS))

.PHONY: all cgi dep clean print $(INSTALL)

all:cgi $(INSTALL)
	
cgi:$(TARGET) $(OBJS) 

dep:$(DEPS)

$(BIN_DIR)%:$(OBJS_DIR)%.o $(DEP_LIB)
	@if( $(CXX) $(C_FLAGS) $(INC) $< $(DEP_LIB) $(LIB) -o $@ );then \
        echo -e "linking $@"[$(GREEN)OK$(RESET)]; \
        rm -f $<; \
        rm -f $@.d; \
     else \
        echo -e "linking $@"[$(RED)FAIL$(RESET)]; \
        exit 1;\
     fi
	@(if [ "$(STRIP)" != "nostrip" ];then strip $@;fi)

$(OBJS_DIR)%.o:$(SRC_DIR)%.cpp
	@if( $(CXX) $(C_FLAGS) $(INC) -c $< -o $@ );then \
		echo -e "Compiling $<"[$(GREEN)OK$(RESET)];\
	else \
		echo -e "Compiling $<"[$(RED)FAIL$(RESET)];\
		exit 1;\
	fi

#生成依赖文件，如果需要可以加上
ifndef NOT_COMPILE_DEP
ifneq ($(MAKECMDGOALS),clean)
-include $(DEPS)
endif

$(OBJS_DIR)%.d:$(SRC_DIR)%.cpp
	@echo -ne $(YELLOW)
	@if ($(CXX) -MM -MT "$(patsubst %.d,$@ %,$@)" $(C_FLAGS) $(INC) $< > $@ );then \
        echo -ne $(RESET);\
     else \
        echo -ne $(RESET); \
        exit 1;\
     fi
endif

clean:
	rm -f $(TARGET) $(OBJS) $(DEPS)

print:
	@echo "INC:$(INC)"
	@echo "LIB:$(LIB)"
	@echo "DEP_LIB:$(DEP_LIB)"
	@echo "SRCS:$(SRCS)"
	@echo "DEPS:$(DEPS)"
	@echo "OBJS:$(OBJS)"
	@echo "BIN_DIR:$(BIN_DIR)"
	@echo "OBJS_DIR:$(OBJS_DIR)"
	@echo "SRC_DIR:$(SRC_DIR)"
	@echo "C_FLAGS:$(C_FLAGS)"
