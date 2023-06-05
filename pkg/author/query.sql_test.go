package author

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func Test_CreateAuthor(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreateAuthorParams
	}

	q := `-- name: CreateAuthor :one
		INSERT INTO authors (
			name, bio
		) VALUES (
			$1, $2
		)
		RETURNING id, name, bio
	`

	qGet := `-- name: GetAuthor :one
		SELECT id, name, bio FROM authors
		WHERE id = $1 LIMIT 1
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     Author
		wantErr  bool
	}{
		{
			name: "success create author",
			args: args{
				ctx: context.Background(),
				arg: CreateAuthorParams{
					Name: "boba",
					Bio: sql.NullString{
						String: "hello",
						Valid:  true,
					},
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(1, "boba", "hello")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("boba", "hello").WillReturnRows(rows)

				rows = sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(1, "boba", "hello")
				mock.ExpectQuery(regexp.QuoteMeta(qGet)).WithArgs(int64(1)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: Author{
				ID:   1,
				Name: "boba",
				Bio: sql.NullString{
					String: "hello",
					Valid:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "error create author",
			args: args{
				ctx: context.Background(),
				arg: CreateAuthorParams{
					Name: "boba",
					Bio: sql.NullString{
						String: "hello",
						Valid:  true,
					},
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("boba", "hello").WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    Author{},
			wantErr: true,
		},
		{
			name: "error get author",
			args: args{
				ctx: context.Background(),
				arg: CreateAuthorParams{
					Name: "boba",
					Bio: sql.NullString{
						String: "hello",
						Valid:  true,
					},
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(1, "boba", "hello")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("boba", "hello").WillReturnRows(rows)

				rows = sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(1, "boba", "hello")
				mock.ExpectQuery(regexp.QuoteMeta(qGet)).WithArgs(int64(1)).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    Author{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.CreateAuthor(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeleteAuthor(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	q := `-- name: DeleteAuthor :exec
		DELETE FROM authors
		WHERE id = $1
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		wantErr  bool
	}{
		{
			name: "success delete author",
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(int64(1)).WillReturnResult(sqlmock.NewResult(1, 1))

				return &Queries{
					db: dbMock,
				}
			},
			wantErr: false,
		},
		{
			name: "error delete author",
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(int64(1)).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			err := p.DeleteAuthor(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetAuthor(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	q := `-- name: GetAuthor :one
		SELECT id, name, bio FROM authors
		WHERE id = $1 LIMIT 1
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     Author
		wantErr  bool
	}{
		{
			name: "success get author",
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(1, "boba", "hello")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(int64(1)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: Author{
				ID:   1,
				Name: "boba",
				Bio: sql.NullString{
					String: "hello",
					Valid:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "error get author",
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(int64(1)).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    Author{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.GetAuthor(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ListAuthors(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	q := `-- name: ListAuthors :many
		SELECT id, name, bio FROM authors
		ORDER BY name
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     []Author
		wantErr  bool
	}{
		{
			name: "success get list author",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(1, "boba", "hello").AddRow(2, "doda", "goga")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: []Author{
				{
					ID:   1,
					Name: "boba",
					Bio: sql.NullString{
						String: "hello",
						Valid:  true,
					},
				},
				{
					ID:   2,
					Name: "doda",
					Bio: sql.NullString{
						String: "goga",
						Valid:  true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error get list author",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error scan get list author",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow("boba", 2, "hello")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.ListAuthors(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAuthors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAuthors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UpdateAuthor(t *testing.T) {
	type args struct {
		ctx context.Context
		arg UpdateAuthorParams
	}

	q := `-- name: UpdateAuthor :one
		UPDATE authors
		set name = $2,
		bio = $3
		WHERE id = $1
		RETURNING id, name, bio
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     Author
		wantErr  bool
	}{
		{
			name: "success get list author",
			args: args{
				ctx: context.Background(),
				arg: UpdateAuthorParams{
					ID:   int64(1),
					Name: "yeye",
					Bio: sql.NullString{
						String: "hey",
						Valid:  true,
					},
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(1, "yeye", "hey")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: Author{
				ID:   1,
				Name: "yeye",
				Bio: sql.NullString{
					String: "hey",
					Valid:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "error get list author",
			args: args{
				ctx: context.Background(),
				arg: UpdateAuthorParams{
					ID:   int64(1),
					Name: "yeye",
					Bio: sql.NullString{
						String: "hey",
						Valid:  true,
					},
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    Author{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.UpdateAuthor(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}
