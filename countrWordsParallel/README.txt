go run frequencyWithoutChannel.go shake.txt read1.txt read2.txt
----above run will take an avg of 500micro sec

go run frequency.go shake.txt read1.txt read2.txt
----above code takes an avg of 250 micro sec

----thus with channels a minimum of 300 micro seconds faster.
