package folder

import (
	"slices"
	s "strings"

	"github.com/gofrs/uuid"
)

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

type folderInfo struct {
	index int
	size int
}

type driver struct {
	folders      []Folder
	folderMap    map[string]folderInfo
	orgMap       map[uuid.UUID][]Folder
}


// Comparator function to sort folders in tree order
func pathOrderComparator(f1, f2 Folder) int {
	path1 := s.Split(f1.Paths, ".")
	path2 := s.Split(f2.Paths, ".")

	// If paths diverge sort alphabetically
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] != path2[i] {
			if path1[i] < path2[i] {
				return -1
			}
			return 1
		}
	}

	// If paths are same till certain length return based on shorter length
	if len(path1) < len(path2) {
		return -1
	}
	if len(path1) > len(path2) {
		return 1
	}

	return 0
}
func NewDriver(folders []Folder) *driver {
	slices.SortStableFunc(folders, pathOrderComparator)

	d := &driver{
		folders:   folders,
		folderMap: make(map[string]folderInfo),
		orgMap:    make(map[uuid.UUID][]Folder),
	}

	for i, folder := range folders {
		d.folderMap[folder.Name] = folderInfo{index: i, size: 1}
		d.orgMap[folder.OrgId] = append(d.orgMap[folder.OrgId], folder)
	}

	return d
}
