#include <stdio.h>

void print_time(int hour, int minute)
{
  printf("%d:%d\n", hour, minute);
}

int main(void)
{
  print_time(23, 59);
  return 0;
}
