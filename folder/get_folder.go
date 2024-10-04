package folder

import (
	"errors"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (d *driver) GetFoldersByOrgID(orgID uuid.UUID) ([]Folder, error) {
	// Retrieve all folders for the specified orgID
	folders, exists := d.orgMap[orgID]
	if !exists {
		return nil, errors.New("Error: No such organisation")
	}
	return folders, nil
}

func (d *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Check if the folder exists in the folder map
	folderInfo, exists := d.folderMap[name]
	if !exists {
		return nil, errors.New("Error: Folder does not exist")
	}

	_, orgExists := d.orgMap[orgID]
	if !orgExists {
		return nil, errors.New("Error: No such organisation")
	}
	// Ensure the folder belongs to the correct organisation
	if d.folders[folderInfo.index].OrgId != orgID {
		return nil, errors.New("Error: Folder does not exist in the specified organisation")
	}

	// Extract all child folders based on the calculated size from folderInfo
	startIndex := folderInfo.index
	endIndex := folderInfo.index + folderInfo.size

	return d.folders[startIndex:endIndex], nil
}
