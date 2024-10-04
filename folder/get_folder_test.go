package folder_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_invalid_foldername(t *testing.T) {
	data := folder.GenerateData()
	d := folder.NewDriver(data)


	res, err := d.GetFoldersByOrgID(uuid.FromStringOrNil("invalidID1234@abcd"))
	assert.Equal(t, res, []folder.Folder(nil))
	assert.NotNil(t, err)
}

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name        string
		orgID       uuid.UUID
		folders     []folder.Folder
		want        []folder.Folder
		expectedErr error
	}{
		{
			name:        "empty_folders",
			orgID:       uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"),
			folders:     []folder.Folder{},
			want:        []folder.Folder{},
			expectedErr: errors.New("Error: No such organisation\n"),
		},
		{
			name: "single_folder_match",
			orgID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"),
			folders: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"},
			},
			want:        []folder.Folder{{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"}},
			expectedErr: nil,
		},
		{
			name: "single_folder_no_match",
			orgID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000def"),
			folders: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"},
			},
			want:        []folder.Folder{},
			expectedErr: errors.New("Error: No such organisation\n"),
		},
		{
			name: "multiple_folders_some_match",
			orgID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"),
			folders: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000def"), Paths: "Folder2"},
				{Name: "Folder3", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder3"},
			},
			want: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"},
				{Name: "Folder3", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder3"},
			},
			expectedErr: nil,
		},
		{
			name: "multiple_folders_all_match",
			orgID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"),
			folders: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder2"},
			},
			want: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder2"},
			},
			expectedErr: nil,
		},
		{
			name:  "no_folders_with_same_org",
			orgID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000ghi"),
			folders: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000abc"), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000def"), Paths: "Folder2"},
			},
			want:        []folder.Folder{},
			expectedErr: errors.New("Error: No such organisation\n"),
		},
		// Sneaky Match Test Case
		{
			name: "sneaky_match",
			orgID: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")),
			folders: []folder.Folder{
				{Name: "folder_prefix_wont_work", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work"},
				{Name: "folder_prefix_wont_work_a", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a"},
				{Name: "folder_prefix_wont_work_a_1", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a_1"},
			},
			want: []folder.Folder{
				{Name: "folder_prefix_wont_work", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work"},
				{Name: "folder_prefix_wont_work_a", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a"},
				{Name: "folder_prefix_wont_work_a_1", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a_1"},
			},
			expectedErr: nil,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.GetFoldersByOrgID(tt.orgID)

			t.Log("Test: ", tt.name, "\n")
			// Compare the error messages
			if !errors.Is(err, tt.expectedErr) && (err == nil || tt.expectedErr == nil || err.Error() != tt.expectedErr.Error()) {
				t.Errorf("GetFoldersByOrgID() error = %v, want %v", err, tt.expectedErr)
			}

			// Use reflect.DeepEqual to compare slices
			if !slices.Equal(got, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want %v", got, tt.want)
			}
		})
	}
}
