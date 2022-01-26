#include<stdio.h>

long int fib(long int n)
{
   long int first = 1, second = 1, next, c, sum;

   sum = 0;
   for ( c = 0 ; c < n ; c++ ){
      if ( c <= 1 )
         next = c;
      else{
         next = first + second;
         first = second;
         second = next;
      }
   }

   return(next);
}

long int main(){
  long int x;
  scanf("%ld",&x);

  printf("%ld\n",fib(x));
}
