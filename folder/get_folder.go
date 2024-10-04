package folder

import (
	"errors"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

// O(1) solution that returns slice based on our map
func (d *driver) GetFoldersByOrgID(orgID uuid.UUID) ([]Folder, error) {
	folders, exists := d.orgMap[orgID]
	if !exists {
		return nil, errors.New("Error: No such organisation")
	}
	return folders, nil
}

// O(1) solution that returns slice based on the index and size of a subtree in the sorted list
func (d *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folderInfo, exists := d.folderMap[name]
	if !exists {
		return nil, errors.New("Error: Folder does not exist")
	}

	_, orgExists := d.orgMap[orgID]
	if !orgExists {
		return nil, errors.New("Error: No such organisation")
	}

	if d.folders[folderInfo.index].OrgId != orgID {
		return nil, errors.New("Error: Folder does not exist in the specified organisation")
	}

	startIndex := folderInfo.index
	endIndex := folderInfo.index + folderInfo.size

	return d.folders[startIndex:endIndex], nil
}
