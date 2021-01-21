package command

import(
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestBluejeansCommand(t *testing.T) {
  cmd := NewBluejeansResolver();
  runCommandTests(t, cmd, "bluejeans")
}

func TestGoogleCommand(t *testing.T) {
  cmd := NewGoogleCommand();
  runCommandTests(t, cmd, "google")
}

func TestJiraCommand(t *testing.T) {
  cmd := NewJiraResolver();
  runCommandTests(t, cmd, "jira")
}

func TestListCommand(t *testing.T) {
  cmd := NewListCommand();
  assert.Equal(t, "list", cmd.name)
}

func TestTwitterCommand(t *testing.T) {
  cmd := NewTwitterCommand();
  runCommandTests(t, cmd, "twitter")
}

func TestWikipediaCommand(t *testing.T) {
  cmd := NewWikipediaCommand();
  runCommandTests(t, cmd, "wikipedia")
}

func TestZoomCommand(t *testing.T) {
  num := "1234"
  name := "jane1"

  cmd := NewZoomCommand();
  runCommandTests(t, cmd, "zoom")

  result := cmd.handler(name)
  assert.Equal(t, ZoomUsernameBaseString + name, result.String())
  result = cmd.handler(num)
  assert.Equal(t, ZoomBaseString + num, result.String())
}

func TestXrpAddressCommand(t *testing.T) {
  cmd := NewXrpAddressResolver();
  runCommandTests(t, cmd, "xrp")
}

func runCommandTests(t *testing.T, cmd *Command, name string) {
  assert.Equal(t, name, cmd.name)

  result := cmd.handler("foo")
  assert.True(t, result.IsAbs())
}
