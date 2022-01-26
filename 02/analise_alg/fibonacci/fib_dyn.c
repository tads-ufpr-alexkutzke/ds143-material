#include <stdio.h>

long int knownF[100];

long int fib(long int n)
{
  long int t;

  if(knownF[n] != 0) return knownF[n];
  if(n == 0) return(0);
  if(n == 1) return(1);
  t = fib(n-1) + fib(n-2);
  return knownF[n] = t;
}

long int main(){
  long int x,i;
  for(i=0;i<100;i++) knownF[i]=0;
  scanf("%ld",&x);

  printf("%ld\n",fib(x));
}
