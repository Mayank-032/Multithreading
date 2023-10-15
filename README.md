# Multithreading

MULTI-THREADING

HOW TO WRITE GOOD MULTI-THREADED PROGRAM

Problem Statement: -> We have to traverse 100,000,000 numbers and have to find

There are 3 basic approaches:

# 1.) Sequential Search:
-> One by one we will traverse each number upto 100,000,000 and for each number we will run another logic to know if it is prime or not.

-> This prime number checking will generally take O(sqrt(n)) n here is 100,000,000.

-> Well running the program and time taken to executed could depend on which machine or how many processes running concurrently with it etc., etc.

# RESULTS OF SEQUENTIAL APPROACH TO CHECK PRIME NUMBERS IN 100,000,000 NUMBERS<br />
checking till 100000000 found 5761455 prime numbers. Took 29.210291875s<br />


# 2.) Using Multi-Threading
-> Now suppose I say that I am concurrently running 10 workers which means per worked 10,000,000 numbers i.e. 10,000,000 will be evaluated concurrently.

-> Now this process can be divided in 2 ways

#	2.1) Unfair Approach to multi-threading:
		-> Let’s say we assign first 10,000,000 numbers to first thread, another 10,000,000 to second and another to 3rd.. and goes on.
		
		-> Now why above approach is unfair because, calculating primes for first 10,000,000 would be fast since the square root of first 10,000,000 is less but what about the thread which is calculating for 90,000,000 to 100,000,000 the square root is going to be way larger then for first thread and thus more load on it thus more time.
	
		-> Thus the scenario is like its not actual concurrency because when the work of first thread is finished its sitting idol now, no work done while other threads are still processing.

# RESULTS OF UNFAIR MULTI-THREADING APPROACH TO CHECK PRIME NUMBERS IN 100,000,000 NUMBERS <br />
thread 0 from [3 - 10000003] completed in 2.508178958s<br />
thread 1 from [10000003 - 20000003] completed in 3.905038208s<br />
thread 2 from [20000003 - 30000003] completed in 4.715970875s<br />
thread 3 from [30000003 - 40000003] completed in 5.351821042s<br />
thread 4 from [40000003 - 50000003] completed in 5.780914083s<br />
thread 5 from [50000003 - 60000003] completed in 6.041456625s<br />
thread 6 from [60000003 - 70000003] completed in 6.32560725s<br />
thread 7 from [70000003 - 80000003] completed in 6.6594625s<br />
thread 8 from [80000003 - 90000003] completed in 6.72908425s<br />
thread 9 from [90000003 - 100000000] completed in 6.99657725s<br />
checking till 100000000 found 5761455 prime numbers. Took 6.996670667s<br />

	
#	2.2) Fair Approach to multi-threading
		-> Now fair approach is, instead of threads sitting idol, they all work concurrently no matter if the number assigned to it is processed or not.
		
		-> When first processing finishes move on to different, then different and so on…

		-> So this approach is like, first assign threads to first 10 numbers and then as soon as the processing (checking if prime or not) for that number is finished move on 			     to the 11th number, so in this case each and every thread will be doing almost equal amount of work (might be a bit more less but that difference is so negligible).

# RESULTS OF FAIR(OPTIMAL) MULTI-THREADING APPROACH TO CHECK PRIME NUMBERS IN 100,000,000 NUMBERS<br />
thread 2 completed in 6.853569375s<br />
thread 7 completed in 6.845954583s<br />
thread 9 completed in 6.853380583s<br />
thread 3 completed in 6.845876875s<br />
thread 1 completed in 6.853332125s<br />
thread 8 completed in 6.843340375s<br />
thread 0 completed in 6.853395042s<br />
thread 6 completed in 6.832699208s<br />
thread 5 completed in 6.846038625s<br />
thread 4 completed in 6.843755208s<br />
checking till 100000000 found 5761455 prime numbers. Took 6.853754583s<br />
