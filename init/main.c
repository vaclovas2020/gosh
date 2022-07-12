#include "util.h"
#include "sys.h"

int main()
{
    sys_print("Gosh v0.0.1 starting...\n");

    execute_process("/bin/gosh");

    while (1)
    {
        sleep_sec(1);
    }
    return 0;
}