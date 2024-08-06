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