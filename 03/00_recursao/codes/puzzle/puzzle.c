#include <stdio.h>

int puzzle(int n)
{
  if(n == 1) return(1);
  if(n % 2 == 0)
    return(puzzle(n/2));
  else return(puzzle(3*n+1));
}

int main(){
  int x;
  scanf("%d",&x);

  printf("%d\n",puzzle(x));
}
