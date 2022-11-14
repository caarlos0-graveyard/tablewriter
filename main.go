// Package tablewriter provides a simple way to print tables using text/tabwriter and lpgloss.
package tablewriter

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/charmbracelet/lipgloss"
)

var boldStyle = lipgloss.NewStyle().Bold(true)

// Liner takes an item and returns the values its columns should have.
type Liner[T any] func(t T) ([]string, error)

// Render renders the table to the given io.Writer.
func Render[T any](w io.Writer, items []T, columns []string, liner Liner[T]) error {
	tw := newTabWriter(w, columns...)
	for _, item := range items {
		line, err := liner(item)
		if err != nil {
			return err
		}
		fmt.Fprintln(tw, strings.Join(line, "\t"))
	}
	if len(items) == 0 {
		fmt.Fprintln(tw, "No items found")
	}
	return tw.Flush()
}

func newTabWriter(w io.Writer, columns ...string) *tabwriter.Writer {
	fmt.Fprintln(w)
	tw := tabwriter.NewWriter(w, 8, 4, 4, ' ', 0)
	var h string
	for _, c := range columns {
		h = h + c + "\t    "
	}
	if h != "" {
		fmt.Fprintln(tw, boldStyle.Render(h))
	}
	return tw
}
