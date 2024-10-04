package folder

import (
	"errors"
	s "strings"

	"slices"
)


// The same move folder function from the initial solution as state does not persist
// O(n^1)
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := make([]Folder, len(f.folders))
	copy(folders, f.folders)

	if name == dst {
		return nil, errors.New("Error: Cannot move a folder to itself")
	}

	var destFolder *Folder = nil
	var srcFolder *Folder = nil

	// Find our source and destination folders
	for i := range folders {
		switch folders[i].Name {
		case name:
			srcFolder = &folders[i]
		case dst:
			destFolder = &folders[i]
		}

		if destFolder != nil && srcFolder != nil {
			break
		}
	}

	if destFolder == nil {
		return nil, errors.New("Error: Destination folder does not exist")
	}

	if srcFolder == nil {
		return nil, errors.New("Error: Source folder does not exist")
	}

	if srcFolder.OrgId != destFolder.OrgId {
		return nil, errors.New("Error: Cannot move a folder to a different organization")
	}

	if s.HasPrefix(destFolder.Paths, srcFolder.Paths+".") {
		return nil, errors.New("Error: Cannot move a folder to a child of itself")
	}

	srcParts := s.Split(srcFolder.Paths, ".")
	dstParts := s.Split(destFolder.Paths, ".")

	// Find the longest common path
	commonLength := 0
	for i := 0; i < min(len(srcParts), len(dstParts)); i++ {
		if srcParts[i] == dstParts[i] {
			commonLength++
		} else {
			break
		}
	}

	newBasePath := destFolder.Paths
	srcPath := srcFolder.Paths

	// Prevent undefined behaviour but does open program to DOS Attacks if event
	// Never should be case - suggests that src is child or is dest but handled above
	if commonLength >= len(srcParts) {
		panic("Common length should never be greater than or equal to ")
	}

	newBasePath += "." + s.Join(srcParts[commonLength:], ".")

	// Update path
	for i, folder := range folders {
		if folder.Name == srcFolder.Name {
			folders[i].Paths = newBasePath
		} else if s.HasPrefix(folder.Paths, srcPath+".") {
			folders[i].Paths = s.Replace(folder.Paths, srcPath+".", newBasePath+".", 1)
		}
	}

	slices.SortStableFunc(folders, pathOrderComparator)

	// Commented-out logic for updating maps (size, index, orgMap) for the cae state is persistant
	// Essentially recreating the loop as in newDriver
	/*
		// Initialize a new folderMap and orgMap
		newFolderMap := make(map[string]folderInfo)
		newOrgMap := make(map[uuid.UUID][]Folder)

		// Loop through all folders to populate the new maps
		for index, folder := range folders {
			// Update the folderMap with the new index and size values
			newFolderMap[folder.Name] = folderInfo{
				index: index,
				size:  calculateSize(folder), // Function to calculate size of the folder
			}

			// Update the orgMap to reflect the changes
			newOrgMap[folder.OrgId] = append(newOrgMap[folder.OrgId], folder)
		}

		// Assign the new maps to the driver's folderMap and orgMap
		f.folderMap = newFolderMap
		f.orgMap = newOrgMap
	*/

	return folders, nil
}

// Utility function to get the minimum of two integers
// Math func exists only for floats
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
