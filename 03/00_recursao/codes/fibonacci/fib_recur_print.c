#include <stdio.h>

int fib_print(int n,int tabs)
{
  int i;

  for(i=0;i<tabs;i++) printf("  ");
  printf("fib(%d)\n",n);

  if(n == 0) return(0);
  if(n == 1) return(1);
  return(fib_print(n-1,tabs+1) + fib_print(n-2,tabs+1));
}

int main(){
  int x;
  scanf("%d",&x);

  printf("%d\n",fib_print(x,0));
}
