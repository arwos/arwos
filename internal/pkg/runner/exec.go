package runner

import (
	"context"

	"go.osspkg.com/ioutils/shell"
)

type executer struct {
	sh shell.TShell
}

func newExecuter(ctx context.Context, dir string) *executer {
	sh := shell.New()
	sh.SetDir(dir)

	return &executer{sh: sh}
}
