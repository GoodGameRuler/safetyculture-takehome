package folder

import "github.com/gofrs/uuid"

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

// {} Hash Map from (OrgID, folderPath) to folderNode
// Tree implmented using lists as children
	// Parent is a Node
	// Parent Path (maybe trivial with acess to parent node)
	// Children is a list
	// Reference to the folder in the folders slice

// Still Have folders for reference


// GetAllChildFolders given a path returns multiple folder objects
// Implemented as a HashMap lookup to the tree
// Followed by a simple DFS on the tree
// Keep a track of all folder references visited and return them as another slice

// Error Cases
	// Invalid Folder, or Org (easy with hashmap lookup)

// MoveFolder given two paths moves folders
// Two hashmap lookups, as simple as removing a node from the parent object and adding it to another
// As returning the entire slice is required
// Conduct a DFS on all child nodes of the source updating
	// the reference to the slice with the new path
// This is trivialwith the parent nodes path


// Error Cases
	// Moving parent folder to child folder
	// Moving folder to self
	// Same subfolder already present


type folderNode struct {
	children []*folderNode
	parent *folderNode
	path string
	name string
}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data

	// example: feel free to change the data structure, if slice is not what you want
	MapPathToNode map[string]*folderNode
	treeHead folderNode
	folders []Folder
}

// func (d *driver) addDevice() {
// }
//
// func (d *driver) initStructs() {
// 	d.treeHead = folderNode{
// 		// children: ma,
// 		parent: nil,
// 		path: "root",
// 		name: "root",
// 	}
//
// 	for folder := range(d.folders) {
// 		d.MapPathToNode[folder.Name] = &folderNode{}
// 	}
// }

func NewDriver(folders []Folder) IDriver {
	d := driver{
		folders: folders,
		// MapPathToNode: make(map[string]*folderNode, 0),

	}

	// d.initStructs()

	return &d
}
