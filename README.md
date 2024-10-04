# SC Take-home README

## Introduction
This Udit Samant's submission for the SafetyCulture Summer Internship 2024 Take Home Assessment.

The final solution that I am submitting as the following properties

**Time Complexity**
Prepossessing: O(n^2)

GetSubFolders: O(1)

GetByOrgID: O(1)

MoveFolders: O(n^2)

**Space Complexity**
O(n) across

I argue that the main way a client would use this is by receiving data through 'get' services more often than using the 'move' service. In real-time systems, we don't usually pre-process data that much, so it made sense to focus on making get operations faster.

I really enjoyed tackling this task in go. There definitely were a few issues and lessons that I came across, some more painful than others, but all informative. As a means to truly experience the Golang Programming Experience I used the following resources.

- [Effective Go](https://go.dev/doc/effective_go) for styling and design
- [Go By Example](https://gobyexample.com) to learn the language

Makefiles seemed a common build tool for Go Programmers, and it was something I also adopted to create a mini testing program that load's a json file and serves as a command-line API to use the functions that our interface exports.

## Disclaimer and Changes
_Going over changes to the original codebase_

- Added a load data function to `static.go`, to load json data from the terminal, without modifying the pre-existing functions.
- Changed the interface functions to return both a slice `[]Folder`, and an error `error`, as I felt all the functions did need proper handling of errors.
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

The overall

However, the Makefile introduces the following:
- **E2E Tests:** While E2E were not implemented I did create a program, and these can be tested using simple in and out files in the `tests` folder.
- **Building:** The program is built using the `Makefile` the default `make` rule does this and runs tests as well.

## Methods
> I experimented with a total of three methods before coming to the conclusion. In chronological order I present my methods.

### Initial Brute force Solution
I started with brute force approach. Every time a "find" operation (locating a particular folder in the hierarchy) was need I would loop through all folders and do string matching. This meant that all operations took O(n^2) time, but there was no pre-processing.

Please see `udit-initial-sol`

### Tree Based Solution
A second solution I explored was a node based Tree structure, with an auxiliary data-structured that mapped a folder name to folder node. This meant that some time was saved in finding the source and destination nodes during operations, but operations were still O(n^2) as the collection process ran through all elments with comparisons on the file path which is an O(n) operation.

Please see `udit-tree-sol`

### The O(1) Solution

The hint I got from this was maps. Obviously maps provide an amortised time-complexity, but they are simple and elegant as compared to a string prefix match. This time around I encode the tree structure in the slice itself by ordering by the path. I then use an auxiliary data-structure, to map folderName to the index and (subtree) size of each node. Following, the get operations simply return a subarray. Which is a constant timer operation as a slice is represented by a ptr, a length, and a capacity ([Slice Expression](https://go.dev/ref/spec#Slice_expressions)).

Please see `main`

## Testing
Testing was a large part of developing this solution. I planned to test alongside developing the brute force solution, and extend this to any other solutions developed, as a means of Test Driven Development. This was the first time I had encountered Table Driven Tests, and while quite difficult to wrap my head around at first, I got well acquainted.

## My Experience
Following learning Zig, Go seemed quite like a distant cousin. The experience was quite enjoyable. I enjoy Go's simple and easy to use declarative structure.

I am grateful to have a reason to learn go. I would love to follow through and spend some more time learning Go. Particularly I would love to learn about using go to do following
- Creating an HTTP Server
- Using Channels
- Processes and Systems Development

Although, I did face the following issues
- I had some trouble navigating documentation. I found the Go by Example guide a much quicker way to navigate Go than the documentation at first.
- Testing in Go was quite different to any other library I have used, and did take some time to get accustomed to.
