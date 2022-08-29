#include <syscall.h>
#include <time.h>

long _syscall(int num, void *a0, void *a1, void *a2, void *a3, void *a4, void *a5);
long sys_fork();
long sys_execve(char *filename, char **argv, char **envp);
long unsigned sys_nanosleep(struct timespec *req, struct timespec *rem);