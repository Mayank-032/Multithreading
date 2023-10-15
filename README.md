# Multithreading

RESULTS OF SEQUENTIAL APPROACH TO CHECK PRIME NUMBERS IN 100,000,000 NUMBERS
checking till 100000000 found 5761455 prime numbers. Took 29.210291875s



RESULTS OF UNFAIR MULTI-THREADING APPROACH TO CHECK PRIME NUMBERS IN 100,000,000 NUMBERS
thread 0 from [3 - 10000003] completed in 2.508178958s
thread 1 from [10000003 - 20000003] completed in 3.905038208s
thread 2 from [20000003 - 30000003] completed in 4.715970875s
thread 3 from [30000003 - 40000003] completed in 5.351821042s
thread 4 from [40000003 - 50000003] completed in 5.780914083s
thread 5 from [50000003 - 60000003] completed in 6.041456625s
thread 6 from [60000003 - 70000003] completed in 6.32560725s
thread 7 from [70000003 - 80000003] completed in 6.6594625s
thread 8 from [80000003 - 90000003] completed in 6.72908425s
thread 9 from [90000003 - 100000000] completed in 6.99657725s
checking till 100000000 found 5761455 prime numbers. Took 6.996670667s



RESULTS OF FAIR(OPTIMAL) MULTI-THREADING APPROACH TO CHECK PRIME NUMBERS IN 100,000,000 NUMBERS
thread 2 completed in 6.853569375s
thread 7 completed in 6.845954583s
thread 9 completed in 6.853380583s
thread 3 completed in 6.845876875s
thread 1 completed in 6.853332125s
thread 8 completed in 6.843340375s
thread 0 completed in 6.853395042s
thread 6 completed in 6.832699208s
thread 5 completed in 6.846038625s
thread 4 completed in 6.843755208s
checking till 100000000 found 5761455 prime numbers. Took 6.853754583s