package server

import (
	_ "fmt"

	"lynx/internal/command"
	"lynx/internal/db"

	"github.com/kataras/iris/v12"
)

type LynxServer struct {
	iris       *iris.Application
	config     *Configuration
	commandSet *command.CommandSet
	urlStore   db.URLStore
}

func (s LynxServer) Start() {
	addr := ":" + s.config.Port

	if s.urlStore != nil {
		iris.RegisterOnInterrupt(s.urlStore.Close)
	}
	s.iris.Logger().SetLevel("debug")
	s.iris.Run(iris.Addr(addr))
}

func (s LynxServer) setupContent() {
	view := iris.Handlebars("./content/views", ".html").
		Layout("layouts/main.html").
		Reload(true)
	s.iris.RegisterView(view)
	s.iris.HandleDir("/static", "./content/resources")
}

func (s LynxServer) setupRouting() {
	s.iris.Get("/", func(ctx iris.Context) {
		ctx.ViewData("title", "Lynx")
		ctx.View("index.html")
	})

	s.iris.Get("/list", func(ctx iris.Context) {
		commandInfo := s.commandSet.GetCommandInfo()
		ctx.ViewData("title", "Commands")
		ctx.ViewData("commandCount", len(commandInfo))
		ctx.ViewData("commands", commandInfo)
		ctx.View("list.html")
	})

	s.iris.Get("/s/{searchText}", func(ctx iris.Context) {
		url := s.commandSet.Resolve(ctx.Params().Get("searchText"))
		ctx.Redirect(url.String(), iris.StatusTemporaryRedirect)
	})

	// URL Shortener - Not Implemented
	s.iris.Get("/u/{shortcut}", func(ctx iris.Context) {
		ctx.NotFound()
	})
	s.iris.Post("/api/shorten", func(ctx iris.Context) {
		ctx.NotFound()
	})
}

func NewLynxServer(config *Configuration) *LynxServer {
	server := &LynxServer{
		iris:       iris.New(),
		config:     config,
		commandSet: command.NewCommandSet(),
		urlStore:   nil,
	}

	server.setupContent()
	server.setupRouting()
	return server
}
