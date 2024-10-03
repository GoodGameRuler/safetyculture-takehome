package folder

import (
	"errors"
	s "strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	folders := make([]Folder, len(f.folders))
	copy(folders, f.folders)

	if name==dst {
		return nil, errors.New("Error: Cannot move a folder to itself")
	}

	var dest_folder *Folder = nil
	var src_folder *Folder = nil

	for _, folder := range folders {
		switch folder.Name {
		case name:
			src_folder = &folder
		case dst:
			dest_folder = &folder
		}

		if dest_folder!=nil && src_folder!=nil {
			break
		}
	}

	if dest_folder==nil {
		return nil, errors.New("Error: Destination folder does not exist")
	}

	if src_folder==nil {
		return nil, errors.New("Error: Source folder does not exist")
	}

	if src_folder.OrgId!=dest_folder.OrgId {
		return nil, errors.New("Error: Cannot move a folder to a different organization")
	}

	if len(dest_folder.Paths)>=len(src_folder.Paths) && src_folder.Paths==dest_folder.Paths[:len(dest_folder.Paths)] {
		return nil, errors.New("Error: Cannot move a folder to a child of itself")
	}

	source_base_dir := src_folder.Paths[:len(src_folder.Paths) - len(src_folder.Name) - 1]

	for i, folder := range folders {
		if(s.HasPrefix(folder.Paths, src_folder.Paths)) {
			folders[i].Paths = s.Replace(folder.Paths, source_base_dir, dest_folder.Paths, 1)
		}
	}

	return folders, nil
}
