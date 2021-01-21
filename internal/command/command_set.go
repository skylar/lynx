package command

import (
	"net/url"
	"sort"
	"strings"
)

type CommandSet struct {
	commands       commandMap
	commandIndex   commandMap
	commandInfo    []CommandInfo
	detectors      []*Command
	defaultCommand *Command
}

type commandMap map[string]*Command
type commandList []*Command

const GoogleCalendarUrlString = "https://calendar.google.com"
const GoogleMailUrlString = "https://mail.google.com"

func NewCommandSet() *CommandSet {
	googleCommand := NewGoogleCommand()
	commands := []*Command{
		NewZoomCommand(),
		NewBluejeansResolver(),
		googleCommand,
		NewTwitterCommand(),
		NewJiraResolver(),
		NewListCommand(),
		NewWikipediaCommand(),
		NewBtcAddressResolver(),
		NewXrpAddressResolver(),
	}
	bookmarks := buildBookmarkCommands(
		map[string]string{
			"cal":  GoogleCalendarUrlString,
			"mail": GoogleMailUrlString,
		})
	commands = append(commands, bookmarks...)

	commandMap := buildCommandMap(commands)
	e := &CommandSet{
		commands:       commandMap,
		commandIndex:   buildCommandIndex(commands),
		commandInfo:    buildCommandInfoList(commandMap),
		detectors:      buildDetectorList(commands),
		defaultCommand: googleCommand,
	}

	return e
}

func (cs CommandSet) GetCommandInfo() []CommandInfo {
	return cs.commandInfo
}

func buildBookmarkCommands(bookmarks map[string]string) commandList {
	list := make([]*Command, 0, len(bookmarks))
	for key, urlString := range bookmarks {
		cmd := NewBookmarkCommand(key, urlString, urlString)
		list = append(list, cmd)
	}
	return list
}

func buildCommandMap(commands []*Command) commandMap {
	commandMap := make(map[string]*Command)

	for _, r := range commands {
		commandMap[r.name] = r
	}
	return commandMap
}

func buildCommandIndex(commands []*Command) commandMap {
	commandIndex := make(map[string]*Command)

	for _, r := range commands {
		commandIndex[r.name] = r
		for _, nick := range r.nicknames {
			commandIndex[nick] = r
		}
	}
	return commandIndex
}

func buildCommandInfoList(commands commandMap) []CommandInfo {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	info := make([]CommandInfo, 0, len(commands))
	for _, name := range names {
		info = append(info, commands[name])
	}
	return info
}

func buildDetectorList(commands []*Command) commandList {
	detectors := make(commandList, len(commands))
	count := 0

	for _, r := range commands {
		if r.detector != nil {
			detectors[count] = r
			count += 1
		}
	}
	return detectors[:count]
}

func (cs CommandSet) Resolve(searchText string) *url.URL {
	var r *Command = nil
	var ok bool = false

	cmd, param := parseSearch(searchText)

	// see if we can find a command by name/nickname
	r, ok = cs.commandIndex[cmd]
	if ok {
		return r.handler(param)
	}

	// try to detect the resolver based on content
	r, ok = cs.detectResolver(cmd)
	if ok {
		// implict command
		param = cmd
		return r.handler(param)
	}

	// default
	return cs.defaultCommand.handler(searchText)
}

func parseSearch(input string) (string, string) {
	var parameter = ""
	parts := strings.SplitN(strings.TrimSpace(input), " ", 2)
	command := strings.TrimSpace(parts[0])
	if len(parts) > 1 {
		parameter = strings.TrimSpace(parts[1])
	}
	return command, parameter
}

func (e CommandSet) detectResolver(input string) (*Command, bool) {
	for _, command := range e.detectors {
		if command.detector.MatchString(input) {
			return command, true
		}
	}
	return nil, false
}
