#include "util.h"
#include "sys.h"

int main()
{
    execute_process("/bin/gosh");

    while (1)
    {
        sleep_sec(1);
    }
    return 0;
}