package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"log/slog"
	"posts/internal/lib/pagination"
	mock_repo "posts/internal/repository/mocks"

	"posts/internal/model"
	"testing"
)

func IsEqualPost(want, got model.Post) bool {

	if want.ID != got.ID || want.CommentsAllowed != got.CommentsAllowed ||
		want.Content != got.Content {
		return false
	}

	return true
}

func TestPostsService_CreatePost(t *testing.T) {

	defaultPost := model.Post{
		ID:              1,
		Content:         "Cnt",
		CommentsAllowed: true,
		Name:            "Name",
	}

	tests := []struct {
		name     string
		post     model.Post
		want     model.Post
		wantErr  bool
		repoRes  model.Post
		repoErr  error
		repoSkip bool
	}{
		{
			name:     "Positive",
			post:     defaultPost,
			want:     defaultPost,
			wantErr:  false,
			repoErr:  nil,
			repoRes:  defaultPost,
			repoSkip: false,
		},
		{
			name:     "Error from repo",
			post:     defaultPost,
			want:     model.Post{},
			wantErr:  true,
			repoErr:  errors.New("some error"),
			repoRes:  model.Post{},
			repoSkip: false,
		},
	}

	log := slog.Default()
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctl := gomock.NewController(t)
			defer ctl.Finish()

			repo := mock_repo.NewMockRepoPosts(ctl)

			if !tt.repoSkip {
				repo.EXPECT().CreatePost(tt.post).Return(tt.repoRes, tt.repoErr)
			}

			p := NewPostsService(repo, log)

			got, err := p.CreatePost(ctx, tt.post)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !IsEqualPost(got, tt.want) {
				t.Errorf("CreatePost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostsService_GetAllPosts(t *testing.T) {

	data := []*model.Post{
		{
			ID:              1,
			Name:            "N1",
			Content:         "C1",
			CommentsAllowed: true,
		},
		{
			ID:              2,
			Name:            "N2",
			Content:         "nya",
			CommentsAllowed: true,
		},
		{
			ID:              3,
			Name:            "N3",
			Content:         "cool stuff",
			CommentsAllowed: false,
		},
	}

	type args struct {
		page     *int
		pageSize *int
	}
	tests := []struct {
		name     string
		want     []*model.Post
		wantErr  bool
		repoRes  []*model.Post
		repoErr  error
		repoSkip bool
		args     args
	}{
		{
			name:     "Positive",
			args:     args{page: nil, pageSize: nil},
			want:     data,
			wantErr:  false,
			repoRes:  data,
			repoErr:  nil,
			repoSkip: false,
		},
		{
			name:     "Error from repo",
			args:     args{page: nil, pageSize: nil},
			want:     nil,
			wantErr:  true,
			repoRes:  nil,
			repoErr:  errors.New("some error"),
			repoSkip: false,
		},
		{
			name:     "Empty result from repo",
			args:     args{page: nil, pageSize: nil},
			want:     nil,
			wantErr:  false,
			repoRes:  nil,
			repoErr:  nil,
			repoSkip: false,
		},
		{
			name:     "With not null page number",
			args:     args{page: ptr(1), pageSize: nil},
			want:     data,
			wantErr:  false,
			repoRes:  data,
			repoErr:  nil,
			repoSkip: false,
		},
		{
			name:     "With not null page size",
			args:     args{page: nil, pageSize: ptr(1)},
			want:     data,
			wantErr:  false,
			repoRes:  data,
			repoErr:  nil,
			repoSkip: false,
		},
		{
			name:     "With not null page size and number",
			args:     args{page: ptr(1), pageSize: ptr(3)},
			want:     data,
			wantErr:  false,
			repoRes:  data,
			repoErr:  nil,
			repoSkip: false,
		},
	}

	log := slog.Default()
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctl := gomock.NewController(t)
			defer ctl.Finish()

			repo := mock_repo.NewMockRepoPosts(ctl)

			limit, offset, err := pagination.GetLimitAndOffset(tt.args.page, tt.args.pageSize)
			if !tt.repoSkip {
				repo.EXPECT().GetAllPosts(limit, offset).Return(tt.repoRes, tt.repoErr)
			}

			p := NewPostsService(repo, log)
			got, err := p.GetAllPosts(ctx, limit, offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("GetAllPosts() got = %v len= %d, want %v len= %d", got, len(got), tt.want, len(tt.want))
				return
			}
			for i := 0; i < len(got); i++ {
				if !IsEqualPost(*got[i], *tt.want[i]) {
					t.Errorf("GetAllPosts() got = %v, want %v", got[i], tt.want[i])
					return
				}
			}

		})
	}
}

func TestPostsService_GetPostById(t *testing.T) {

	defaultPost := model.Post{
		ID:              1,
		Content:         "Cnt",
		CommentsAllowed: true,
		Name:            "Name",
	}

	log := slog.Default()
	ctx := context.Background()

	tests := []struct {
		name     string
		postId   int
		want     model.Post
		wantErr  bool
		repoRes  model.Post
		repoErr  error
		repoSkip bool
	}{
		{
			name:     "Positive",
			postId:   1,
			want:     defaultPost,
			wantErr:  false,
			repoRes:  defaultPost,
			repoErr:  nil,
			repoSkip: false,
		},
		{
			name:     "Error from repo",
			postId:   1,
			want:     model.Post{},
			wantErr:  true,
			repoRes:  model.Post{},
			repoErr:  errors.New("some error"),
			repoSkip: false,
		},
		{
			name:     "Sql no errors from repo",
			postId:   1,
			want:     model.Post{},
			wantErr:  true,
			repoRes:  model.Post{},
			repoErr:  sql.ErrNoRows,
			repoSkip: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			repo := mock_repo.NewMockRepoPosts(ctl)

			if !tt.repoSkip {
				repo.EXPECT().GetPostById(tt.postId).Return(tt.repoRes, tt.repoErr)
			}

			p := NewPostsService(repo, log)

			got, err := p.GetPostById(ctx, tt.postId)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetPostById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !IsEqualPost(got, tt.want) {
				t.Errorf("GetPostById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
