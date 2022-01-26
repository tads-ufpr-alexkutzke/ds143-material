#include <stdio.h>

int mdc(int m, int n)
{
  if(n == 0) return(m);
  return(mdc(n, m % n));
}

int main(){
  int x,y;
  scanf("%d %d",&x,&y);

  printf("%d\n",mdc(x,y));
}
