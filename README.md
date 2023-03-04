Colors is a tool for setting up VSCode

There are 5 tweakable colors available. The whole theme revolves
around them. Check type annotations for more details.

```go
type Colors struct {
  BackgroundDark  string
  BackgroundLight string
  Text            string
  Comments        string
  Cursor          string
}

type Config struct {
  Target string // path to settings.json
  Colors Colors
}
```
*Colors uses the 'Black on Light Gray' theme as base.
Make sure it's installed before running the tool.