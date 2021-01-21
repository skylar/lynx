package command

import (
  "net/url"
  "regexp"
)

type Command struct {
  name        string
  description string
  handler     CommandHander
  nicknames   []string
  detector    *regexp.Regexp
}

type CommandHander func(string) *url.URL


func (c *Command) GetName() string {
  return c.name
}

func (c *Command) GetDescription() string {
  return c.description
}

func (c *Command) GetNicknames() []string {
  return c.nicknames
}

func NewBookmarkCommand(name string, path string, desc string) *Command {
  commandUrl, err := url.ParseRequestURI(path)
  if err != nil {
    return nil
  }

  return &Command{
    name:         name,
    description:  desc,
    nicknames:    []string{},
    detector:     nil,
    handler:      func(param string) *url.URL {
    	return commandUrl
    },
  }
}

func NewSearchCommand(name string, searchPath string, desc string, nicks []string) *Command {
  return &Command{
    name:         name,
    description:  desc,
    nicknames:    nicks,
    detector:     nil,
    handler:      func(param string) *url.URL {
      url, _ := url.ParseRequestURI(searchPath + param)
    	return url
    },
  }
}

func NewIdResolverCommand(name string, base string, desc string, nicks []string, regex string) *Command {
  var detector *regexp.Regexp = nil
  var err error = nil
  if len(regex) > 0 {
    detector, err = regexp.Compile(regex)
    if err != nil {
      detector = nil
    }
  }

  return &Command{
    name:         name,
    description:  desc,
    nicknames:    nicks,
    detector:     detector,
    handler:      func(param string) *url.URL {
      urlString := base
      if len(param) > 0 {
        urlString += param
      }
      url, _ := url.ParseRequestURI(urlString)
    	return url
    },
  }
}
