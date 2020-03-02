import React, { Component, createRef } from "react";
import { TableData, SankeyData, GraphData } from "../services/api";
import Sankey from "./Sankey";
import Table from "./Table";
import { Graph as GraphViz } from "react-d3-graph";

type GraphProps = {
  type: string;
  data: any;
};

export default class Graph extends Component<GraphProps, {}> {
  private container = createRef<HTMLDivElement>();

  constructor(props: GraphProps) {
    super(props);
    this.state = {
      data: null
    };
  }

  render() {
    let width = this.container.current?.offsetWidth || 800;
    let height = this.container.current?.offsetHeight || 600;

    let graphEl: any;
    if (this.props.data) {
      switch (this.props.type) {
        case "sankey":
          if ((this.props.data as SankeyData).nodes) {
            graphEl = (
              <Sankey
                data={this.props.data as SankeyData}
                width={width}
                height={height}
              />
            );
          }
          break;
        case "table":
          if ((this.props.data as TableData).rows) {
            graphEl = <Table data={this.props.data as TableData} />;
          }
          break;

        case "graph":
          if ((this.props.data as GraphData).nodes) {
            graphEl = (
              <GraphViz
                id="grapviz-graph"
                data={this.props.data}
                config={{
                  width: width,
                  height: height,
                  directed: true,
                  nodeHighlightBehavior: true,
                  node: {
                    color: "lightgreen",
                    size: 120,
                    highlightStrokeColor: "blue"
                  },
                  link: {
                    highlightColor: "lightblue"
                  },
                  d3: {
                    alphaTarget: 0.05,
                    gravity: -100,
                    linkLength: 100,
                    linkStrength: 1,
                    disableLinkForce: false
                  }
                }}
              />
            );
          }
          break;
      }
    }

    return (
      <div ref={this.container} style={{ height: "100%" }}>
        {graphEl}
      </div>
    );
  }
}
