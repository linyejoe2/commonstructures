# CommonStructures ⚒️

**CommonStructures** is a Go library providing a variety of reusable data structures. The project aims to implement commonly used structures like stack, queue, and list using both array and linked list.

> NOTICE: All array is implement by go build-in struct `slice` in this project!

## 📦 Current Version: v0.0.5

## 🏗️ Available Data Structures

### 1. `rarraystack`
   - **Type**: Stack (LIFO)
   - **Implementation**: Array-based
   - **Description**: A stack that provides operations for pushing, popping, and peeking using an array.
   - **Usage**: `import "github.com/linyejoe2/commonstructures/rstack/rarraystack"`

### 2. `rlinkedliststack`
   - **Type**: Stack (LIFO)
   - **Implementation**: Linked list-based
   - **Description**: A stack that uses linked list nodes to implement the stack's functionality.
   - **Usage**: `import "github.com/linyejoe2/commonstructures/rstack/rlinkedliststack`

### 3. `rarrayqueue`
   - **Type**: Queue (FIFO)
   - **Implementation**: Array-based
   - **Description**: A queue that uses an array to manage enqueue and dequeue operations.
   - **Usage**: `import "github.com/linyejoe2/commonstructures/rqueue/rarrayqueue`


### 4. `rlinkedlistqueue`
   - **Type**: Queue (FIFO)
   - **Implementation**: Linked list-based
   - **Description**: A queue that operates using a linked list for efficient head/tail operations.
   - **Usage**: `import "github.com/linyejoe2/commonstructures/rqueue/rlinkedlistqueue`

### 5. `rlist`
   - **Type**: List
   - **Implementation**: Array-based and Linked list-based
   - **Description**: A basic list structure supporting multiple implementations, designed for easy element management.
   - **Usage**: `import "github.com/linyejoe2/commonstructures/rlist`

### 6. `rarray` (Utility Functions)
   - **Type**: Utility functions for arrays
   - **Description**: A set of helper functions to perform common operations on arrays.
   - **Usage**: `import "github.com/linyejoe2/commonstructures/rarray`

### 7. `rlinkedlist`
   - **Type**: linked lists
   - **Description**: A set of helper functions to simplify linked list operations, including node creation and traversal.
   - **Usage**: `import "github.com/linyejoe2/commonstructures/rlinkedlist`

<!-- --- -->

## 🚀 Getting Started

### Prerequisites

- Go 1.18+ installed on your machine.

### Installation

To use **CommonStructures** in your Go project, run:

```bash
go get github.com/linyejoe2/commonstructures@v0.0.5
```

## 💻 Example Usage

Here are a few examples showing how to use different structures from this package.

### Example: Using `rarraystack`

```go
package main

import (
    "fmt"
    "github.com/linyejoe2/commonstructures/rarraystack"
)

func main() {
    stack := rarraystack.NewRArrayStackte a stack of capacity 10
    stack.Push(42)
    stack.Push(99)
    fmt.Println(stack.Pop()) // Outputs: 99
    fmt.Println(stack.Peek()) // Outputs: 42
}
```

### Example: Using `rlinkedlistqueue`

```go
package main

import (
    "fmt"
    "github.com/linyejoe2/commonstructures/rlinkedlistqueue"
)

func main() {
    queue := rlinkedlistqueue.NewRLinkListQueue[int]()
    queue.Enqueue(10)
    queue.Enqueue(20)
    fmt.Println(queue.Dequeue()) // Outputs: 10
    fmt.Println(queue.Peek())    // Outputs: 20
}
```
