package server

import (
	_ "fmt"

	"lynx/internal/command"
	"lynx/internal/db"
	"lynx/internal/shortcuts"

	"github.com/kataras/iris/v12"
)

type LynxServer struct {
	iris       *iris.Application
	config     *Configuration
	commandSet *command.CommandSet
	shortener  *shortcuts.Controller
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

	// Commands
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

	// URL Shortener
	s.iris.Get("/u/{shortcut}", func(ctx iris.Context) {
		shortcut := ctx.Params().Get("shortcut")
		url := s.urlStore.Get(shortcut)

		if len(url) == 0 {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.Writef("No entry for shortcut '%s'.", shortcut)
			return
		}
		ctx.Redirect(url, iris.StatusTemporaryRedirect)
	})
	s.iris.Post("/api/shorten", func(ctx iris.Context) {
		urlString := ctx.FormValue("url")
		shortcode := ctx.FormValue("shortcode")
		response := s.shortener.Shorten(urlString, shortcode)

		ctx.StatusCode(response.StatusCode)
		ctx.JSON(response.Data)
	})
}

func NewLynxServer(config *Configuration) *LynxServer {
	irisApp := iris.New()
	store := createStore(config)
	server := &LynxServer{
		iris:       irisApp,
		config:     config,
		commandSet: command.NewCommandSet(),
		shortener: shortcuts.NewController(
			irisApp.ConfigurationReadOnly().GetVHost(),
			store,
		),
		urlStore: store,
	}

	server.setupContent()
	server.setupRouting()
	return server
}

func createStore(config *Configuration) db.URLStore {
	switch config.StoreType {
	case BoltStoreType:
		return db.NewBoltStore(config.StoreConnectionString)
	case MapStoreType:
		return db.NewMapStore()
	default:
		return nil
	}
}
