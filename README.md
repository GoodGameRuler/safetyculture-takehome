# SC Take Home Assessment README

## Introduction
This is Udit Samant's submission for the SafetyCulture Summer Internship 2024 Take Home Assessment.

The solution that I am submitting as the following properties

**Time Complexity**
Preprocessing: O(n^2)

`GetAllChildFolders`: O(1)

`GetFoldersByOrgID`: O(1)

`MoveFolder`: O(n^2)

**Space Complexity**
O(n) across

I argue that the main way a client would use this is by receiving data through 'get' services more often than using the 'move' service. In real-time systems, we don't usually pre-process data that much, so it made sense to focus on making get operations faster.

I really enjoyed tackling this task in go. There definitely were a few issues and lessons that I came across, some more painful than others, but all informative. As a means to truly experience the Golang Programming Experience I used the following resources.

- [Effective Go](https://go.dev/doc/effective_go) for styling and design
- [Go By Example](https://gobyexample.com) to learn the language

Makefiles seemed a common build tool for Go Programmers, and it was something I also adopted to create a mini testing program that loads a json file and serves as a command-line API to use the functions that our interface exports.

## Disclaimer and Changes
_Going over changes to the original codebase_

- Added a load data function to `static.go`, to load json data from the terminal, without modifying the pre-existing functions.
- Changed the interface functions to return both a slice `[]Folder`, and an error `error`, to ensure that we are handling edge cases where folder retrieval could fail properly.
- Moved the previous `REAMDE.md` to `LEGACY_README.md`. I wanted to use this as a means of communicating my thinking process and learning strategy.

## Structure
```
| go.mod
| README.md
| main.go
| Makefile
| folder
    | get_folder.go
    | get_folder_test.go
    | move_folder.go
    | static.go
    | sample.json
```

The overall structure remains the same.

However, the Makefile introduces the following:
- **E2E Tests:** While E2E were not implemented I did create a program, and these can be tested using simple in and out files in the `tests` folder.
- **Building:** The program is built using the `Makefile` the default `make` rule does this and runs tests as well.

## Methods
> I experimented with a total of three methods before coming to the conclusion. In chronological order I present my methods.

### Initial Brute force Solution
I started with a brute force approach, where every time a "find" operation was needed (to locate a particular folder in the hierarchy), I would loop through all folders and perform string matching. This meant that all operations had a time complexity of O(n^2), but there was no preprocessing involved.

Please see `udit-initial-sol`

### Tree Based Solution
For my second solution, I explored a node-based tree structure, enhanced by an auxiliary data structure that mapped folder names to their corresponding folder nodes. This setup improved the efficiency of locating source and destination nodes during operations, reducing the lookup time compared to the previous method. However, the overall operations still remained O(n^2) as we still traverse all elements and perform comparisons on the file paths, which is still an O(n) operation.

Please see `udit-tree-sol`

### The O(1) Solution

The final solution came by leveraging maps to dramatically optimise the process. Maps, with their O(1) lookup time, allowed us to encode the tree structure directly into the slice by sorting folders based on their paths. We used an auxiliary map to link each folder name to its corresponding index and the size of its subtree. This approach ensures that when we execute a "get" operation, we can return a subarray in constant time—since Go slices are essentially a pointer, length, and capacity, this becomes an O(1) operation ([Slice Expression](https://go.dev/ref/spec#Slice_expressions)). The move folder operation remains the same from the brute force approach and so remains O(n^2). There is also a similar pre=processing step that occurs when the data comes in. This is also O(n^2). However, the advantage of having constant time get requests irrespective of the size of the query outweighs the cost. A simple elegant solution.

Please see `main`

## Testing
Testing was a large part of developing this solution. I planned to test alongside developing the brute force solution, and extend this to any other solutions developed, as a means of Test Driven Development. This was the first time I had encountered Table Driven Tests, and while quite difficult to wrap my head around at first, I got well acquainted. In the end Table Driven Tests made it easier to structure tests across multiple methods and branches.

Particularly, I am thankful for my tests as they uncovered an issue that was raised with prefix matching for folders: `folder` was processed as a child of `folder11` even if that wasn't the case.

## My Experience
Following learning Zig, Go seemed quite like a distant cousin. The experience was quite enjoyable. I enjoyed Go's simple and easy-to-use declarative structure.

I am grateful to have a reason to learn go. I would love to follow through and spend some more time learning Go. Particularly I would love to learn about using Go to do following
- Creating an HTTP Server
- Using Channels
- Processes and Systems Development
- Benchmark Testing

Although, I did face the following issues
- I had some trouble navigating documentation. I found the Go by Example guide a much quicker way to navigate Go than the official documentation.
- Testing in Go was quite different to any other library I have used, and did take some time to get accustomed to.
