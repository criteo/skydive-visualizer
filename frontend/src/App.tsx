import React, { Component } from "react";
import TopBar from "./components/TopBar";
import Menu from "./components/menu/Menu";
import { Api, Attribute, GraphParams } from "./services/api";
import "./App.css";
import Graph from "./components/Graph";

type AppState = {
  loading: boolean;
  attributes: Array<Attribute>;
  graphType: string;
  graphParams: GraphParams;
  graphData: any;
};

export class App extends Component<{}, AppState> {
  constructor(props: any) {
    super(props);
    this.state = {
      loading: false,
      attributes: [],

      graphType: "sankey",
      graphParams: {
        maxValues: 50,
        volume: 0,
        dimension: 0,
        dimensions: {
          src: [],
          dst: []
        },
        filters: {
          src: new Map<number, string>(),
          dst: new Map<number, string>()
        }
      },
      graphData: null
    };

    this.fetchGraph = this.fetchGraph.bind(this);
  }

  async fetchGraph(type: string, params: GraphParams) {
    let urlParam = JSON.stringify({
      type: type,
      params: params
    });
    window.location.hash = "#" + btoa(urlParam);

    this.setState({
      graphType: type,
      graphParams: params
    });
    this.setState({ loading: true });

    switch (type) {
      case "sankey":
        this.setState({
          graphData: await Api.getSankey(params)
        });
        break;

      case "table":
        this.setState({
          graphData: await Api.getTable(params)
        });
        break;

      case "graph":
        this.setState({
          graphData: await Api.getGraph(params)
        });
        break;
    }

    this.setState({ loading: false });
  }

  async componentWillMount() {
    this.setState({ loading: true });
    let attrs = await Api.getAttributes();
    this.setState({ attributes: attrs, loading: false });

    if (window.location.hash.length > 1) {
      let params: {
        type: string;
        params: GraphParams;
      } = JSON.parse(atob(window.location.hash.slice(1)));

      this.fetchGraph(params.type, params.params);
    }
  }

  render() {
    return (
      <div className="App">
        <TopBar loading={this.state.loading} />
        <Menu
          attributes={this.state.attributes}
          graphType={this.state.graphType}
          graphParams={this.state.graphParams}
          go={this.fetchGraph}
        />
        <div className="Graph-container">
          <Graph type={this.state.graphType} data={this.state.graphData} />
        </div>
      </div>
    );
  }
}

export default App;
