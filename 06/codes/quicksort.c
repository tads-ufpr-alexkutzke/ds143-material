#include <stdio.h>
#include <stdlib.h>

/* Arrange the N elements of ARRAY in random order.
   Only effective if N is much smaller than RAND_MAX;
   if this may not be the case, use a better random
number generator. */
void shuffle(int *array, size_t n)
{
  if (n > 1)
  {
    size_t i;
    for (i = 0; i < n - 1; i++)
    {
      size_t j = i + rand() / (RAND_MAX / (n - i) + 1);
      int t = array[j];
      array[j] = array[i];
      array[i] = t;
    }
  }
}

void swap(int * a, int i, int j)
{
  int aux = a[i];
  a[i] = a[j];
  a[j] = aux;
}

int partition(int * a, int lo, int hi)
{
  int i = lo, j = hi+1;
  while(1)
  {
    while(a[++i] < a[lo])
    if(i == hi) break;

    while(a[lo] < a[--j])
    if(j == lo) break;

    if(i >= j) break;
    swap(a, i, j);
  }

  swap(a, lo, j);
  return(j);
}

void quicksort_(int * a, int lo, int hi)
{
  if(hi <= lo) return;
  int j = partition(a, lo, hi);
  quicksort_(a, lo, j-1);
  quicksort_(a, j+1, hi);
}

void quicksort(int * a, int n)
{
  shuffle(a,n);
  quicksort_(a,0,n-1);
}

int main()
{
  int i;
  int a[10] = {5,2,8,4,10,9,7,6,3,1};

  quicksort(a,10);
  for(i=0; i<10; i++)
  printf("%d, ", a[i]);

  printf("\n");

  return(0);
}
