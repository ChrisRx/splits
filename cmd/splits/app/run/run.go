package run

import (
	"context"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/spf13/cobra"

	"github.com/ChrisRx/splits/pkg/run"
	"github.com/ChrisRx/splits/pkg/tui"
	"github.com/ChrisRx/splits/pkg/util"
)

var (
	defaultStyle = tcell.StyleDefault.
			Background(tcell.NewHexColor(0x2f3542)).
			Foreground(tcell.ColorWhite)
	segmentStyle = tcell.StyleDefault.
			Background(tcell.NewHexColor(0x2f3542)).
			Foreground(tcell.NewHexColor(0xa4b0be))
	timerStyle = tcell.StyleDefault.
			Background(tcell.NewHexColor(0x2f3542)).
			Foreground(tcell.ColorSienna)
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "run",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := run.Load(args[0])
			if err != nil {
				return err
			}

			encoding.Register()

			s, err := tcell.NewScreen()
			if err != nil {
				return err
			}
			if err := s.Init(); err != nil {
				return err
			}

			s.SetStyle(defaultStyle)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			go func() {
				ticker := time.NewTicker(100 * time.Millisecond)
				defer ticker.Stop()

				for {
					select {
					case <-ticker.C:
						s.Clear()
						tui.DrawTitle(s, r.Name, r.Category)
						tui.DrawSegments(s, segmentStyle, r.Segments())
						tui.DrawTimer(s, timerStyle, util.FormatDuration(r.CurrentTime()))
						s.Show()
					case <-ctx.Done():
						return
					}
				}
			}()

			for {
				switch ev := s.PollEvent().(type) {
				case *tcell.EventResize:
					s.Sync()
					s.Clear()
					tui.DrawTitle(s, r.Name, r.Category)
					tui.DrawSegments(s, segmentStyle, r.Segments())
					tui.DrawTimer(s, timerStyle, util.FormatDuration(r.CurrentTime()))
					s.Show()
				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyEscape:
						s.Fini()
						os.Exit(0)
					case tcell.KeyRune:
						switch ev.Rune() {
						case 'q':
							s.Fini()
							os.Exit(0)
						case ' ':
							if !r.Running() {
								r.Start()
								continue
							}
							r.Split()
						case 'r':
							r.Reset()
						case 's':
							r.Save(args[0])
						}
					}
				}
			}
		},
	}
	return cmd
}
