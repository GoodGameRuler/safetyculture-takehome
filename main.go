package main

import (
	"fmt"
	"os"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

var HELP_STR = [...]string{
	"query-all-folders",
	"query-subfolders sub-folder",
	"query-folders-by-org organisation",
	"query-move-folder src dest",
}

func help() {
	print("HELP -- Commands\n")
	print("General Usage - ./folders source.json subcomand[...args]\n\nSubommands\n")
	for _, f := range HELP_STR {
		print("\t", f, "\n")
	}

}

func check_commands(query string, expected int, actual int) int {
	if actual != expected {
		fmt.Printf("Error: %s expected %d got %d arguments\n", query, expected, actual)
		return -1
	}

	return 0
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Error: Expecting a JSON file and a subcommand")
        help()
        os.Exit(1)
    }

    jsonFile := os.Args[1]
    folders, err := folder.LoadData(jsonFile)
    if err != nil {
        fmt.Printf("Error loading data: %v\n", err)
        os.Exit(1)
    }

    folderDriver := folder.NewDriver(folders)

    switch os.Args[2] {
    case "query-all-folders":
        if check_commands("query-all-folders", 0, len(os.Args)-3) < 0 {
			break
        }
		folder.PrettyPrint(folders)

    case "query-subfolders":
        if check_commands("query-subfolders", 1, len(os.Args)-3) < 0 {
            os.Exit(-1)
        }
        orgIDStr := os.Args[3]
        orgID := uuid.FromStringOrNil(orgIDStr)
        subFolders, err := folderDriver.GetAllChildFolders(orgID, os.Args[4])
        if err != nil {
            fmt.Printf("Error getting subfolders: %v\n", err)
			break
        }
        folder.PrettyPrint(subFolders)

    case "query-folders-by-org":
        if check_commands("query-folders-by-org", 1, len(os.Args)-3) < 0 {
			break
        }
        orgIDStr := os.Args[3]
        orgID := uuid.FromStringOrNil(orgIDStr)
        orgFolders, err := folderDriver.GetFoldersByOrgID(orgID)
        if err != nil {
            fmt.Printf("Error getting folders by orgID: %v\n", err)
			break
        }
        folder.PrettyPrint(orgFolders)

    case "query-move-folder":
        if check_commands("query-move-folder", 2, len(os.Args)-3) < 0 {
            os.Exit(-1)
        }
        src := os.Args[3]
        dest := os.Args[4]
        movedFolders, err := folderDriver.MoveFolder(src, dest)
        if err != nil {
            fmt.Printf("Error moving folder: %v\n", err)
            os.Exit(1)
        }
        folder.PrettyPrint(movedFolders)

    case "help":
        help()

    default:
        fmt.Println("Invalid Command")
        help()
    }
}
