#include <stdio.h>
#include <stdlib.h>

void merge(int * c, int * a, int N,
                    int * b, int M)
{
  int i,j,k;

  for(i=0,j=0,k=0; k<M+N; k++)
  {
    if(i == N){c[k] = b[j++]; continue;}
    if(j == M){c[k] = a[i++]; continue;}
    c[k] = (a[i] < b[j]) ? a[i++] : b[j++];
  }
}

int main(){
  int N,M,i;
  int *a, *b, *c;

  scanf("%d",&N);
  a = (int *) malloc(sizeof(int) * N);
  for(i=0; i<N; i++)
    scanf("%d",&a[i]);

  scanf("%d",&M);
  b = (int *) malloc(sizeof(int) * M);
  for(i=0; i<M; i++)
    scanf("%d",&b[i]);

  c = (int *) malloc(sizeof(int) * N+M);

  merge(c,a,N,b,M);
  for(i=0; i<N+M; i++)
    printf("%d ",c[i]);
  printf("\n");
}
