package qbuilder

import (
	"strings"

	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/repository"
)

type QBuilder struct {
	query   string
	filters []repository.Filter
	l       log.Logger
	b       *strings.Builder
}

func NewQBuilder(logger log.Logger, query string, filters ...repository.Filter) QBuilder {
	return QBuilder{
		query:   query,
		filters: filters,
		l:       logger,
		b:       &strings.Builder{},
	}
}

func (qb QBuilder) Build() string {
	qb.b.WriteString(qb.query)

	if len(qb.filters) > 0 {
		qb.where()
	}

	return qb.b.String()
}

func (qb QBuilder) where() {
	if !strings.Contains(qb.b.String(), "WHERE") {
		qb.b.WriteByte(' ')
		qb.b.WriteString("WHERE")
	}

	for i, f := range qb.filters {
		if err := f.Apply(qb.b); err != nil {
			qb.l.Warn(err)
			continue
		}

		if i < len(qb.filters)-1 {
			qb.and(qb.b)
		}
	}
}

func (qb QBuilder) and(b *strings.Builder) *strings.Builder {
	b.WriteByte(' ')
	b.WriteString("AND")
	return b
}
