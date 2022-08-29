#include "util.h"

void sleep_sec(int sec){
    struct timespec tm;
    tm.tv_nsec = 0;
    tm.tv_sec = sec;
    sys_nanosleep(&tm, NULL);
}

int execute_process(char *filename)
{
    long pid = sys_fork();
    if (!pid){
        char * argv[2];
        argv[0] = filename;
        argv[1] = 0;
        char * envp[1];
        envp[0] = 0;
        return sys_execve(filename, argv, envp);
    }
}
