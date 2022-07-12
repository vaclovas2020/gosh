#include "init.h"
#define DELAY 1000000000

int main()
{

   delay(DELAY);
   str_print("Gosh v0.0.1");
   while(1){
    delay(DELAY);
    str_print("Tick!");
   }
   return 0;
}