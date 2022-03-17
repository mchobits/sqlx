package sqlx

import (
	"context"
	"errors"
	"fmt"
	"github.com/SkyAPM/go2sky"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"strings"
)

const (
	componentIDUnknown = 0
	componentIDMysql   = 5012
)

// ErrUnsupportedOp operation unsupported by the underlying driver
var ErrUnsupportedOp = errors.New("operation unsupported by the underlying driver")

func argsToString(args []interface{}) string {
	sb := strings.Builder{}
	for _, arg := range args {
		sb.WriteString(fmt.Sprintf("%v, ", arg))
	}
	return sb.String()
}

func createSpan(ctx context.Context, tracer *go2sky.Tracer, opts *options, operation string) (go2sky.Span, error) {
	s, _, err := tracer.CreateLocalSpan(ctx,
		go2sky.WithSpanType(go2sky.SpanTypeExit),
		go2sky.WithOperationName(opts.getOpName(operation)),
	)
	if err != nil {
		return nil, err
	}
	s.SetPeer(opts.peer)
	s.SetComponent(opts.componentID)
	s.SetSpanLayer(agentv3.SpanLayer_Database)
	s.Tag(go2sky.TagDBType, string(opts.dbType))
	s.Tag(go2sky.TagDBInstance, opts.peer)
	return s, nil
}
