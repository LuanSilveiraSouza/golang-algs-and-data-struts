## Binary Search Tree

Binary search trees (or BSTs) probably are the most used type of tree. It consists in a tree-like structure that has a root node, and each node only points to two another nodes. The special condition is that the left element always has to be lower than its parent, and the same to the right always be greater than its parent, composing a ordened tree that has some operations that runs with great efficiency.

#### Complexity 

| Algorithm | Average Case | Worst Case |
|-----------|--------------|------------|
|  Search   |   O(log n)   |    O(n)    |
|  Insert   |   O(log n)   |    O(n)    |
|  Delete   |   O(log n)   |    O(n)    |