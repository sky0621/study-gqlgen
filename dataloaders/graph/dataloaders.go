package graph

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sky0621/study-gqlgen/dataloaders/graph/model"
)

const loadersKey = "dataLoaders"

type Loaders struct {
	UsersByIDs     UserLoader
	TodosByUserIDs TodoLoader
}

func Middleware(conn *sqlx.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UsersByIDs: UserLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				// 最大１ミリ秒待機した結果、ないし、最大 100 個のGraphQLクエリが溜まった分の id のスライスが ids という名前で渡ってくる。
				fetch: func(ids []int64) ([]*model.User, []error) {
					if len(ids) == 0 {
						return nil, nil
					}

					sql := "SELECT * FROM user WHERE id IN (" + toPKs(ids) + ")"
					log.Print(sql)

					var users []*model.User
					if err := conn.SelectContext(r.Context(), &users, sql); err != nil {
						log.Print(err)
						return nil, []error{err}
					}

					// ids の中の id 毎にデータをマッピングする必要がある。
					userById := map[int64]*model.User{}
					for _, user := range users {
						userById[user.ID] = user
					}
					results := make([]*model.User, len(ids))
					for i, id := range ids {
						results[i] = userById[id]
					}

					return results, nil
				},
			},
			TodosByUserIDs: TodoLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				// 最大１ミリ秒待機した結果、ないし、最大 100 個のGraphQLクエリが溜まった分の id のスライスが ids という名前で渡ってくる。
				fetch: func(userIDs []int64) ([][]*model.Todo, []error) {
					if len(userIDs) == 0 {
						return nil, nil
					}

					sql := "SELECT * FROM todo WHERE user_id IN (" + toPKs(userIDs) + ")"
					log.Print(sql)

					var todos []*model.Todo
					if err := conn.SelectContext(r.Context(), &todos, sql); err != nil {
						log.Print(err)
						return nil, []error{err}
					}

					// ids の中の id 毎にデータをマッピングする必要がある。
					todoByUserId := map[int64][]*model.Todo{}
					for _, todo := range todos {
						todoByUserId[todo.UserID] = append(todoByUserId[todo.UserID], todo)
					}
					results := make([][]*model.Todo, len(userIDs))
					for i, id := range userIDs {
						results[i] = todoByUserId[id]
					}

					return results, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func toPKs(ids []int64) string {
	// ids をSQL文の IN 句に指定できる形に変換
	var pks []string
	for _, id := range ids {
		pks = append(pks, strconv.FormatInt(id, 10))
	}
	return strings.Join(pks, ",")
}
