## Quick Sort

<p align="center">
    <img src="../../.github/quicksort.gif">
</p>

Quick Sort is one of the most used sorting algorithms. It doesn't have the best performance on its worst case-scenarios, but its medium case-scenarios performs very well. It applies the Divide-And-Conquer strategy together a recursive methodology. 
Quick Sort consists in choosing a element of the array that will be the pivot (there have multiple ways to choose, the most common is picking the last element of the collection). With pivot in hands, make the left side of the pivot being only the elements lower than it and its right side only with elements higher than it. After that, the pivot will be in its correctly place in the final sorted array. The next steps is doing quick sorts in the two sides of the pivot and continue doing recursive/divisive sorts until every element of the array are sorted.

#### Complexity 
##### Worst Cases: O(n²)
##### Medium and Best Cases: O(n log n)