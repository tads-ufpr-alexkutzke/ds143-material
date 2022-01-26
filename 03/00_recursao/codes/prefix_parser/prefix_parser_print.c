#include <stdio.h>

char a[100]; int i;

int eval(int tabs)
{
  int x = 0,j;

  while(a[i] == ' ') i++;

  for(j=0;j<tabs;j++) printf("  ");
  printf("eval()\n");

  if(a[i] == '+') { i++; return(eval(tabs+1) + eval(tabs+1)); }
  if(a[i] == '*') { i++; return(eval(tabs+1) * eval(tabs+1)); }
  while((a[i] >= '0') && (a[i] <= '9'))
    x = 10*x + (a[i++]-'0');
  return(x);
}

int main(){
  fgets(a,99,stdin);

  printf("%d\n",eval(0));
}
