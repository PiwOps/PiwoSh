// logger/styles.go
package logger

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func GetPackageLoggerStyle(packageName string) *log.Styles {
	styles := log.DefaultStyles()

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("[DEBU] [%s]", packageName)).
		Padding(0).
		Foreground(lipgloss.Color("#1E90FF")).
		Bold(true)
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("[INFO] [%s]", packageName)).
		Padding(0).
		Foreground(lipgloss.Color("#CCCCCC")).
		Bold(true)
	styles.Levels[log.Level(SuccessLevel)] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("[SUCC] [%s]", packageName)).
		Padding(0).
		Foreground(lipgloss.Color("#00FF00")).
		Bold(true)
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("[WARN] [%s]", packageName)).
		Padding(0).
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true)
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("[ERRO] [%s]", packageName)).
		Padding(0).
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("[FATA] [%s]", packageName)).
		Padding(0).
		Foreground(lipgloss.Color("#8B0000")).
		Bold(true).
		Blink(true)

	return styles
}

func GetDefaultLoggerStyle() *log.Styles {
	styles := log.DefaultStyles()

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("[DEBU]").
		Padding(0).
		Foreground(lipgloss.Color("#1E90FF")).
		Bold(true)
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("[INFO]").
		Padding(0).
		Foreground(lipgloss.Color("#CCCCCC")).
		Bold(true)
	styles.Levels[log.Level(SuccessLevel)] = lipgloss.NewStyle().
		SetString("[SUCC]").
		Padding(0).
		Foreground(lipgloss.Color("#00FF00")).
		Bold(true)
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("[WARN]").
		Padding(0).
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true)
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("[ERRO]").
		Padding(0).
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("[FATA]").
		Padding(0).
		Foreground(lipgloss.Color("#8B0000")).
		Bold(true).
		Blink(true)

	return styles
}
