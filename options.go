package sqlx

type DBType string

const (
	// MYSQL mysql
	MYSQL DBType = "mysql"
	// IPV4 others database type
	IPV4 DBType = "others"
)

// Option set plugin option
type Option func(*options)

type options struct {
	dbType      DBType
	peer        string
	componentID int32

	reportQuery bool
	reportParam bool
}

// WithSQLDBType set dbType option,
// dbType is used for parsing dsn string to peer address
// and setting componentID, if DB type is not support in DBType
// list, please use WithPeerAddr to set peer address manually
func WithSQLDBType(t DBType) Option {
	return func(o *options) {
		o.dbType = t
		o.setComponentID()
	}
}

// WithPeerAddr set the peer address to report
func WithPeerAddr(addr string) Option {
	return func(o *options) {
		o.peer = addr
	}
}

// WithQueryReport if set, the sql would be collected
func WithQueryReport() Option {
	return func(o *options) {
		o.reportQuery = true
	}
}

// WithParamReport if set, the parameters of the sql would be collected
func WithParamReport() Option {
	return func(o *options) {
		o.reportParam = true
	}
}

func (o options) getOpName(op string) string {
	switch o.dbType {
	case MYSQL:
		return "Mysql/Go2Sky/" + op
	default:
		return "Sql/Go2Sky/" + op
	}
}

func (o *options) setComponentID() {
	switch o.dbType {
	case MYSQL:
		o.componentID = componentIDMysql
	default:
		o.componentID = componentIDUnknown
	}
}
