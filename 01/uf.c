#include <stdio.h>
#include <stdlib.h>

typedef struct {
  int * id;
  int count;
  int n;
} UF;

UF * init_UF(int n);
int count_UF(UF * uf);
int connected_UF(UF * uf, int p, int q);
int find_UF(UF * uf, int p);
void union_UF(UF * uf, int p, int q);

UF * init_UF(int n){
  int i;
  UF * uf;

  uf = malloc(sizeof(UF));

  uf->id = malloc(sizeof(int) * n);
  uf->n = n;
  uf->count = n;
  for (i = 0; i < uf->count; ++i) {
    uf->id[i] = i;
  }

  return(uf);
}

int count_UF(UF * uf){
  return(uf->count);
}

int connected_UF(UF * uf, int p, int q){
  return(find_UF(uf, p) == find_UF(uf,q));
}

int find_UF(UF * uf, int p){
  return(0);
}

void union_UF(UF * uf, int p, int q){
  return;
}

int main(){
  int n, p, q;
  UF * uf;

  scanf("%d", &n);
  
  uf = init_UF(n);

  scanf("%d %d", &p, &q);
  while(p > -1 && q > -1){

    if(!connected_UF(uf, p, q)){
      printf("%d %d\n", p, q);
      union_UF(uf, p, q);
    }

    scanf("%d %d", &p, &q);
  }
}
