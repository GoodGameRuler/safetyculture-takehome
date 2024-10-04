package folder

import (
	"errors"
	_ "errors"
	_ "fmt"

	s "strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}


func (d *driver) GetFoldersByOrgID(orgID uuid.UUID) ([]Folder, error) {
	res := []Folder{}
	for _, f := range d.folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("Error: No such organisation\n")
	}
	return res, nil
}

func (d *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folderNode, exists := d.folderMap[name]

	if !exists {
		return nil, errors.New("Error: Folder does not exist\n")
	}

	if folderNode.Folder.OrgId != orgID {
		return nil, errors.New("Error: Folder does not exist in the specified organization\n")
	}
	var result []Folder
	for _, f := range d.folders {
		if f.OrgId == orgID && s.HasPrefix(f.Paths, folderNode.Folder.Paths+".") {
			result = append(result, f)
		}
	}

	return result, nil
}
