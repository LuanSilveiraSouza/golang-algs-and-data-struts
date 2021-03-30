## Interpolation Search

Interpolation search uses binary search as a base point, but with a improvement in the middle element choose. Instead of just pick the middle based in the actual lowest and highest value that are being used in that specific round, it do a calculation based in that values that returns a probe value that are most aligned with the input value.

```mid = low + ((x – A[low]) * (high – low) / (A[high] – A[low]));```

#### Complexity 
##### Worst and Medium Cases: O(n)
##### Best Cases: O(log(log n))