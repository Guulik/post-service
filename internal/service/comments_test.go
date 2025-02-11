package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"log/slog"
	"posts/internal/lib/response_error"
	"posts/internal/model"
	mock_repo "posts/internal/repository/mocks"
	mock_service "posts/internal/service/mocks"
	"testing"
)

func IsEqualComment(want, got model.Comment) bool {
	if want.ID != got.ID || want.Post != got.Post || want.Content != got.Content {
		return false
	}

	if (want.ReplyTo == nil) != (got.ReplyTo == nil) {
		return false
	}

	if (want.ReplyTo != nil) && (got.ReplyTo != nil) && *got.ReplyTo != *want.ReplyTo {
		return false
	}

	return true
}

func TestCommentsService_CreateComment(t *testing.T) {

	type FuncRes struct {
		com  model.Comment
		err  error
		skip bool
	}

	defaultComment := model.Comment{
		ID:      1,
		Post:    1,
		Content: "Text1",
		ReplyTo: nil,
	}

	defaultPost := model.Post{
		ID:              1,
		Content:         "Cnt2",
		CommentsAllowed: true,
		Name:            "Name",
	}

	postWithoutAllowedComments := model.Post{
		ID:              1,
		Content:         "Cnt3",
		CommentsAllowed: false,
		Name:            "Name3",
	}

	type args struct {
		comment model.Comment
	}
	tests := []struct {
		name       string
		args       args
		want       FuncRes
		repoRes    FuncRes
		getterRes  model.Post
		getterErr  error
		skipGetter bool
	}{
		{
			name:      "Positive creation",
			args:      args{comment: defaultComment},
			want:      FuncRes{com: defaultComment, err: nil},
			repoRes:   FuncRes{com: defaultComment, err: nil},
			getterRes: defaultPost,
			getterErr: nil,
		},
		{
			name:      "Comments Not Allowed",
			args:      args{comment: defaultComment},
			want:      FuncRes{com: model.Comment{}, err: response_error.ResponseError{}},
			repoRes:   FuncRes{com: defaultComment, err: nil, skip: true},
			getterRes: postWithoutAllowedComments,
			getterErr: nil,
		},
		{
			name:      "Post Not Found",
			args:      args{comment: defaultComment},
			want:      FuncRes{com: model.Comment{}, err: response_error.ResponseError{}},
			repoRes:   FuncRes{com: model.Comment{}, err: nil, skip: true},
			getterRes: model.Post{},
			getterErr: sql.ErrNoRows,
		},
		{
			name:      "Error with creating",
			args:      args{comment: defaultComment},
			want:      FuncRes{com: model.Comment{}, err: response_error.ResponseError{}},
			repoRes:   FuncRes{com: model.Comment{}, err: errors.New("some error")},
			getterRes: defaultPost,
			getterErr: nil,
		},
	}

	log := slog.Default()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mock_repo.NewMockRepoComments(controller)
			getter := mock_service.NewMockPostProvider(controller)

			if !tt.repoRes.skip {
				repo.EXPECT().CreateComment(tt.args.comment).Return(tt.repoRes.com, tt.repoRes.err)
			}

			if !tt.skipGetter {
				getter.EXPECT().GetPostById(tt.args.comment.Post).Return(tt.getterRes, tt.getterErr)
			}

			c := NewCommentsService(repo, getter, log)

			got, err := c.CreateComment(ctx, tt.args.comment)

			if (err != nil) != (tt.want.err != nil) {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tt.want.err)
				return
			}

			require.Equal(t, true, IsEqualComment(got, tt.want.com))
		})
	}
}

func TestCommentsService_GetRepliesOfComment(t *testing.T) {

	data := []*model.Comment{
		{
			ID:      2,
			Post:    1,
			Content: "Test2",
			ReplyTo: ptr(1),
		},
		{
			ID:      3,
			Post:    1,
			Content: "Test3",
			ReplyTo: ptr(1),
		},
		{
			ID:      4,
			Post:    1,
			Content: "Test6",
			ReplyTo: ptr(1),
		},
		{
			ID:      5,
			Post:    1,
			Content: "Test8",
			ReplyTo: ptr(1),
		},
	}

	type args struct {
		commentId int
	}
	tests := []struct {
		name     string
		args     args
		want     []*model.Comment
		wantErr  bool
		repoRes  []*model.Comment
		repoErr  error
		repoSkip bool
	}{
		{
			name:     "Positive Getting",
			args:     args{commentId: 1},
			want:     data,
			wantErr:  false,
			repoRes:  data,
			repoSkip: false,
			repoErr:  nil,
		},
		{
			name:     "Error from repo",
			args:     args{commentId: 1},
			want:     nil,
			wantErr:  true,
			repoRes:  nil,
			repoSkip: false,
			repoErr:  errors.New("error"),
		},
		{
			name:     "Nil result from repo",
			args:     args{commentId: 1},
			want:     nil,
			wantErr:  false,
			repoRes:  nil,
			repoSkip: false,
			repoErr:  nil,
		},
	}

	log := slog.Default()
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctl := gomock.NewController(t)
			defer ctl.Finish()

			repo := mock_repo.NewMockRepoComments(ctl)
			getter := mock_service.NewMockPostProvider(ctl)

			if !tt.repoSkip {
				repo.EXPECT().GetRepliesOfComment(tt.args.commentId).Return(tt.repoRes, tt.repoErr)
			}

			c := NewCommentsService(repo, getter, log)

			got, err := c.GetRepliesOfComment(ctx, tt.args.commentId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepliesOfComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("GetCommentsByPost() got = %v len= %d, want %v len= %d", got, len(got), tt.want, len(tt.want))
				return
			}
			for i := 0; i < len(got); i++ {
				if !IsEqualComment(*got[i], *tt.want[i]) {
					t.Errorf("GetCommentsByPost() got = %v, want %v", got[i], tt.want[i])
					return
				}
			}
		})
	}
}

func ptr(i int) *int {
	p := i
	return &p
}
