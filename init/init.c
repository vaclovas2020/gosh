#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdio.h>
#include "init.h"

int main()
{
    int onefd = open("/dev/console", O_RDONLY, 0);
    dup2(onefd, 0); // stdin
    int twofd = open("/dev/console", O_RDWR, 0);
    dup2(twofd, 1); // stdout
    dup2(twofd, 2); // stderr

    if (onefd > 2)
        close(onefd);
    if (twofd > 2)
        close(twofd);
    int pid = start_process("/bin/gosh", "gosh");
    if (pid == 0)
    {
        while (1)
            ;
    }
}

int start_process(const char *path, const char *arg)
{
    int pid = fork();
    if (!pid)
    {
        printf("Process `%s` exit with code %d", path, execl(path, arg, NULL));
    }
    else
    {
        printf("Starting process `%s` (arg: %s) (pid: %d)...", path, arg, pid);
    }
    return pid;
}