## AVL Tree

<p align="center">
    <img src="../.github/avltree.gif">
</p>

AVL Trees are basicaly binary search trees that have height balancing. In this context, the balance of a tree means that each node has a height. This height is the difference of extension of the right and left subtrees. By default, AVL Trees implements self-balancing methods every time a new element is inserted or deleted and it breaks the balance of some node of the tree.
To balance the tree it is used rotation operations There aree 4 cases of disbalance possibles, and each one will require a specific movement in the tree.
The self-balance attribute makes AVL Trees operations acts faster than common BSTs in its worst cases.   

#### Complexity 

| Algorithm | Average Case | Worst Case |
|-----------|--------------|------------|
|  Search   |   O(log n)   |  O(log n)  |
|  Insert   |   O(log n)   |  O(log n)  |
|  Delete   |   O(log n)   |  O(log n)  |