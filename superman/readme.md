## Documentation
### Problem 2 : Superman's Chicken Rescue
To implement this module, I have identified the following approaches:

1. Unit Testing: This approach involves testing the functionality through unit tests to ensure correctness.

to run :
```
go test ./superman/... -v // at root level(band-protocol-test)
```

Approach to Solve the Problem: two pointers

- Since the positions are sorted, you can use a sliding window to efficiently count the number of chickens within any segment [p, p+k).
- Initialize two pointers, start and end, to represent the current window.
- Move the end pointer to expand the window until the difference between positions[end] and positions[start] is less than k.
- Move the start pointer to shrink the window if the difference is >= k.
- Keep track of the maximum number of chickens within any window.

### Unit Test Coverage
test coverage is 100.0%


### Time and Memory Complexity 
Time Complexity: O(n), loop iterates over each position from 0 to n-1 (where n is the number of positions). Thus, the outer loop runs n times, which is O(n).

Memory Complexity: O(1),The function uses a few variables (start, n, k, maxCarriedChicken, end). These require constant space, O(1).
