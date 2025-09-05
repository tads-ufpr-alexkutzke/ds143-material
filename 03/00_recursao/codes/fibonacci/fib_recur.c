#include <stdio.h>

long int cont=0;

long int fib(long int n)
{
  cont++;
  if(n == 0) return(0);
  if(n == 1) return(1);
  return(fib(n-1) + fib(n-2));
}

long int main(){
  long int x;
  scanf("%ld",&x);

  printf("%ld\n",fib(x));
  printf("%ld\n",cont);

}
