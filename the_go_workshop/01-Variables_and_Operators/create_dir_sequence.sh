#!/bin/bash
#mkdir  ./Chapter{1..15}


for i in $(seq -f "%02g" 2 4);
	do mkdir ./act_1.${i} ;
done
