package main

/*
There is no do-while loop in Go. To emulate the C/Java code
do {
	work();
} while (condition);

you may use a for loop in one of these ways:

1) skip the init and post statemens
ok := true
for ok = condition {
	work()
}

2) To ensure that it executes at least once, the <condition> must be true for the first iteration and be dependent on the loop stop condition in subsequent iterations:

for next := true; next; next=<condition> {
    <loopBody>
}

3) 
for {
	work()
	if !condition {
		break
	}
}

*/

func main() {

}
