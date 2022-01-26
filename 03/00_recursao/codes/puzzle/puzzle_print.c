#include <stdio.h>

int puzzle_print(int n,int tabs)
{
  int i;
  for(i=0;i<tabs;i++) printf("  ");
  printf("puzzle(%d)\n",n);
  if(n == 1) return(1);
  if(n % 2 == 0)
    return(puzzle_print(n/2,tabs+1));
  else return(puzzle_print(3*n+1,tabs+1));
}

int main(){
  int x;
  scanf("%d",&x);

  printf("%d\n",puzzle_print(x,0));
}
