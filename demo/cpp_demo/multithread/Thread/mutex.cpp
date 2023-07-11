#include <iostream>
#include <thread>
#include <mutex>

using namespace std;

volatile int counter(0);
std::mutex mtx;

void attempt_10k_increases()
{
    for ( int i = 0; i < 10; ++i)
    {
        if(mtx.try_lock())
        {
            std::cout<< "thread id:" << std::this_thread::get_id() 
                << "counter=" << counter <<endl;
            ++counter;
            mtx.unlock();
        }
        else
        {
//            std::cout<< "thread id:" << std::this_thread::get_id() 
 //               << "false to get lock" <<endl;
        }
    }
}

int main(int argc, const char *argv[])
{
    std::thread th[10]; 
    std::cout << " spawing 10 thread... " <<std::endl;
    for ( int i = 0; i < 10; ++i)
    {
        th[i] = std::thread(attempt_10k_increases);
    }
    std::cout << "Done  spawing 10 thread " <<std::endl;
    for(auto& t: th)
    {
        t.join();
    }
    

    std::cout<< "All thread join, final counter= " << counter  << endl;
    return EXIT_SUCCESS;
}
