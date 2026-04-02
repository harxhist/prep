
# Horizontal vs Vertical Scaling

Assume a SaaS.
Client send requests to Server. Server sends back response.

More clients = More requests.
Server has a limit to handle max_requests.

Solution is SCALABILITY:
1. Buy Bigger machine(Vertical Scaling) or
2. Buy More machines(Horizontal Scaling)

More request handling by throwing more money at the problem.

-------

| Horizontal                          | Vertical                     |
|------------------------------------|------------------------------|
| Load Balancing                     | Not Applicable               |
| Resilient                          | Single Point of Failure      |
| Network Calls (Remote Procedure Calls)(slow) | Inter Process Communication(fast) |
| Data Inconsistency                 | Consistent                   |
| Scales well as users increase      | Hardware                     |

-------

Both are used in real world. 
Is your system scalable, resilient & consistent?