#include "init.h"

long _syscall(int num, void *a0, void *a1, void *a2, void *a3, void *a4, void *a5);

unsigned long _strlen(const char *sz)
{
    int count = 0;
    while (*sz++)
    {
        count++;
    }
    return count;
}

void delay(int ticks)
{
    for (int i = 0; i < ticks; i++)
        ;
}

void str_print(const char *msg)
{
    _syscall(SYS_write, (void *)1 /* stdout */, (void *)msg, (void *)_strlen(msg), 0, 0, 0);
}
