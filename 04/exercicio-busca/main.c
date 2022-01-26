#include <stdio.h>
#include <stdlib.h>

#include "t_time.h"
#include "t_timetable.h"

int main(){
	t_time *ta;
	t_timetable * tt;
	char * str;
	size_t len;

	int h,m,s,size;

	scanf("%d",&size);
	tt = t_timetable_init(size);

	scanf("%d:%d:%d",&h,&m,&s);
	while(h >= 0){
		getchar();

		ta = t_time_init(h,m,s);
		
		// read string
		str = NULL;
		len = getline(&str, &len, stdin);
		str[len-1] = '\0';

		t_timetable_put(tt,ta,str);
		//t_timetable_put_seq(tt,ta,str);

		scanf("%d:%d:%d",&h,&m,&s);
	}

	t_timetable_print(tt);

	scanf("%d:%d:%d",&h,&m,&s);
	while(h >= 0){
		ta = t_time_init(h,m,s);
		str = t_timetable_get(tt,ta);
		//str = t_timetable_get_seq(tt,ta);

		if(str)
			printf("%02d:%02d:%02d => %s\n",h,m,s,str);
		else
			printf("%02d:%02d:%02d => nao encontrado\n",h,m,s);

		scanf("%d:%d:%d",&h,&m,&s);
	}	
}
