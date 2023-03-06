package main

const SETTINGS = `{
    "workbench.statusBar.visible": false,
    "workbench.activityBar.visible": false,
    "git.enabled": false,
    "editor.renderLineHighlight": "none",
    "breadcrumbs.enabled": false,
    "editor.minimap.enabled": false,
    "editor.detectIndentation": false,
    "explorer.confirmDelete": false,
    "editor.guides.highlightActiveIndentation": false,
    "explorer.confirmDragAndDrop": false,
    "workbench.startupEditor": "none",
    "editor.semanticHighlighting.enabled": false,
    "editor.occurrencesHighlight": false,
    "workbench.colorTheme": "Black on Light Gray",
    "workbench.iconTheme": "vs-minimal",
    "workbench.tree.renderIndentGuides": "none",
    "editor.matchBrackets": "never",
    "editor.bracketPairColorization.enabled": false,
    "window.commandCenter": false,
    "elixirLS.suggestSpecs": false,
    "window.zoomLevel": 2,
    "go.toolsManagement.autoUpdate": true,
    "editor.renderWhitespace": "none",
    "editor.cursorStyle": "line",
    "telemetry.telemetryLevel": "off",
    "editor.unicodeHighlight.allowedCharacters": {
        " ": true
    },
    "workbench.colorCustomizations": {

        "tab.activeBackground": "{{ .BackgroundDark }}",
        "tab.inactiveBackground": "{{ .BackgroundLight }}",
        "tab.border": "{{ .BackgroundDark }}",
        "tab.activeForeground": "{{ .Text }}",
        "editorGroupHeader.tabsBackground": "{{ .BackgroundLight }}",
        "scrollbar.shadow": "#a3a3a300",
        "editorGroupHeader.tabsBorder": "#a3a3a300",
        "tab.hoverForeground": "{{ .Text }}",

        "editor.background": "{{ .BackgroundDark }}",
        "editor.foreground": "{{ .Text }}",
        "editor.selectionBackground": "{{ .BackgroundLight }}",
        "editor.selectionHighlightBackground": "{{ .BackgroundLight }}",

        "editorLineNumber.foreground": "{{ .Comments }}",
        "editorLineNumber.activeForeground": "{{ .Text }}",

        "editorIndentGuide.background": "{{ .BackgroundLight }}",

        "editorSuggestWidget.background": "{{ .BackgroundLight }}",
        "editorSuggestWidget.border": "{{ .BackgroundLight }}",
		"editorHoverWidget.background": "{{ .BackgroundLight }}",

        "editorCursor.background": "{{ .BackgroundLight }}",
        "editorCursor.foreground": "{{ .Cursor }}",

        "sideBar.background": "{{ .BackgroundLight }}",
        "list.activeSelectionBackground": "{{ .BackgroundDark }}",
        "list.focusOutline": "{{ .BackgroundDark }}",
        "list.focusForeground": "{{ .Text }}",
        "list.hoverForeground": "{{ .Text }}",
        "list.activeSelectionForeground": "{{ .Text }}",
        "list.inactiveSelectionBackground": "{{ .BackgroundDark }}",
        "list.inactiveSelectionForeground": "{{ .Text }}",
        "list.hoverBackground": "#00000000",
		
        "input.background": "{{ .BackgroundLight }}",
        "dropdown.background": "{{ .BackgroundLight }}"
    },
    "editor.tokenColorCustomizations": {
        "[Black on Light Gray]": {
            "comments": "{{ .Comments }}"
        }
    }
}
`
