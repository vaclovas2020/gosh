#pragma once
#include <syscall.h>

long _syscall(int num, void *a0, void *a1, void *a2, void *a3, void *a4, void *a5);
unsigned long _strlen(char * sz);
void delay(int ticks);
void str_print(char *msg);