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
			input: "73718693-0a2b-41da-ba08-de7bf7940340",
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

func TestTodoService_List(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			page, limit int64
		}
		wants []*pb.Todo
	}{
		{
			name: "succesful",
			input: struct {
				page, limit int64
			}{
				page:  1,
				limit: 2,
			},
			wants: []*pb.Todo{
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
			for i, want := range tc.wants {
				got.Todos[i].Id = ""
				want.CreatedAt = got.Todos[i].CreatedAt
				want.UpdatedAt = got.Todos[i].UpdatedAt
				if !reflect.DeepEqual(want, got.Todos[i]) {
					t.Fatalf("%s: expected:%v got:%v", tc.name, want, got.Todos[i])
				}
			}
		})
	}
}

func TestTodoService_Update(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Todo
		want  pb.Todo
	}{
		{
			name: "successful",
			input: pb.Todo{
				Id:       "73718693-0a2b-41da-ba08-de7bf7940340",
				Assignee: "albert",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T00:00:00Z",
				Status:   "active",
			},
			want: pb.Todo{
				Assignee: "albert",
				Title:    "todo service",
				Summary:  "rpc implement",
				Deadline: "2021-12-15T00:00:00Z",
				Status:   "active",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to update todo", err)
			}
			got.Id = ""
			tc.want.CreatedAt = got.CreatedAt
			tc.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected:%v got:%v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoService_Delete(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  pb.Empty
	}{
		{
			name:  "successful",
			input: "73718693-0a2b-41da-ba08-de7bf7940340",
			want:  pb.Empty{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Delete(context.Background(), &pb.ByIdReq{Id: tc.input})
			if err != nil {
				t.Error("failed to delete todo", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected:%v got:%v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoService_ListOverdue(t *testing.T) {
	tests := []struct {
		name  string
		input *pb.ListTime
		wants []*pb.Todo
	}{
		{
			name: "succesful",
			input: &pb.ListTime{
				ListPage: &pb.ListReq{
					Page:  1,
					Limit: 2,
				},
				ToTime: "2021-12-04",
			},
			wants: []*pb.Todo{
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
			got, err := client.ListOverdue(context.Background(), tc.input)
			if err != nil {
				t.Error("failed to listOverdue todo", err)
			}
			for i, want := range tc.wants {
				got.Todos[i].Id = ""
				want.CreatedAt = got.Todos[i].CreatedAt
				want.UpdatedAt = got.Todos[i].UpdatedAt
				if !reflect.DeepEqual(want, got.Todos[i]) {
					t.Fatalf("%s: expected:%v got:%v", tc.name, want, got.Todos[i])
				}
			}
		})
	}
}
