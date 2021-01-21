package command

type CommandInfo interface {
  GetName()         string
  GetDescription()  string
  GetNicknames()    []string
}
