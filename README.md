# Motivations
> What has been done?

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

Perhaps there is a solution the also optimises for pre-processing. I make the argument that the use case this taks this would apply to for a client, would be in receiving data using the get services more often than using move service. I also argue that in a real-time system we are not pre-processing data often. Hence, for this task I take this perhaps _opinionated_ optimisation for the getter operations.

I really enjoyed tackling this task in go. There definitely were a few issues and lessons that I came across, some more painful than others, but all Very informative. As a means to truly experience the Golang Progamming Experience I used the following resources.

- [Effective Go](https://go.dev/doc/effective_go) for styling and design
- [Go By Example](https://gobyexample.com) to learn the language

Makefiles seemed a common build tool for Go Programmers, and it was something I also adopted to create a mini testing program that load's a json file and serves as a command-line API to use the functions that our interface exports.

## Disclaimer and Changes
_I do want to quickly address any changes I made to the original code_

- I added loading function to `static.go`, as a means to load json data form the terminal, without modifying the pre-existing functions.
- I changed the interface functions to return both a slice `[]Folder`, and an err `error`, as I felt all the functions did need error handling.

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

Follows the same structure.

Although the Makefile adds the following:
- While E2E were not implemented I did create a program, and these can be implemented using simple in and out files in the `tests` folder
- The program is built using the `Makefile` the default `make` rule does this and runs tests as well.

## Methods
> I experimented with a total of three methods before coming to the conclusion. In chronological order I present my methods.

### Initial Brute force Solution

### Tree Based Solution

### The O(1) Solution

## Testing

## My Experience
