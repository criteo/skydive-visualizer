import React, { Component } from "react";
import { Attribute, GraphParams, Filters } from "../../services/api";
import "./Menu.css";
import FiltersSection from "./FilterSection";

type MenuProps = {
  attributes: Array<Attribute>;
  graphType: string;
  graphParams: GraphParams;
  go: (type: string, params: GraphParams) => void;
};

type MenuState = {
  graphType: string;
  graphParams: GraphParams;
};

export default class Menu extends Component<MenuProps, MenuState> {
  constructor(props: MenuProps) {
    super(props);

    this.state = {
      graphType: props.graphType,
      graphParams: props.graphParams
    };

    this.go = this.go.bind(this);
    this.onMaxValuesChanged = this.onMaxValuesChanged.bind(this);
    this.onFiltersChanged = this.onFiltersChanged.bind(this);
    this.onDimChange = this.onDimChange.bind(this);
  }

  onDimsChange(
    e: React.ChangeEvent<HTMLInputElement>,
    id: number,
    dir: string
  ) {
    let params = this.state.graphParams;
    if (e.target.checked) {
      params.dimensions[dir].push(id);
    } else {
      params.dimensions[dir] = params.dimensions[dir].filter(
        (e: number) => e !== id
      );
    }
    params.dimensions.src.sort((a, b) => a - b);
    params.dimensions.dst.sort((a, b) => b - a);

    this.setState({ graphParams: params });
  }

  onDimChange(e: React.ChangeEvent<HTMLSelectElement>) {
    let params = this.state.graphParams;
    params.dimension = parseInt(e.target.value);
    this.setState({ graphParams: params });
  }

  componentWillReceiveProps(props: MenuProps) {
    this.setState({
      graphType: props.graphType,
      graphParams: props.graphParams
    });
  }

  onVolumeChanged(e: React.ChangeEvent<HTMLSelectElement>) {
    let params = this.state.graphParams;
    params.volume = parseInt(e.target.value);
    this.setState({ graphParams: params });
  }

  onMaxValuesChanged(e: React.ChangeEvent<HTMLInputElement>) {
    let params = this.state.graphParams;
    params.maxValues = parseInt(e.target.value);
    this.setState({ graphParams: params });
  }

  onFiltersChanged(filters: Filters) {
    let params = this.state.graphParams;
    params.filters = filters;
    this.setState({ graphParams: params });
  }

  go() {
    let params = this.state.graphParams;
    if (params.volume === 0) {
      params.volume = this.props.attributes.filter(
        e => e.type === "edge"
      )[0].id;
    }

    this.props.go(this.state.graphType, params);
  }

  render() {
    var multiDim = (
      <>
        <ul id="dimensions-src">
          {this.props.attributes
            .filter(e => e.type === "node")
            .map(e => {
              return (
                <li key={"src" + e.id}>
                  <label>
                    <input
                      type="checkbox"
                      value={e.id}
                      onChange={ev => this.onDimsChange(ev, e.id, "src")}
                      checked={this.state.graphParams.dimensions.src.includes(
                        e.id
                      )}
                    />{" "}
                    ↦ {e.name}
                  </label>
                </li>
              );
            })}
        </ul>
        <ul id="dimensions-dst">
          {this.props.attributes
            .filter(e => e.type === "node")
            .reverse()
            .map(e => {
              return (
                <li key={"dst" + e.id}>
                  <label>
                    <input
                      type="checkbox"
                      value={e.id}
                      onChange={ev => this.onDimsChange(ev, e.id, "dst")}
                      checked={this.state.graphParams.dimensions.dst.includes(
                        e.id
                      )}
                    />{" "}
                    ↤ {e.name}
                  </label>
                </li>
              );
            })}
        </ul>
      </>
    );

    var singleDim = (
      <select
        value={this.state.graphParams.dimension}
        onChange={this.onDimChange}
      >
        {this.props.attributes
          .filter(e => e.type === "node")
          .map(e => (
            <option value={e.id} key={e.id}>
              {e.name}
            </option>
          ))}
      </select>
    );

    return (
      <div className="Menu">
        <div className="Section">
          <div className="Title">Graph type</div>
          <select
            id="graph-type"
            value={this.state.graphType}
            onChange={e => this.setState({ graphType: e.target.value })}
          >
            <option value="sankey">Sankey</option>
            <option value="table">Table</option>
            <option value="graph">Graph</option>
          </select>
        </div>
        <div className="Section">
          <div className="Title">Volume</div>
          <select
            id="volume"
            onChange={this.onVolumeChanged}
            value={this.state.graphParams.volume}
          >
            {this.props.attributes
              .filter(e => e.type === "edge")
              .map(e => {
                return (
                  <option key={e.id} value={e.id}>
                    {e.name}
                  </option>
                );
              })}
          </select>
        </div>
        <div className="Section">
          <div className="Title">Dimensions</div>
          {this.state.graphType === "graph" ? singleDim : multiDim}
        </div>

        <FiltersSection
          attributes={this.props.attributes}
          filters={this.state.graphParams.filters}
          onChange={this.onFiltersChanged}
        />

        <div className="Section">
          <div className="Title">Max values</div>
          <input
            type="number"
            value={this.state.graphParams.maxValues}
            onChange={this.onMaxValuesChanged}
          />
        </div>

        <div className="Section">
          <button className="Go" onClick={this.go}>
            Go
          </button>
        </div>
      </div>
    );
  }
}
