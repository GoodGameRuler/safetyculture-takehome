package folder

import (
	"errors"
	s "strings"
)

func (d *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	srcNode, srcExists := d.folderMap[name]
	if !srcExists {
		return nil, errors.New("Error: Source folder does not exist")
	}
	dstNode, dstExists := d.folderMap[dst]
	if !dstExists {
		return nil, errors.New("Error: Destination folder does not exist")
	}

	if srcNode.Folder.Paths == dstNode.Folder.Paths {
		return nil, errors.New("Error: Cannot move a folder to itself")
	}

	if s.HasPrefix(dstNode.Folder.Paths, srcNode.Folder.Paths+".") {
		return nil, errors.New("Error: Cannot move a folder to a child of itself")
	}

	if srcNode.Folder.OrgId != dstNode.Folder.OrgId {
		return nil, errors.New("Error: Cannot move a folder to a different organization")
	}

	// Use srcNode.Parent to check for the parent folder.
	if srcNode.Parent != nil {
		parentNode := srcNode.Parent
		for i, child := range parentNode.Children {
			if child.Folder.Name == srcNode.Folder.Name {
				parentNode.Children = append(parentNode.Children[:i], parentNode.Children[i+1:]...)
				break
			}
		}

	// Check if the folder is in the root folders
	} else {
		for i, rootNode := range d.rootFolders {
			if rootNode.Folder.Name == srcNode.Folder.Name {
				d.rootFolders = append(d.rootFolders[:i], d.rootFolders[i+1:]...)
				break
			}
		}
	}


	newPath := s.Join([]string{dstNode.Folder.Paths, srcNode.Folder.Name}, ".")
	srcNode.Folder.Paths = newPath

	var updatePaths func(node *FolderNode, basePath string)
	updatePaths = func(node *FolderNode, basePath string) {
		for _, child := range node.Children {
			child.Folder.Paths = s.Join([]string{basePath, child.Folder.Name}, ".")
			updatePaths(child, child.Folder.Paths)
		}
	}
	updatePaths(srcNode, newPath)

	dstNode.Children = append(dstNode.Children, srcNode)

	var allFolders []Folder
	var collectFolders func(node *FolderNode)

	collectFolders = func(node *FolderNode) {
		allFolders = append(allFolders, node.Folder)
		for _, child := range node.Children {
			collectFolders(child)
		}
	}

	for _, rootNode := range d.rootFolders {
		collectFolders(rootNode)
	}

	return allFolders, nil
}
