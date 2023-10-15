# Multithreading

MULTI-THREADING

HOW TO WRITE GOOD MULTI-THREADED PROGRAM

Problem Statement: -> We have to traverse 100,000,000 numbers and have to find

There are 3 basic approaches:

# 1.) Sequential Search:
-> One by one we will traverse each number up to 100,000,000 and for each number we will run another logic to know if it is prime or not.

-> This prime number checking will generally take O(sqrt(n)) n here is 100,000,000.

-> Well running the program and the time taken to execute could depend on which machine or how many processes run concurrently with it etc., etc.

# RESULTS OF SEQUENTIAL APPROACH TO CHECK PRIME NUMBERS IN 100,000,000 NUMBERS<br />
checking till 100000000 found 5761455 prime numbers. Took 29.210291875s<br />


# 2.) Using Multi-Threading
-> Now suppose I say that I am concurrently running 10 workers which means per worked 10,000,000 numbers i.e. 10,000,000 will be evaluated concurrently.

-> Now this process can be divided into 2 ways

#	2.1) Unfair Approach to Multi-threading:
		-> Let’s say we assign the first 10,000,000 numbers to the first thread, another 10,000,000 to the second, and another to 3rd.. and goes on.
		
		-> Now why the above approach is unfair, calculating primes for the first 10,000,000 would be fast since the square root of the first 10,000,000 is less but what about the 
                   the thread which is calculated for 90,000,000 to 100,000,000 the square root is going to be way larger than for the first thread and thus more load on it thus more time.
	
		-> Thus the scenario is like it's not actual concurrency because when the work of the first thread is finished it's sitting idol now, no work is done while other threads 
                   are still processing.

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

	
#	2.2) Fair Approach to Multi-threading
		-> Now fair approach is, instead of threads sitting idol, they all work concurrently no matter if the number assigned to it is processed or not.
		
		-> When first processing finishes move on to different, then different, and so on…

		-> So this approach is like this, first, assign threads to the first 10 numbers, and then as soon as the processing (checking if prime or not) for that number is finished 
                   move on to the 11th number, in this case, each and every thread will be doing an almost equal amount of work (might be a bit less but that difference is so 
                   negligible).

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



# Creating Our Own TCP Server and handling multiple connections (Multi-Threading in TCP Connections)

**STEP-1: Listen to a PORT**<br />
-> Pick a PORT<br />
-> Start listening to the PORT<br />

**STEP-2: Start Accepting connections over the PORT**<br />
-> Invoke “ACCEPT” system call which will accept the connections on the server<br />
-> This will be a blocking call, which means the server will be blocked until it doesn’t accept the connection<br />

**STEP-3: Start Reading and writing the requests**<br />
-> Now once the client is connected, start reading the request with which the client connected, i.e. READ system call.<br />
-> Once the read is done, process the request and write something i.e. WRITE system call.<br />
-> After the connection write the response, and close the connection.<br />

**STEP-4: Start accepting multiple TCP connections over the PORT**<br />
-> We don’t let our server stop when the first client connects, instead, it should constantly listen for new connections<br />
-> Let’s say if the first connection request is processed and the response is returned, then our server’s port should not be closed.<br />

**STEP-5: Handling multiple connections simultaneously / concurrently instead of one by one**<br />
-> Now our objective is to handle the connection requests concurrently i.e. in such a way that another connection should not be blocked until the first connection request is resolved till 
   then no request will be handled.<br />

**IMPROVEMENTS:**
-> We cannot spin up an infinite number of threads, so in order to handle that we have to maintain a limitation on the number of threads.<br />
-> Thread pooling means we can use the thread and once the processing is done put it back in the pool<br />
-> Connection Timeout: Do not wait for a long time for the connection to send the request, if no request is made.. do the connection timeout<br />
