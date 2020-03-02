package chef

import (
	"fmt"
	"network/skydive-visualizer-go/graph"
	"network/skydive-visualizer-go/source"
	"network/skydive-visualizer-go/source/skydive"
	"strings"
	"sync"

	"github.com/go-chef/chef"
	log "github.com/sirupsen/logrus"
)

type Chef struct {
	cfg     *Config
	clients []*chef.Client
	inner   source.Source
}

func New(inner source.Source, cfg *Config) (*Chef, error) {
	s := &Chef{
		cfg:   cfg,
		inner: inner,
	}

	for _, url := range cfg.Servers {
		// build a client
		client, err := chef.NewClient(&chef.Config{
			Name:    cfg.User,
			Key:     cfg.Key,
			BaseURL: url,
			SkipSSL: true,
			Timeout: 10,
		})
		if err != nil {
			return nil, err
		}
		s.clients = append(s.clients, client)
	}

	return s, nil
}

func (s *Chef) Fetch() (graph.Graph, error) {
	log.Info("chef source: starting")
	defer log.Info("chef source: done")

	grc := source.WillFetch(s.inner)

	type task struct {
		clientID int
		node     string
	}

	nodesAttrs := map[string]map[graph.Attribute]string{}
	var nodesAttrsLock sync.Mutex

	log.Info("fetching chef data")

	tasks := make(chan task)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for task := range tasks {
				ip, attrs, err := s.nodeAttrs(task.node, task.clientID)
				if err != nil {
					log.Warn(err)
					continue
				}
				if ip == "" {
					continue
				}

				nodesAttrsLock.Lock()
				nodesAttrs[ip] = attrs
				nodesAttrsLock.Unlock()
			}
		}()
	}

	for i, client := range s.clients {
		nodes, err := client.Nodes.List()
		if err != nil {
			log.Error(err)
			continue
		}

		for node := range nodes {
			tasks <- task{i, node}
		}
	}

	close(tasks)
	wg.Wait()

	log.Info("done fetching chef data")

	gr := <-grc
	if gr.Error != nil {
		return gr.Graph, gr.Error
	}
	g := gr.Graph

	for _, a := range s.cfg.AttrsMapping {
		g.AddAttribute(a.ID, a.Name, graph.AttributeTypeNode)
	}
	g.AddAttribute(AttributeChefPolicy, "Chef policy", graph.AttributeTypeNode)

	for k := range g.Nodes {
		ip := g.NodeAttr(k, skydive.AttributeIPAddr)
		if ip == "" {
			continue
		}

		attrs, ok := nodesAttrs[ip]
		if !ok {
			log.Warnf("no chef data found for %s", ip)
			continue
		}

		for ak, av := range attrs {
			g.SetNodeAttr(k, ak, av)
		}
	}

	return g, nil
}

func (s *Chef) nodeAttrs(node string, clientID int) (string, map[graph.Attribute]string, error) {
	details, err := s.clients[clientID].Nodes.Get(node)
	if err != nil {
		return "", nil, err
	}

	ip, ok := details.AutomaticAttributes["ipaddress"].(string)
	if !ok {
		return "", nil, nil
	}

	res := map[graph.Attribute]string{
		AttributeChefPolicy: details.PolicyName,
	}

	for _, a := range s.cfg.AttrsMapping {
		v, ok := getChefAttr(details, a.Key)
		if ok {
			res[a.ID] = v
		}
	}

	return ip, res, nil
}

func getChefAttr(node chef.Node, key string) (string, bool) {
	attrSources := []map[string]interface{}{
		node.AutomaticAttributes,
		node.NormalAttributes,
		node.OverrideAttributes,
		node.DefaultAttributes,
	}

	keyParts := strings.Split(key, ".")

	for _, source := range attrSources {
		base := source
		for _, part := range keyParts[:len(keyParts)-1] {
			if v, ok := base[part].(map[string]interface{}); ok {
				base = v
			} else {
				base = nil
				break
			}
		}
		if base != nil {
			if v, ok := base[keyParts[len(keyParts)-1]]; ok {
				return fmt.Sprint(v), true
			}
		}
	}

	return "", false
}
