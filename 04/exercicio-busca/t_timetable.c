#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include "t_time.h"
#include "t_timetable.h"

t_timetable * t_timetable_init(int size){
	t_timetable * tt;

	tt = (t_timetable *) malloc(sizeof(t_timetable));
	tt->table = (t_timetable_item*) malloc(sizeof(t_timetable_item)* size); tt->size = size;
	tt->n = 0;

	return(tt);
}

void t_timetable_put(t_timetable * tt,t_time * key, char * value){
	int i,j;

	if(tt->n == tt->size) exit(-1);
	
	i = 0;
  while(i < tt->n && t_time_cmp(key, tt->table[i].key) > 0){
		i++;
	}

	if(tt->n > i && t_time_cmp(key, tt->table[i].key) != 0){
		for (j = tt->n-1; j >= i; --j) {
			tt->table[j+1].key = tt->table[j].key;		
			tt->table[j+1].value = tt->table[j].value;		
		}
	}

	tt->table[i].key = key;
	tt->table[i].value = value;
	if(tt->n == i) tt->n++;
}

char * t_timetable_get(t_timetable * tt, t_time * key){
	int lo,hi,m,x;

	lo = 0;
	hi = tt->n-1;

	while(lo <= hi){
		m = (lo + hi)/2;

		x = t_time_cmp(tt->table[m].key, key);
		if(x == 0){
			return(tt->table[m].value);
		}
		else if(x == 1){
			hi = m - 1;
		}
		else{
			lo = m + 1;
		}
	}

	return(NULL);
}

void t_timetable_print(t_timetable * tt){
	int i, h, m, s;
	t_time * t;

	for (i = 0; i < tt->n; ++i) {
		t = tt->table[i].key;
		h = t_time_get_h(t);
		m = t_time_get_m(t);
		s = t_time_get_s(t);

		printf("[%02d] - %02d:%02d:%02d => %s\n",i, h, m, s, tt->table[i].value);
	}
}


void t_timetable_put_seq(t_timetable * tt,t_time * key, char * value){
	if(tt->n == tt->size) exit(-1);
	
	int i;

	for (i = 0; i < tt->n; ++i) {
		if(t_time_cmp(key, tt->table[i].key) == 0)
			break;
	}
	tt->table[i].key = key;
	tt->table[i].value = value;
	
	if(i == tt->n)
		tt->n++;
}

char * t_timetable_get_seq(t_timetable * tt, t_time * key){
	int i;

	for (i = 0; i < tt->n; ++i) {
		if(t_time_cmp(key, tt->table[i].key) == 0)
			return(tt->table[i].value);
	}

	return(NULL);
}
