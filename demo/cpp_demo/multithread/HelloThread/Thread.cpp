#include <stdio.h>
#include <stdlib.h>

#include <thread>
#include <iostream>

void thread_task()
{
    std::cout<<"hello thread" << std::endl;
}

int main(int argc, const char *argv[])
{
    std::thread t(thread_task);
    t.join();
    return EXIT_SUCCESS;
}
