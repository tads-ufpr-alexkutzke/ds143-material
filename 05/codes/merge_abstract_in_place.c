#include <stdio.h>
#include <stdlib.h>

void merge(int * a, int l, int m, int r)
{
  int i,j,k;
  int *aux;

  aux = (int *) malloc(sizeof(int) * r+1);

  for(i=m+1; i > l; i--) aux[i-1] = a[i-1];
  for(j=m; j < r; j++) aux[r+m-j] = a[j+1];

  for(k=l; k<=r; k++)
  {
    if(aux[j] < aux[i])
      a[k] = aux[j--];
    else
      a[k] = aux[i++];
  }
}

int main(){
  int N,M,i,x;
  int *a;

  scanf("%d",&N);
  scanf("%d",&M);
  a = (int *) malloc(sizeof(int) * (N+M));
  for(i=0; i<N; i++)
    scanf("%d",&a[i]);
  for(i=N; i<N+M; i++)
    scanf("%d",&a[i]);

  merge(a,0,N-1,N+M-1);
  for(i=0; i<N+M; i++)
    printf("%d ",a[i]);
  printf("\n");
}
