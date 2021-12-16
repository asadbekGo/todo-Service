package postgres

import (
	"reflect"
	"testing"
	"time"

	pb "github.com/asadbekGo/todo-service/genproto"
)

func TestTodoRepo_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Todo
		want    pb.Todo
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Todo{
				Assignee: "asadbek",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			want: pb.Todo{
				Assignee: "asadbek",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			wantErr: false,
		},
		{
			name: "successful",
			input: pb.Todo{
				Assignee: "muhammad",
				Title:    "API gateway",
				Summary:  "restfull ",
				Deadline: "2021-12-18T18:00:10Z",
				Status:   "active",
			},
			want: pb.Todo{
				Assignee: "muhammad",
				Title:    "API gateway",
				Summary:  "restfull ",
				Deadline: "2021-12-18T18:00:10Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			got.Id = 0
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoRepo_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    pb.Todo
		wantErr bool
	}{
		{
			name:  "successful",
			input: 1,
			want: pb.Todo{
				Id:       1,
				Assignee: "asadbek",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoRepo_List(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			page, limit int64
		}
		want    []*pb.Todo
		wantErr bool
	}{
		{
			name: "succesful",
			input: struct {
				page, limit int64
			}{
				page:  1,
				limit: 2,
			},
			want: []*pb.Todo{
				{
					Id:       1,
					Assignee: "asadbek",
					Title:    "todo service",
					Summary:  "rpc implement",
					Deadline: "2021-12-15T14:12:14Z",
					Status:   "active",
				},
				{
					Id:       2,
					Assignee: "muhammad",
					Title:    "API gateway",
					Summary:  "restfull ",
					Deadline: "2021-12-18T18:00:10Z",
					Status:   "active",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.List(tc.input.page, tc.input.limit)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.wantErr, err, count)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.want, got, count)
			}
		})
	}
}

func TestTodoRepo_Update(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Todo
		want    pb.Todo
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Todo{
				Id:       1,
				Assignee: "asadbek",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T00:00:00Z",
				Status:   "active",
			},
			want: pb.Todo{
				Assignee: "asadbek",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T00:00:00Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Update(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			got.Id = 0
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoRepo_Delete(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    error
		wantErr bool
	}{
		{
			name:    "successful",
			input:   2,
			want:    nil,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := pgRepo.Delete(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
		})
	}
}

func TestTodoRepo_ListOverdue(t *testing.T) {
	layoutISO := "2006-01-02"
	toTime, err := time.Parse(layoutISO, "2021-12-10")
	if err != nil {
		t.Fatal("failed to time parse", err)
	}
	tests := []struct {
		name  string
		input struct {
			time        time.Time
			page, limit int64
		}
		want    []*pb.Todo
		wantErr bool
	}{
		{
			name: "succesful",
			input: struct {
				time        time.Time
				page, limit int64
			}{
				time:  toTime,
				page:  1,
				limit: 2,
			},
			want: []*pb.Todo{
				{
					Id:       1,
					Assignee: "asadbek",
					Title:    "todo service",
					Summary:  "rpc implement",
					Deadline: "2021-12-15T14:12:14Z",
					Status:   "active",
				},
				{
					Id:       2,
					Assignee: "muhammad",
					Title:    "API gateway",
					Summary:  "restfull ",
					Deadline: "2021-12-18T18:00:10Z",
					Status:   "active",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.ListOverdue(tc.input.time, tc.input.page, tc.input.limit)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.wantErr, err, count)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.want, got, count)
			}
		})
	}
}
