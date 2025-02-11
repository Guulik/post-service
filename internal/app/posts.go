package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"os"
	"posts/configure"
	"posts/graph"
	ql "posts/internal/api/graph"
	"posts/internal/constants"
	"posts/internal/repository"
	in_memory "posts/internal/repository/in-memory"
	"posts/internal/repository/postgres"
	"posts/internal/service"
)

const (
	envLocal = "local"
	envStage = "stage"
)

type App struct {
	cfg    *configure.Config
	server *http.Server
	svc    *service.Services
	repo   *repository.Repo
	pool   *sqlx.DB
}

func (a *App) routes() *http.ServeMux {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &ql.Resolver{
		PostsService:      a.svc.Posts,
		CommentsService:   a.svc.Comments,
		CommentsObservers: ql.NewCommentsObserver(),
	}}))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)
	return mux
}

func New() *App {
	app := &App{}

	cfg := configure.MustLoadConfig()
	log := setupLogger(cfg.Env)

	app.cfg = cfg

	if cfg.InMemory {
		posts := in_memory.NewPostsInMemory(constants.PostsPullSize)
		comments := in_memory.NewCommentsInMemory(constants.CommentsPullSize)
		app.repo = repository.NewRepo(posts, comments)
	} else {
		app.pool = configure.NewPostgres(cfg)
		posts := postgres.NewPostsPostgres(app.pool)
		comments := postgres.NewCommentsPostgres(app.pool)
		app.repo = repository.NewRepo(posts, comments)

		err := cfg.MigrateUp()
		if err != nil {
			fmt.Println("migration error:", err)
		}
	}

	app.svc = service.NewServices(app.repo, log)

	app.server = &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	return app
}

func (a *App) Run() {
	fmt.Println("server running")
	fmt.Println(fmt.Sprintf("Connect to http://localhost:%d/ for GraphQL playground", a.cfg.Port))

	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server error:", err)
	}
}

func (a *App) Stop(ctx context.Context) error {
	fmt.Println("stopping server...")

	if err := a.server.Shutdown(ctx); err != nil {
		fmt.Println("failed to shutdown server:", err)
		return err
	}

	if a.pool != nil {
		if err := a.pool.Close(); err != nil {
			fmt.Println("failed to close connection:", err)
		}
	}

	return nil
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envStage:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
