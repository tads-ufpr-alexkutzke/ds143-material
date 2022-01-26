// Weighted Quick Union para Union-find
// Prof. Alex Kutzke
//
// Baseado no exemplo dado no livro Algorithms,4th Edition 
// de Robert Sedgewick e Kevin Wayne.
// Versão original em java: https://algs4.cs.princeton.edu/15uf/WeightedQuickFindUF.java

#include <stdio.h>
#include <stdlib.h>

typedef struct {
  int * parent;   // parent[i] = parent of i
  int * size;     // size[i] = number of elements in subtree rooted at i 
  int n;
  int count;
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

  uf->n = n;
  uf->parent = malloc(sizeof(int) * n);
  uf->size = malloc(sizeof(int) * n);
  uf->count = n;
  for (i = 0; i < uf->count; ++i) {
    uf->parent[i] = i;
    uf->size[i] = 1;
  }

  return(uf);
}

int count_UF(UF * uf){
  return(uf->count);
}

int connected_UF(UF * uf, int p, int q){
  return(find_UF(uf,p) == find_UF(uf,q));
}

int find_UF(UF * uf, int p){
  int orig_p = p;
  while(p != uf->parent[p])
    p = uf->parent[p];
  
  //uf->parent[orig_p] = p;
  return(p);
}

void union_UF(UF * uf, int p, int q){
  int root_p = find_UF(uf,p);
  int root_q = find_UF(uf,q);

  int * parent = uf->parent;
  int * size = uf->size;

  if (root_p == root_q) return;

  if(size[root_p] < size[root_q]){
    parent[root_p] = root_q;
    size[root_q] += size[root_p];
  }
  else {
    parent[root_q] = root_p;
    size[root_p] += size[root_q];
  }
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
