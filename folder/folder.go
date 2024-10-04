package folder

import (
	"slices"
	s "strings"

	"github.com/gofrs/uuid"
)

type IDriver interface {
	GetFoldersByOrgID(orgID uuid.UUID) []Folder

	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	MoveFolder(name string, dst string) ([]Folder, error)
}

// Custom tuple if you will
// Keeps track of the position of
// @node (of a non binary tree implementing using an array)
// @size of the subtree
type folderInfo struct {
	index int
	size int
}

type driver struct {
	folders      []Folder

	// Maps folder to the custom tuple
	folderMap    map[string]folderInfo

	// Maps org to slice of org's folders
	orgMap       map[uuid.UUID][]Folder
}


// Comparator function to sort folders in tree order
func pathOrderComparator(f1, f2 Folder) int {
	path1 := s.Split(f1.Paths, ".")
	path2 := s.Split(f2.Paths, ".")

	// If paths diverge sort alphabetically on divergent point
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

	// Sort slice on our path-based sorter
	slices.SortStableFunc(folders, pathOrderComparator)

	d := &driver{
		folders:   folders,
		folderMap: make(map[string]folderInfo),
		orgMap:    make(map[uuid.UUID][]Folder),
	}

	// Now on the sorted list build results
	for i, folder := range folders {
		d.folderMap[folder.Name] = folderInfo{index: i, size: 1}
		d.orgMap[folder.OrgId] = append(d.orgMap[folder.OrgId], folder)
	}

	return d
}
