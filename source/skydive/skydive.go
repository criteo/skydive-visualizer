package skydive

import (
	"context"
	"net"
	"network/skydive-visualizer-go/graph"
	"network/skydive-visualizer-go/skydive"
	"strconv"

	"github.com/prometheus/common/log"
)

type nodeSocket struct {
	Node   skydive.Node
	Socket skydive.Socket
}

type Skydive struct {
	client *skydive.Skydive
}

func NewSkydive(client *skydive.Skydive) *Skydive {
	return &Skydive{client}
}

func (s *Skydive) Fetch() (graph.Graph, error) {
	log.Info("skydive source: starting")
	defer log.Info("skydive source: done")

	g := graph.Graph{}
	g.AddAttribute(AttributeConnections, "Connections", graph.AttributeTypeEdge)
	g.AddAttribute(AttributeIPAddr, "IP address", graph.AttributeTypeNode)
	g.AddAttribute(AttributeProcess, "Process", graph.AttributeTypeNode)
	g.AddAttribute(AttributePortNum, "Port number", graph.AttributeTypeNode)

	servers := map[string]skydive.Socket{}

	hostNodes, err := s.client.LookupNodes(context.Background(), "G.V().Has('Type', 'host').HasKey('Sockets')")
	if err != nil {
		return g, err
	}

	for _, node := range hostNodes {
		for _, socket := range node.Metadata.Sockets {
			if !isListenSocket(socket) {
				continue
			}

			servers[hostPort(node.Host, socket.LocalPort)] = socket
		}
	}

	remotes := map[string]nodeSocket{}

	for _, node := range hostNodes {
		for _, socket := range node.Metadata.Sockets {
			if isListenSocket(socket) {
				continue
			}

			remoteNodeName := socket.RemoteAddress // TODO resolve
			localAddr := hostPort(node.Host, socket.LocalPort)
			remoteAddr := hostPort(remoteNodeName, socket.RemotePort)

			if isLoopBack(socket.LocalAddress) && isLoopBack(socket.RemoteAddress) {
				continue
			}

			// this node is the server. we save this information in case we do not have an
			// agent on the client
			if _, ok := servers[localAddr]; ok {
				remotes[remoteAddr] = nodeSocket{node, socket}
				continue
			}

			// this node is the client and we have an identified server
			if _, ok := servers[remoteAddr]; ok {
				fromNode := g.CreateNode(map[graph.Attribute]string{
					AttributeIPAddr:  socket.LocalAddress,
					AttributePortNum: strconv.Itoa(socket.LocalPort),
					AttributeProcess: socket.Name,
				})

				toNode := g.CreateNode(map[graph.Attribute]string{
					AttributeIPAddr:  socket.RemoteAddress,
					AttributePortNum: strconv.Itoa(socket.RemotePort),
					AttributeProcess: servers[remoteAddr].Name,
				})

				g.AddEdge(graph.Edge{
					FromNode: fromNode,
					ToNode:   toNode,
					Attrs: graph.EdgeAttributes{
						AttributeConnections: 1,
					},
				})

				// if we had a saved connection from the server side, remove it
				// as we found it on a client
				delete(remotes, localAddr)
				continue
			}

			// we are the client to an unidentified server
			fromNode := g.CreateNode(map[graph.Attribute]string{
				AttributeIPAddr:  socket.LocalAddress,
				AttributePortNum: strconv.Itoa(socket.LocalPort),
				AttributeProcess: socket.Name,
			})

			toNode := g.CreateNode(map[graph.Attribute]string{
				AttributeIPAddr:  socket.RemoteAddress,
				AttributePortNum: strconv.Itoa(socket.RemotePort),
			})

			g.AddEdge(graph.Edge{
				FromNode: fromNode,
				ToNode:   toNode,
				Attrs: graph.EdgeAttributes{
					AttributeConnections: 1,
				},
			})
		}
	}

	// add the unidentified clients to servers we know
	for _, nodeSocket := range remotes {
		node := nodeSocket.Node
		socket := nodeSocket.Socket

		fromNode := g.CreateNode(map[graph.Attribute]string{
			AttributeIPAddr:  socket.RemoteAddress,
			AttributePortNum: strconv.Itoa(socket.RemotePort),
		})

		toSocket := servers[hostPort(node.Host, socket.LocalPort)]
		toNode := g.CreateNode(map[graph.Attribute]string{
			AttributeIPAddr:  socket.LocalAddress,
			AttributePortNum: strconv.Itoa(socket.LocalPort),
			AttributeProcess: toSocket.Name,
		})

		g.AddEdge(graph.Edge{
			FromNode: fromNode,
			ToNode:   toNode,
			Attrs: graph.EdgeAttributes{
				AttributeConnections: 1,
			},
		})
	}

	return g, nil
}

func isListenSocket(socket skydive.Socket) bool {
	return (socket.State == "LISTEN" || (socket.Protocol == "UDP" &&
		socket.State == "CLOSE" &&
		(socket.RemoteAddress == "0.0.0.0" ||
			socket.RemoteAddress == "::")))

}

func hostPort(host string, port int) string {
	return host + ":" + strconv.Itoa(port)
}

func isLoopBack(ip string) bool {
	return net.ParseIP(ip).IsLoopback()
}
