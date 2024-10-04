package folder

import (
	s "strings"

	"github.com/gofrs/uuid"
)

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) ([]Folder, error)
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error)

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct {
	folders     []Folder
	folderMap   map[string]*FolderNode
	rootFolders []*FolderNode
}

type FolderNode struct {
	Folder   Folder
	Parent   *FolderNode
	Children []*FolderNode
}

func (d *driver) buildFolderTree() {
	for _, folder := range d.folders {
		parts := s.Split(folder.Paths, ".")
		currentNode := d.getOrCreateNode(parts, folder)

		d.folderMap[folder.Name] = currentNode

		// Add to root folders if number of parts is 1
		if len(parts) == 1 {
			d.rootFolders = append(d.rootFolders, currentNode)
		}
	}
}

func (d *driver) getOrCreateNode(parts []string, folder Folder) *FolderNode {
	var currentNode *FolderNode

	for _, part := range parts {
		if node, exists := d.folderMap[part]; exists {
			currentNode = node
		} else {
			newNode := &FolderNode{Folder: folder}

			if currentNode != nil {
				newNode.Parent = currentNode
				currentNode.Children = append(currentNode.Children, newNode)
			}
			currentNode = newNode


			d.folderMap[part] = newNode
		}
	}

	return currentNode
}

func NewDriver(folders []Folder) IDriver {
	d := &driver{
		folders:   folders,
		folderMap: make(map[string]*FolderNode),
	}

	d.buildFolderTree()
	return d
}
