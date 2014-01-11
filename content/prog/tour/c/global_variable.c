#include <stdio.h>

int hour = 23, minute = 59;

void print_time(void)
{
  printf("%d:%d in print_time\n", hour, minute);
}

int main(void)
{
  print_time();
  printf("%d:%d in main\n", hour, minute);
  return 0;
}
