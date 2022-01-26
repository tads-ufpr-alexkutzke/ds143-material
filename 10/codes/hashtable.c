#include<stdio.h>
#include<string.h>
#include<stdlib.h>

int string_hash(char * s) {
  char c;
  int p = 31, m = 97;
  int hash_value = 0, p_pow = 1;

  while (c = *s++){
      hash_value = (hash_value + (c - 'a' + 1) * p_pow) % m;
      p_pow = (p_pow * p) % m;
  }
  return(hash_value);
}

int main(){
  char s[255];

  fgets(s,255,stdin);
  s[strcspn(s, "\n")] = '\0';

  printf("%s = %d\n",s, string_hash(s));
}
