

#!/bin/bash
#mkdir  ./act_4.{1..15}


for i in $(seq -f "%02g" 1 2);
	do mkdir ./act_19.${i} ;
done

for i in $(seq -f "%02g" 1 2);
	do mkdir ./ex_19.${i} ;
done
