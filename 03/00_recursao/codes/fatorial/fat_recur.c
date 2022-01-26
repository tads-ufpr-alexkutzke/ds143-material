#include <stdio.h>

int fat(int n)
{
  if(n == 1) return(1);
  return(n * fat(n-1));
}

int fat2(int n)
{
  int t,i;
  for(t=1,i=1;i<=n;i++) t*=i;
  return(t);
}

int main(){
  int x;
  scanf("%d",&x);

  printf("%d\n",fat(x));
  printf("%d\n",fat2(x));
}
