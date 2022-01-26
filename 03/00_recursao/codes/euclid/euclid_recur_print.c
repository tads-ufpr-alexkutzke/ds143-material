#include <stdio.h>

int mdc_print(int m, int n, int tabs)
{
  int i;

  for(i=0;i<tabs;i++) printf("  ");
  printf("mdc(%d,%d)\n",m,n);

  if(n == 0) return(m);
  return(mdc_print(n, m % n,tabs+1));
}

int main(){
  int x,y;
  scanf("%d %d",&x,&y);

  printf("%d\n",mdc_print(x,y,0));
}
