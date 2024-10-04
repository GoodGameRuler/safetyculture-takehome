package folder

import (
	"errors"
	_ "fmt"
	s "strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) ([]Folder, error) {
	folders := f.folders

	res := []Folder{}
	found := false

	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
			found = true
		}
	}

	if !found {
		return nil, errors.New("Error: No such organisation\n")
	}

	return res, nil

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error){
	var subTree *Folder = nil

	for _, folder := range f.folders {
		if folder.Name==name {
			subTree = &folder
			break
		}
	}

	if subTree == nil {

		return nil, errors.New("Error: Folder does not exist\n")
	}

	if subTree.OrgId != orgID {
		return nil, errors.New("Error: Folder does not exist in the specified organization\n")
	}

	var returnList = make([]Folder, 0);

	for _, folder := range f.folders {
		if(s.HasPrefix(folder.Paths, subTree.Paths + ".")) {
			returnList = append(returnList, folder)
		}
	}

	return returnList, nil
}
