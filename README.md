Colors is a minimalist vscode setup

How to install:
```
1. in vscode press ctrl + shift + x
2. find and install 'Black on Light Gray' theme
3. find path of settings.json (somewhere in vscode folder)
4. add path as target in colors.json 
5. maker sure colors.json and colors bin are in the same folder
6. run colors bin
```

There are 5 tweakable colors available. The whole theme revolves
around them. Check the type annotations for more details.

```go
type Colors struct {
	BackgroundDark  string
	BackgroundLight string
	Text            string
	Comments        string
	Cursor          string
}
```