
#include "sys.h"
#include "util.h"

void sys_print(char *msg)
{
    _syscall(SYS_write, (void *)1 /* stdout */, (void *)msg, (void *)_strlen(msg), 0, 0, 0);
}

long sys_fork()
{
    return _syscall(SYS_fork, 0, 0, 0, 0, 0, 0);
}

long sys_execve(char *filename, char **argv, char **envp)
{
    return _syscall(SYS_execve, (void *)filename, (void *)argv, (void *)envp, 0, 0, 0);
}

long unsigned sys_nanosleep(struct timespec *req, struct timespec *rem)
{
    return _syscall(SYS_nanosleep, (void *)req, (void *)rem, 0, 0, 0, 0);
}