package ipam

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"network/skydive-visualizer-go/graph"
	"network/skydive-visualizer-go/source"
	"network/skydive-visualizer-go/source/skydive"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
)

type IPAM struct {
	cfg   *Config
	inner source.Source
	db    *sql.DB
}

func New(inner source.Source, cfg *Config) (*IPAM, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	))
	if err != nil {
		return nil, errors.Wrap(err, "ipam: open")
	}

	return &IPAM{
		cfg:   cfg,
		inner: inner,
		db:    db,
	}, nil
}

func (s *IPAM) Fetch() (graph.Graph, error) {
	log.Info("ipam source: starting")
	defer log.Info("ipam source: done")

	grc := source.WillFetch(s.inner)

	subnets, err := s.fetchSubnets()
	if err != nil {
		return graph.Graph{}, errors.Wrap(err, "ipam")
	}

	gr := <-grc
	if gr.Error != nil {
		return gr.Graph, gr.Error
	}

	g := gr.Graph

	for _, mapping := range s.cfg.Mapping {
		g.AddAttribute(mapping.ID, mapping.Name, graph.AttributeTypeNode)

		reg, err := regexp.Compile(mapping.Match)
		if err != nil {
			return g, errors.Wrap(err, "ipam: mapping regexp")
		}

		for subnet, value := range subnets {
			if !reg.MatchString(value) {
				continue
			}

			value = reg.ReplaceAllString(value, mapping.Value)
			g = s.doSubnet(g, mapping.ID, subnet, value)
		}
	}

	return g, nil
}

func (s *IPAM) doSubnet(g graph.Graph, id graph.Attribute, subnet *net.IPNet, value string) graph.Graph {
	for i := range g.Nodes {
		ip := g.NodeAttr(i, skydive.AttributeIPAddr)
		if ip == "" {
			continue
		}

		if !subnet.Contains(net.ParseIP(ip)) {
			continue
		}

		g.SetNodeAttr(i, id, value)
	}

	return g
}

func (s *IPAM) fetchSubnets() (map[*net.IPNet]string, error) {
	res := map[*net.IPNet]string{}

	rows, err := s.db.QueryContext(
		context.Background(),
		`
			SELECT INET_NTOA(subnet), mask, description
			FROM subnets
			WHERE subnet IS NOT NULL
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subnet sql.NullString
		var maskStr string
		var desc string
		err := rows.Scan(
			&subnet,
			&maskStr,
			&desc,
		)
		if err != nil {
			return nil, err
		}

		if maskStr == "" {
			continue
		}
		if !subnet.Valid {
			continue
		}

		_, net, err := net.ParseCIDR(subnet.String + "/" + maskStr)
		if err != nil {
			return nil, err
		}
		res[net] = desc
	}

	return res, nil
}
