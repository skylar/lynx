package command

import(
  "testing"

  "github.com/stretchr/testify/assert"
)


func TestCommand(t *testing.T) {
  const searchName = "search"
  const searchPath = "http://mysearch.com/?q="
  const searchDescription = "Find anything."
  var searchNicks = []string{"s", "?"}

  cmd := NewSearchCommand(searchName, searchPath, searchDescription, searchNicks)

  assert.Equal(t, searchName, cmd.GetName())
  assert.Equal(t, searchDescription, cmd.GetDescription())
  assert.Equal(t, searchNicks, cmd.GetNicknames())
}
