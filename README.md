# Go LeetCode Solutions

This is a Go project where I solve problems from LeetCode using the Go programming language. The project is structured to include solutions for various LeetCode problems, organized in a clear and maintainable way.

## LeetCode Profile

You can find my LeetCode profile here: [etharrra](https://leetcode.com/u/etharrra/).

## Project Structure

The project is organized as follows:

### `solutions/`

This directory contains the Go files for each LeetCode problem I've solved. Each file corresponds to a specific problem and includes the solution code.

### `utils/`

This directory contains utility files that define common data structures like `ListNode` and `TreeNode`, which are frequently used in LeetCode problems.

### `main.go`

This is the entry point of the project. It can be used to test or demonstrate the solutions.

### `go.mod`

This file defines the Go module and its dependencies.

## How to Use

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/etharrra/go-leetcode.git
    cd go-leetcode
    ```

2.  **Run a specific solution:**
    You can run any of the solution files by calling the function name in main function. For example:

    ```go
    func main() {
    	s := "A man, a plan, a canal: Panama"
    	fmt.Println(solutions.IsPalindrome(s))
    }
    ```

    And the run the command

    ```bash
    go run main.go
    ```

3.  **Add new solutions:**
    To add a new solution, create a new Go file in the `solutions/` directory and implement the solution there. You can use the existing files as a reference.

4.  **Test your solutions:**
    You can write tests for your solutions using Go's testing framework. Create a `_test.go` file for each solution and write test cases.

## List of Solutions

Here is a list of the LeetCode problems solved in this project:

-   **Binary Operations:**

    -   `addBinary.go`: Adds two binary strings.
    -   `addBinarySlow.go`: A slower implementation of binary addition.
    -   `decimalToBinary.go`: Converts a decimal number to binary.

-   **Arrays and Strings:**

    -   `containsDuplicate.go`: Checks if an array contains duplicates.
    -   `containsNearbyDuplicate.go`: Checks for duplicates within a certain range.
    -   `fizzBuzz.go`: Implements the FizzBuzz problem.
    -   `intersect.go`: Finds the intersection of two arrays.
    -   `intersection.go`: Another implementation of array intersection.
    -   `isAnagram.go`: Checks if two strings are anagrams.
    -   `isPalindrome.go`: Checks if a string is a palindrome.
    -   `lengthOfLastWord.go`: Finds the length of the last word in a string.
    -   `majorityElement.go`: Finds the majority element in an array.
    -   `merge.go`: Merges two sorted arrays.
    -   `missingNumber.go`: Finds the missing number in an array.
    -   `modifyString.go`: Modifies a string based on certain rules.
    -   `plusOne.go`: Increments a number represented as an array.
    -   `productExceptSelf.go`: Computes the product of an array except for the current element.
    -   `remainderSorting.go`: Sorts an array based on remainders.
    -   `singleNumber.go`: Finds the single number in an array where every other number appears twice.
    -   `topKFrequent.go`: Finds the top K frequent elements in an array.

-   **Linked Lists:**

    -   `deleteDuplicates.go`: Removes duplicates from a sorted linked list.
    -   `getIntersectionNode.go`: Finds the intersection node of two linked lists.
    -   `hasCycle.go`: Checks if a linked list has a cycle.
    -   `nextUniqueNode.go`: Finds the next unique node in a linked list.
    -   `nodeCount.go`: Counts the number of nodes in a linked list.

-   **Trees:**

    -   `getDepth.go`: Computes the depth of a binary tree.
    -   `getRow.go`: Gets a specific row from Pascal's Triangle.
    -   `hasPathSum.go`: Checks if a binary tree has a path with a given sum.
    -   `inorderTraversal.go`: Performs inorder traversal of a binary tree.
    -   `inorderTraversalLoop.go`: Iterative inorder traversal.
    -   `isBalanced.go`: Checks if a binary tree is balanced.
    -   `isSameTree.go`: Checks if two binary trees are the same.
    -   `isSymmetric.go`: Checks if a binary tree is symmetric.
    -   `isSymmetricLoop.go`: Iterative check for symmetric binary trees.
    -   `levelorderTraversal.go`: Performs level-order traversal of a binary tree.
    -   `maxDepth.go`: Finds the maximum depth of a binary tree.
    -   `minDepth.go`: Finds the minimum depth of a binary tree.
    -   `postorderTraversal.go`: Performs postorder traversal of a binary tree.
    -   `postorderTraversalLoop.go`: Iterative postorder traversal.
    -   `preorderTraversal.go`: Performs preorder traversal of a binary tree.
    -   `preorderTraversalLoop.go`: Iterative preorder traversal.
    -   `sortedArrayToBST.go`: Converts a sorted array to a balanced binary search tree.

-   **Dynamic Programming:**

    -   `climbStairs.go`: Solves the climbing stairs problem.
    -   `generate.go`: Generates Pascal's Triangle.
    -   `longestConsecutive.go`: Finds the longest consecutive sequence in an array.
    -   `maxProfit.go`: Computes the maximum profit from stock prices.

-   **Miscellaneous:**
    -   `isValidSudoku.go`: Checks if a Sudoku board is valid.
    -   `mySqrt.go`: Computes the integer square root of a number.

## Contributing

If you'd like to contribute to this project, feel free to fork the repository and submit a pull request. Please ensure that your code follows the existing style and includes appropriate tests.

---

Happy coding! 🚀
