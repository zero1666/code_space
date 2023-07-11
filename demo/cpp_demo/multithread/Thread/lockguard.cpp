#include <iostream>
#include <thread>
#include <mutex>
#include <chrono>
#include <stdexcept>

using namespace std;

std::mutex mtx;
std::timed_mutex mtx_time;

void fireworks()
{
    while( !mtx_time.try_lock_for(std::chrono::milliseconds(200)))
    {
        //std::cout << "- " <<std::this_thread::get_id()<< "-";
        std::cout << "- " ;
    }
    std::this_thread::sleep_for(std::chrono::milliseconds(1000));
    std::cout << std::this_thread::get_id(); 
    std::cout << "*\n"; 
    mtx_time.unlock();
}

void print_even(int x)
{
    if(x%2 == 0) std::cout << x << "is even\n";
    else throw (std::logic_error("not even"));
}

void print_thread_id(int id)
{
    try {
        std::lock_guard<std::mutex> lck(mtx);
        print_even(id);
    }
    catch (std::logic_error&)
    {
        std::cout << "[exception caught]\n";
    }
}

int main(int argc, const char *argv[])
{
    std::thread th[10]; 
    std::cout << " spawing 10 thread... " <<std::endl;
    for ( int i = 0; i < 10; ++i)
    {
        th[i] = std::thread(print_thread_id, i+1);
    }
    std::cout << "Done  spawing 10 thread " <<std::endl;
    for(auto& t: th)
    {
        t.join();
    }
    

    std::cout<< "All thread join " << endl;
    return EXIT_SUCCESS;
}
