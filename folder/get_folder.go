package folder

import (
	"fmt"
	s "strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	var subTree *Folder = nil

	for _, folder := range f.folders {
		if folder.Name==name {
			subTree = &folder
			break
		}
	}

	if subTree == nil {
		fmt.Print("Error: Folder does not exist")
		return nil
	}

	if subTree.OrgId != orgID {
		fmt.Print("Error: Folder does not exist in the specified organization\n")
		return nil;
	}

	var returnList = make([]Folder, 0);

	for _, folder := range f.folders {
		if(s.HasPrefix(folder.Paths, subTree.Paths)) {
			returnList = append(returnList, folder)
		}
	}

	return returnList
}
