#include <iostream>
#include <thread>

using namespace std;
void f1(int n)
{
    for ( int i = 0; i < 5; ++i)
    {
        std::cout << "Thread1, n=" << n << " executing" <<endl;;
        std::this_thread::sleep_for(std::chrono::milliseconds(13));
    }
}

void f2(int& n)
{
    for ( int i = 0; i < 5; ++i)
    {
        std::cout << "Thread2, n=" << n << "executing" <<endl;;
        ++n;
        std::this_thread::sleep_for(std::chrono::milliseconds(10));
    }
}
void thread_task(int n)
{
    std::this_thread::sleep_for(std::chrono::seconds(n));
    std::cout<< "thread id:" << std::this_thread::get_id() 
        << "pause " << n <<" seconds" << std::endl;
}

int main(int argc, const char *argv[])
{
    std::thread th[5]; 
    std::cout << " spawing 5 thread... " <<std::endl;
    for ( int i = 4; i >= 0; --i)
    {
        th[i] = std::thread(thread_task, i+1);
    }
    std::cout << "Done  spawing 5 thread " <<std::endl;
    for(auto& t: th)
    {
        t.join();
    }
    

    std::cout<< "All thread join "  << endl;
    return EXIT_SUCCESS;
}
