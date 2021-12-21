package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/asadbekGo/todo-service/genproto"
)

func TestTodoService_Create(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Todo
		want  pb.Todo
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
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Create(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create todo", err)
			}
			got.Id = ""
			tc.want.CreatedAt = got.CreatedAt
			tc.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoService_Get(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  pb.Todo
	}{
		{
			name:  "successful",
			input: "0a0a2ca4-36ec-455f-97f5-8d2501af9015",
			want: pb.Todo{
				Assignee: "asadbek",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Get(context.Background(), &pb.ByIdReq{Id: tc.input})
			if err != nil {
				t.Error("failed to get todo", err)
			}
			got.Id = ""
			tc.want.CreatedAt = got.CreatedAt
			tc.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

/*
func TestTodoService_List(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			page, limit int64
		}
		want []*pb.Todo
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
					Assignee: "asadbek",
					Title:    "todo service",
					Summary:  "rpc implement",
					Deadline: "2021-12-15T14:12:14Z",
					Status:   "active",
				},
				{
					Assignee: "muhammad",
					Title:    "API gateway",
					Summary:  "restfull ",
					Deadline: "2021-12-18T18:00:10Z",
					Status:   "active",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			listReq := &pb.ListReq{
				Limit: tc.input.limit,
				Page:  tc.input.page,
			}
			got, err := client.List(context.Background(), listReq)
			if err != nil {
				t.Error("failed to list todo", err)
			}
			got.Todos

		})
	}
}
*/

// func TestTodoRepo_Update(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   pb.Todo
// 		want    pb.Todo
// 		wantErr bool
// 	}{
// 		{
// 			name: "successful",
// 			input: pb.Todo{
// 				Id:       "908b32e7-160f-4e6c-be3c-b1637a240b96",
// 				Assignee: "asadbek",
// 				Title:    "todo service",
// 				Summary:  "rpc implement",
// 				Deadline: "2021-12-15T00:00:00Z",
// 				Status:   "active",
// 			},
// 			want: pb.Todo{
// 				Assignee: "asadbek",
// 				Title:    "todo service",
// 				Summary:  "rpc implement",
// 				Deadline: "2021-12-15T00:00:00Z",
// 				Status:   "active",
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := pgRepo.Update(tc.input)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
// 			}
// 			got.Id = ""
// 			if !reflect.DeepEqual(tc.want, got) {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
// 			}
// 		})
// 	}
// }

// func TestTodoRepo_Delete(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   string
// 		want    error
// 		wantErr bool
// 	}{
// 		{
// 			name:    "successful",
// 			input:   "acfa43a9-1166-4d88-a0e4-96490a77b8b8",
// 			want:    nil,
// 			wantErr: false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := pgRepo.Delete(tc.input)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
// 			}
// 		})
// 	}
// }

// func TestTodoRepo_ListOverdue(t *testing.T) {
// 	layoutISO := "2006-01-02"
// 	toTime, err := time.Parse(layoutISO, "2021-12-10")
// 	if err != nil {
// 		t.Fatal("failed to time parse", err)
// 	}
// 	tests := []struct {
// 		name  string
// 		input struct {
// 			time        time.Time
// 			page, limit int64
// 		}
// 		want    []*pb.Todo
// 		wantErr bool
// 	}{
// 		{
// 			name: "succesful",
// 			input: struct {
// 				time        time.Time
// 				page, limit int64
// 			}{
// 				time:  toTime,
// 				page:  1,
// 				limit: 2,
// 			},
// 			want: []*pb.Todo{
// 				{
// 					Id:       "908b32e7-160f-4e6c-be3c-b1637a240b96",
// 					Assignee: "asadbek",
// 					Title:    "todo service",
// 					Summary:  "rpc implement",
// 					Deadline: "2021-12-15T14:12:14Z",
// 					Status:   "active",
// 				},
// 				{
// 					Id:       "acfa43a9-1166-4d88-a0e4-96490a77b8b8",
// 					Assignee: "muhammad",
// 					Title:    "API gateway",
// 					Summary:  "restfull ",
// 					Deadline: "2021-12-18T18:00:10Z",
// 					Status:   "active",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, count, err := pgRepo.ListOverdue(tc.input.time, tc.input.page, tc.input.limit)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.wantErr, err, count)
// 			}
// 			if !reflect.DeepEqual(tc.want, got) {
// 				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.want, got, count)
// 			}
// 		})
// 	}
// }
