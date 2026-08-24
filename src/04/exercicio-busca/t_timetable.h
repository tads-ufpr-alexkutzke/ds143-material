#ifndef __T_TIMETABLE__
#define __T_TIMETABLE__

#include "t_time.h"

typedef struct {
	t_time * key;
	char * value;
} t_timetable_item;

typedef struct {
	t_timetable_item * table;
	int size,n;
} t_timetable;

t_timetable * t_timetable_init(int size);
void t_timetable_put_seq(t_timetable * tt, t_time * key, char * value);
char * t_timetable_get_seq(t_timetable * tt, t_time * key);
void t_timetable_put(t_timetable * tt, t_time * key, char * value);
char * t_timetable_get(t_timetable * tt, t_time * key);
void t_timetable_print(t_timetable * tt);

#endif
