## Documentation
### Problem 2 : Boss Baby's Revenge
To implement this module, I have identified the following approaches:

1. Unit Testing: This approach involves testing the functionality through unit tests to ensure correctness.

to run :
```
go test ./baby/... -v // at root level(band-protocol-test)
```

`from my validation i assume that incorrect input are Bad boy ,it easy to do unit test`

### Time and Memory Complexity
Time Complexity: O(n), where n is the length of the input string. This is because each character in the string is processed once.
Memory Complexity: O(1). The memory usage for storing the counts in the map is constant, regardless of the input size, as it only stores counts for 'S' and 'R'.

### Unit Test Coverage
test coverage is 100.0%