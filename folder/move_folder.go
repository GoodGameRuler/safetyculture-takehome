package folder

import (
	"errors"
	s "strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	folders := make([]Folder, len(f.folders))
	copy(folders, f.folders)

	if name == dst {
		return nil, errors.New("Error: Cannot move a folder to itself")
	}

	var destFolder *Folder = nil
	var srcFolder *Folder = nil

	for _, folder := range folders {
		switch folder.Name {
		case name:
			srcFolder = &folder
		case dst:
			destFolder = &folder
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

	if s.HasPrefix(destFolder.Paths, srcFolder.Paths + ".") {
		return nil, errors.New("Error: Cannot move a folder to a child of itself")
	}

	srcParts := s.Split(srcFolder.Paths, ".")
	dstParts := s.Split(destFolder.Paths, ".")

	// Find the first differing index
	commonLength := 0
	for i := 0; i < min(len(srcParts), len(dstParts)); i++ {
		if srcParts[i] == dstParts[i] {
			commonLength++
		} else {
			break
		}
	}

	// Build the new base path for the source
	newBasePath := destFolder.Paths

	if commonLength >= len(srcParts) {
		panic("This should never be the case\n")
	}

	newBasePath += "." + s.Join(srcParts[commonLength:], ".")

	for i, folder := range folders {
		if folder.Name == srcFolder.Name {
			folders[i].Paths = newBasePath

		} else if s.HasPrefix(folder.Paths, srcFolder.Paths + ".") {
			folders[i].Paths = s.Replace(folder.Paths, srcFolder.Paths + ".", newBasePath + ".", 1)
		}
	}

	return folders, nil
}

// Utility function to get the minimum of two integers
// Math function takes floats instead
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
