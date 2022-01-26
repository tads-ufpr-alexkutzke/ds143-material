#include <stdio.h>

typedef struct node {
   int info;
   struct node* next;
}Lista;

Lista* criar (void) {
  return NULL;
}

Lista* inserir (Lista* lista, int elem) {
  Lista* novo =(Lista*)malloc(sizeof(Lista));
  novo->info = elem;
  novo->next = lista;
  return novo;
}

void imprimir (Lista *lista) {
  Lista *v; /*var. para percorrer a lista*/
  for (v = lista; v != NULL; v = v->next) {
     printf("Valor: %d\n", v->info);
  }
}

int tamanho (Lista *lista) {
  if(!lista) return(0);
  return(1 + tamanho(lista->next));
}

int maximo (Lista *lista) {
  int tmp;

  printf("--- %d\n",lista->info);
  if(lista->next == NULL) return(lista->info);
  tmp = maximo(lista->next);
  if(tmp < lista->info) return(lista->info);
  else return(tmp);
}

void liberar (Lista *lista) {
  while (lista != NULL) {
    Lista *aux = lista->next; /*guarda ref. p/ prox.*/
    free (lista); /*libera a memoria apontada por v*/
    lista = aux; /*faz v apontar p/ o prox. elem.*/
  }
}

int main () {
  Lista *lista;
  lista = criar ();
  lista = inserir (lista, 1);
  lista = inserir (lista, 2);
  lista = inserir (lista, 4);
  lista = inserir (lista, 400);
  lista = inserir (lista, 4);
  lista = inserir (lista, 4);
  lista = inserir (lista, 4);
  imprimir (lista);
  printf("Tamanho da lista: %d\n",tamanho(lista));
  printf("Maior da lista: %d\n",maximo(lista));
  liberar (lista);
  return 0;
}
