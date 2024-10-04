package folder

import (
	"errors"
	s "strings"

	"slices"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := make([]Folder, len(f.folders))
	copy(folders, f.folders)

	if name == dst {
		return nil, errors.New("Error: Cannot move a folder to itself")
	}

	var destFolder *Folder = nil
	var srcFolder *Folder = nil

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

	commonLength := 0
	for i := 0; i < min(len(srcParts), len(dstParts)); i++ {
		if srcParts[i] == dstParts[i] {
			commonLength++
		} else {
			break
		}
	}

	newBasePath := destFolder.Paths
	if commonLength < len(srcParts) {
		newBasePath += "." + s.Join(srcParts[commonLength:], ".")
	}

	for i, folder := range folders {
		if folder.Name == srcFolder.Name {
			folders[i].Paths = newBasePath
		} else if s.HasPrefix(folder.Paths, srcFolder.Paths+".") {
			folders[i].Paths = s.Replace(folder.Paths, srcFolder.Paths+".", newBasePath+".", 1)
		}
	}

	slices.SortStableFunc(folders, func(f1, f2 Folder) int {
		return s.Compare(f1.Paths, f2.Paths)
	})

	// Commented-out code for updating maps (size, index, orgMap) if persistent state is used.
	/*
		// Update folderMap with new index and size values
		f.folderMap[srcFolder.Name] = folderInfo{
			index: <new_index>,
			size:  <new_size>,
		}

		// Update orgMap to reflect changes
		f.orgMap[srcFolder.OrgId] = append(f.orgMap[srcFolder.OrgId], srcFolder)
	*/

	// Return the modified folders slice
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
