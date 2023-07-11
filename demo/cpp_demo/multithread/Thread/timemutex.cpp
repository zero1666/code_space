#include <iostream>
#include <thread>
#include <mutex>
#include <chrono>

using namespace std;

std::timed_mutex mtx;

void fireworks()
{
    while( !mtx.try_lock_for(std::chrono::milliseconds(200)))
    {
        //std::cout << "- " <<std::this_thread::get_id()<< "-";
        std::cout << "- " ;
    }
    std::this_thread::sleep_for(std::chrono::milliseconds(1000));
    std::cout << std::this_thread::get_id(); 
    std::cout << "*\n"; 
    mtx.unlock();
}

int main(int argc, const char *argv[])
{
    std::thread th[10]; 
    std::cout << " spawing 10 thread... " <<std::endl;
    for ( int i = 0; i < 10; ++i)
    {
        th[i] = std::thread(fireworks);
    }
    std::cout << "Done  spawing 10 thread " <<std::endl;
    for(auto& t: th)
    {
        t.join();
    }
    

    std::cout<< "All thread join " << endl;
    return EXIT_SUCCESS;
}
