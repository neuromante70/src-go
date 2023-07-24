package main

/*
There is no do-while loop in Go. To emulate the C/Java code
do {
	work();
} while (condition);

you may use a for loop in one of these ways:

1)
skip the init and post statemens
ok := true
for ok = condition


1)
for ok := true; ok; ok = condition {
	work()
}

2)
for {
	work()
	if !condition {
		break
	}
}

*/

func main() {

}
