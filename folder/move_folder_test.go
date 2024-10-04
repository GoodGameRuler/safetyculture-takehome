package folder_test

import (
	"errors"
	"slices"
	s "strings"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name        string
		src         string
		dst         string
		folders     []folder.Folder
		want        []folder.Folder
		expectedErr error
	}{
		{
			name: "move_to_same_folder",
			src:  "folder1",
			dst:  "folder1",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "folder1"},
			},
			want:        nil,
			expectedErr: errors.New("Error: Cannot move a folder to itself"),
		},
		{
			name: "destination_does_not_exist",
			src:  "folder1",
			dst:  "folder2",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")), Paths: "folder1"},
			},
			want:        nil,
			expectedErr: errors.New("Error: Destination folder does not exist"),
		},
		{
			name: "source_does_not_exist",
			src:  "folder2",
			dst:  "folder1",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("33333333-3333-3333-3333-333333333333")), Paths: "folder1"},
			},
			want:        nil,
			expectedErr: errors.New("Error: Source folder does not exist"),
		},
		{
			name: "move_between_organisations",
			src:  "folder1",
			dst:  "folder2",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("44444444-4444-4444-4444-444444444444")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("55555555-5555-5555-5555-555555555555")), Paths: "folder2"},
			},
			want:        nil,
			expectedErr: errors.New("Error: Cannot move a folder to a different organization"),
		},
		{
			name: "move_to_child_folder",
			src:  "folder1",
			dst:  "folder2",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("66666666-6666-6666-6666-666666666666")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("66666666-6666-6666-6666-666666666666")), Paths: "folder1.folder2"},
			},
			want:        nil,
			expectedErr: errors.New("Error: Cannot move a folder to a child of itself"),
		},
		{
			name: "successful_move",
			src:  "folder1",
			dst:  "folder2",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("77777777-7777-7777-7777-777777777777")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("77777777-7777-7777-7777-777777777777")), Paths: "folder2"},
			},
			want: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("77777777-7777-7777-7777-777777777777")), Paths: "folder2.folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("77777777-7777-7777-7777-777777777777")), Paths: "folder2"},
			},
			expectedErr: nil,
		},
		{
			name: "move_nested_folders",
			src:  "folder1",
			dst:  "folder2",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("88888888-8888-8888-8888-888888888888")), Paths: "folder1"},
				{Name: "nestedfolder", OrgId: uuid.Must(uuid.FromString("88888888-8888-8888-8888-888888888888")), Paths: "folder1.nestedfolder"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("88888888-8888-8888-8888-888888888888")), Paths: "folder2"},
			},
			want: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("88888888-8888-8888-8888-888888888888")), Paths: "folder2.folder1"},
				{Name: "nestedfolder", OrgId: uuid.Must(uuid.FromString("88888888-8888-8888-8888-888888888888")), Paths: "folder2.folder1.nestedfolder"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("88888888-8888-8888-8888-888888888888")), Paths: "folder2"},
			},
			expectedErr: nil,
		},
		{
			name: "move_nested_folders_no_commonality",
			src:  "folder1",
			dst:  "folder11",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder2"},
				{Name: "folder3", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder3"},
				{Name: "folder4", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder4"},
				{Name: "folder5", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5"},
				{Name: "folder6", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder6"},
				{Name: "folder7", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7"},
				{Name: "folder8", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder8"},
				{Name: "folder9", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder9"},
				{Name: "folder10", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10"},
				{Name: "folder11", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11"},
				{Name: "folder12", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder12"},
			},
			want: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder2"},
				{Name: "folder3", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder3"},
				{Name: "folder4", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder4"},
				{Name: "folder5", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder5"},
				{Name: "folder6", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder5.folder6"},
				{Name: "folder7", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder5.folder7"},
				{Name: "folder8", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder5.folder7.folder8"},
				{Name: "folder9", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11.folder1.folder5.folder7.folder9"},
				{Name: "folder10", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10"},
				{Name: "folder11", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11"},
				{Name: "folder12", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder12"},
			},
			expectedErr: nil,
		},
		{
			name: "move_nested_folders_with_some_commonality",
			src:  "folder6",
			dst:  "folder9",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder2"},
				{Name: "folder3", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder3"},
				{Name: "folder4", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder4"},
				{Name: "folder5", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5"},
				{Name: "folder6", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder6"},
				{Name: "folder7", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7"},
				{Name: "folder8", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder8"},
				{Name: "folder9", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder9"},
				{Name: "folder10", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10"},
				{Name: "folder11", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11"},
				{Name: "folder12", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder12"},
			},
			want: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder2"},
				{Name: "folder3", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder3"},
				{Name: "folder4", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder4"},
				{Name: "folder5", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5"},
				{Name: "folder6", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder9.folder6"},
				{Name: "folder7", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7"},
				{Name: "folder8", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder8"},
				{Name: "folder9", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder9"},
				{Name: "folder10", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10"},
				{Name: "folder11", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11"},
				{Name: "folder12", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder12"},
			},
			expectedErr: nil,
		},
		{
			name: "move_nested_folders_with_full_commonality",
			src:  "folder8",
			dst:  "folder9",
			folders: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder2"},
				{Name: "folder3", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder3"},
				{Name: "folder4", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder4"},
				{Name: "folder5", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5"},
				{Name: "folder6", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder6"},
				{Name: "folder7", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7"},
				{Name: "folder8", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder8"},
				{Name: "folder9", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder9"},
				{Name: "folder10", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10"},
				{Name: "folder11", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11"},
				{Name: "folder12", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder12"},
			},
			want: []folder.Folder{
				{Name: "folder1", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1"},
				{Name: "folder2", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder2"},
				{Name: "folder3", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder3"},
				{Name: "folder4", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder4"},
				{Name: "folder5", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5"},
				{Name: "folder6", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder6"},
				{Name: "folder7", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7"},
				{Name: "folder8", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder9.folder8"},
				{Name: "folder9", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder1.folder5.folder7.folder9"},
				{Name: "folder10", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10"},
				{Name: "folder11", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder10.folder11"},
				{Name: "folder12", OrgId: uuid.Must(uuid.FromString("99999999-9999-9999-9999-999999999999")), Paths: "folder12"},
			},
			expectedErr: nil,
		},
		{
			name: "move_sneaky_match",
			src:  "folder_prefix_wont_work_a",
			dst:  "folder_prefix_wont_work_a_1",
			folders: []folder.Folder{
				{Name: "folder_prefix_wont_work", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work"},
				{Name: "folder_prefix_wont_work_a", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a"},
				{Name: "folder_prefix_wont_work_a_1", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a_1"},
			},
			want: []folder.Folder{
				{Name: "folder_prefix_wont_work", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work"},
				{Name: "folder_prefix_wont_work_a", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a_1.folder_prefix_wont_work_a"},
				{Name: "folder_prefix_wont_work_a_1", OrgId: uuid.Must(uuid.FromString("20202021-2021-2021-2021-202020202021")), Paths: "folder_prefix_wont_work.folder_prefix_wont_work_a_1"},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got, err := f.MoveFolder(tt.src, tt.dst)

			if !errors.Is(err, tt.expectedErr) && (err == nil || tt.expectedErr == nil || err.Error() != tt.expectedErr.Error()) {
				t.Errorf("MoveFolder() error = %v, want %v", err, tt.expectedErr)
			}

			// Testing that slices are equal independent of order
			sortingFunc := func (a folder.Folder, b folder.Folder) int {
				return s.Compare(a.Name, b.Name)
			}

			slices.SortStableFunc(got, sortingFunc)
			slices.SortStableFunc(tt.want, sortingFunc)

			if !slices.Equal(got, tt.want) {
				t.Errorf("MoveFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
