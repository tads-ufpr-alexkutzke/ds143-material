#include<stdio.h>
#include<stdlib.h>
#include<string.h>


void bubble_sort(long int * v, long int n){
  long int i, j, aux, x;

	for (i = 0; i < n; ++i) {
		for (j = 0; j < n-i-1; j++) {
			if(v[j] > v[j+1]){
				aux = v[j];
				v[j] = v[j+1];
				v[j+1] = aux;
			}
		}
	}
}

void insertion_sort(long int * v, long int n){
  long int i, j, aux, x;
	
	for(i=+1; i<n; i++){
		x = v[i];
		j = i;
		while(x < v[j-1] && j > 0){
			v[j] = v[j-1];
			j--;
		}
		v[j] = x;
	}
}

void selection_sort(long int * v, long int n){
  long int i, j, aux, m;

  for(i=0; i<n-1; i++){
    m = i;
    for(j=i+1; j<n; j++){
      if(v[j] < v[m])
        m = j;
    }
    aux  = v[m];
    v[m] = v[i];
    v[i] = aux;
  }

}

void print_v(long int * v, long int n){
	long int i;
	for (i = 0; i < n; ++i) {
		printf("%d ", v[i]);
	}
	printf("\n");
}

int main(long int argc, char ** argv){
	int print = 0;
  long int i,n;
  long int * v;
  char * alg;

  srand(1);

  alg = argv[1];
  n = atoi(argv[2]);

	if(argc > 3) print = 1;

  v = malloc(sizeof(long int) * n);
  for(i=0; i<n; i++) v[i] = abs(rand()) % n;

	if(print) print_v(v,n);

  if(strcmp(alg, "selection") == 0){
    selection_sort(v,n);
  }
  else if(strcmp(alg, "insertion") == 0){
		insertion_sort(v,n);
  }
  else if(strcmp(alg, "bubble") == 0){
		bubble_sort(v,n);
  }


	if(print) print_v(v,n);

  for(i=n-1; i>0; i--) 
    if(v[i] < v[i-1]){
      printf("Error!\n"); 
      exit(EXIT_FAILURE);
    }

	return(0);
}
