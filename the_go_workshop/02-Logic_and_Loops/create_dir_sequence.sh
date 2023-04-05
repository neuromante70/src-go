#!/bin/bash
#mkdir  ./ex_2.{1..15}


for i in $(seq -f "%02g" 1 3);
	do mkdir ./act_2.${i} ;
done
