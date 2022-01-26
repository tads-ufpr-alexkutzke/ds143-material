#include <stdio.h>
#include <stdlib.h>

#define RED   1
#define BLACK 0

typedef struct arvoreRB {
  int info;
  int cor;
  struct arvoreRB *esq;
  struct arvoreRB *dir;
} ArvoreRB;

int eh_no_vermelho(ArvoreRB * no){
  if(!no) return BLACK;
  return(no->cor == RED);
}

int buscar (ArvoreRB *a, int v) {
  if (a == NULL) { return 0; } /*Nao achou*/
  else if (v < a->info) {
    return buscar (a->esq, v);
  }
  else if (v > a->info) {
    return buscar (a->dir, v);
  }
  else { return 1; } /*Achou*/
}

void in_order(ArvoreRB *a){
  if(!a)
    return;
  in_order(a->esq);
  printf("%d ",a->info);
  in_order(a->dir);
}

void print(ArvoreRB * a,int spaces){
  int i;
  for(i=0;i<spaces;i++) printf(" ");
  if(!a){
    printf("//\n");
    return;
  }

  printf("%d\n", a->info);
  print(a->esq,spaces+2);
  print(a->dir,spaces+2);
}

int main(){
  ArvoreRB * a;

}
