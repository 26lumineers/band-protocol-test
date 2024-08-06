## Documentation
### Problem 3 : Transaction Broadcasting and Monitoring Client
To implement this module, I have identified the following approaches:

1. Unit Testing: This approach involves testing the functionality through unit tests to ensure correctness.

to run :
```
go test ./transaction/... -v // at root level(band-protocol-test)
```

However, I observed that handling the retry logic and managing API calls requires extensive code. To streamline this process, I picked Resty library. Resty offers robust capabilities for HTTP requests, including configurable retry mechanisms. This allows us to efficiently manage retries for API calls, handle custom retry intervals, and ensure that transactions with PENDING or FAILED statuses are reprocessed until they are completed.

custom configurations: 
- set retry : 3 times
- sets default wait time to sleep before retrying request. : 12 seconds.

`if timeout on unit test consider to decrease number of wait time`

### Unit Test Coverage
test coverage is 80.0%

### Time and Memory Complexity
Time Complexity: The time complexity is dominated by the network I/O operations and retry logic. While the core logic of the functions is 
O(1), practical execution times will be influenced by network delays and the number of retries.

Memory Complexity: The memory complexity of both functions is 
O(1) due to the constant amount of additional memory used by variables and Resty client configurations.

### Design Decisions
#### 1. Transaction Broadcasting
#### Function: TransactionBroadcastResty
```
+------------------------+
| Validate Payload       |
+------------------------+
           |
           v
+------------------------+
| Send Payload (POST)    |
+------------------------+
           |
           v
+------------------------+
| Receive TxHash Response|
+------------------------+

```
- Objective: To send a transaction payload to the broadcasting service and obtain a transaction hash.
- Implementation:
    - Validation: The payload is validated to ensure it meets the required criteria (non-empty symbol, positive price, and timestamp).
    - HTTP Request: Uses the resty library to perform a POST request with the payload. This library is chosen for its ease of use and built-in features like retry mechanisms.
    - Error Handling: Logs errors and returns them if the request fails.
#### Trade-offs:

Pros:
- Using resty abstracts away the lower-level details of making HTTP requests.
- Validation ensures that only valid data is processed, reducing the risk of invalid transactions.

Cons:
- Dependency on an external library might increase the complexity of the codebase.


#### 2. Transaction Status Checking
#### Function: CheckTransactionStatusResty
```
+------------------------+
| Validate TxHash        |
+------------------------+
           |
           v
+------------------------------------------+
| Concurrent execute Check Status (GET)    |
+------------------------------------------+
           |
           v
+------------------------+
| Receive TxStatus       |
+------------------------+

```

- Objective: To check the status of a transaction using its hash by querying the status endpoint.
- Implementation:
    - Retry Mechanism: Configured to retry the request up to RETRY_COUNT times with a wait time of SET_RETRY_WAIT_TIME between retries. This is essential for handling transient errors and improving reliability.
    - Retry Conditions: Custom retry logic based on the transaction status. The request is retried if the status is still PENDING or FAILED.
    - Error Handling: Logs errors and returns them if the request fails.

#### Trade-offs:

Pros:
- Retry mechanism helps in handling temporary network issues or server unavailability.
- Custom retry conditions provide flexibility in handling different transaction statuses.

Cons:
- Excessive retries might lead to delays in obtaining the final status.

### Scalability, Reliability, and Performance
Scalability:

- Scalability of HTTP Requests: Using resty allows for efficient management of concurrent HTTP requests, supporting horizontal scaling of the broadcasting and status checking operations.
- Retry Mechanism: Helps in managing transient failures, which is crucial for systems that need high availability and scalability.

Reliability:
- Validation: Ensures that only valid data is processed, reducing the risk of errors during the broadcasting and status checking phases.
- Retry Logic: Improves reliability by handling intermittent issues gracefully, avoiding immediate failures and improving overall fault tolerance.

Performance:
- Efficient HTTP Client: resty is optimized for performance, and its built-in features like retry conditions help in reducing the overhead of manual error handling and retries.
- Minimal Logging: Logs errors only when necessary, reducing performance impact while still providing useful debugging information.

### Assumptions
1. Network Reliability: The system assumes that network connectivity is generally reliable, but handles transient failures using retries.
2. Service Availability: Assumes that the broadcasting and status-checking services are generally available and responsive.

