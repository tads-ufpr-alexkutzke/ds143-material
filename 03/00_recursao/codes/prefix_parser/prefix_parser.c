#include <stdio.h>

char a[100]; int i;

int eval()
{
  int x = 0;
  while(a[i] == ' ') i++;
  if(a[i] == '+') { i++; return(eval() + eval()); }
  if(a[i] == '*') { i++; return(eval() * eval()); }
  while((a[i] >= '0') && (a[i] <= '9'))
    x = 10*x + (a[i++]-'0');
  return(x);
}

int main(){
  fgets(a,99,stdin);

  printf("%d\n",eval());
}
