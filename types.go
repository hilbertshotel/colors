package main

type Colors struct {
	BackgroundDark  string
	BackgroundLight string
	Text            string
	Comments        string
	Cursor          string
}

type Config struct {
	Target string
	Colors Colors
}
