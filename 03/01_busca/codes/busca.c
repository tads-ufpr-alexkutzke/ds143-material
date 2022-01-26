#include <stdio.h>

#define N 1000000

void preenche_vetor(long v[], long n){
  long i;
  long j;
  for(i=0,j=0; j<n;i+=2) v[j++] = i;
}

void imprime_vetor(long v[], long n){
  long i;
  for(i=0; i<n;i++) printf("v[%ld] = %ld\n",i,v[i]);
}

long busca_sequencial(long v[], long n, long elem){
  long i;
  for(i=0; i<n;i++)
    if (v[i] == elem){
      printf("%ld comparacoes\n",i);
      return i;
    }
  printf("%ld comparacoes\n",i);
  return(-1);
}

long busca_binaria(long v[], long n, long elem){
  long l,r,m,comp=0;
  l = 0;
  r = n-1;

  while(l <= r){
    m = (l+r)/2;
    if(v[m] == elem)
    {
      printf("%ld comparacoes\n",comp);
      return(m);
    }
    else if(v[m] < elem)
      l = m + 1;
    else if(v[m] > elem)
      r = m - 1;
    comp++;
  }
  printf("%ld comparacoes\n",comp);
  return(-1);
}

long busca_binaria_recur(long v[], long n, long elem, long l, long r){
  long m = (l+r)/2;

  if(l > r)
    return(-1);

  if(v[m] == elem)
    return(m);
  else if(v[m] < elem)
    return(busca_binaria_recur(v,n,elem,m+1,r));
  else if(v[m] > elem)
    return(busca_binaria_recur(v,n,elem,l,m-1));
}

int main(){
  long v[N],x,r;

  preenche_vetor(v,N);
  //imprime_vetor(v,N);

  scanf("%ld",&x);

  r = busca_sequencial(v,N,x);
  //r = busca_binaria(v,N,x);
  //r = busca_binaria_recur(v,N,x,0,N-1);
  if(r != -1)
    printf("O elemento %ld esta na posicao %ld do vetor.\n",x,r);
  else
    printf("O elemento %ld nao pertence ao vetor.\n",x);
}
