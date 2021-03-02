## Binary Search

<p align="center">
    <img src="../.github/binarysearch.gif">
</p>

Binary Search is a simple but powerful searching algorithm that can be implemented with recursion or common loops. Its simplicity enables anyone to fast understand it and start to use. 
Binary Search uses the premise of Divide and Conquer, splitting the input list in middle in each cycle. It checks if the requested element is equal to the middle of that splitted array, if not, it continues to split (to left or right of the middle) until reaches the element or don't find it, returning -1 as a pattern.

#### Complexity 
##### Worst and Medium Cases: O(logn)
##### Best Cases: O(1)