#include <stdio.h>

int main(int argc, char *argv[])
{
  int hour, miniute, total_minute;
  hour = 1;
  minute = 20;
  total_minute = hour * 60 + minute;
  printf("%d:%d is %d minutes after 00:00\n", hour, minute, total_minute);
  return 0;
}
