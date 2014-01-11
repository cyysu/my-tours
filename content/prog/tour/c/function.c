#include <stdio.h>

void newline(void)
{
  printf("\n");
}

int main(void)
{
  printf("First Line.\n");
  newline();
  printf("Second Line.\n");
  return 0;
}
