import React, { Component } from "react";
import { Attribute, Api } from "../../services/api";
import "./Filter.css";

type FilterProps = {
  attributes: Array<Attribute>;
  dir: string;
  keyId: number;
  value: string;
  onChange: (
    oldDir: string,
    oldId: number,
    newDir: string,
    newKey: number,
    value: string
  ) => void;
  onRemove: (dir: string, id: number) => void;
};

type FilterState = {
  attrValues: Array<string>;
};

export default class Filter extends Component<FilterProps, FilterState> {
  constructor(props: FilterProps) {
    super(props);

    this.state = {
      attrValues: []
    };
  }

  async componentWillMount() {
    let values = await Api.getAttributeValue(this.props.keyId);
    this.setState({
      attrValues: values
    });
  }

  render() {
    return (
      <li>
        <div className="Filter-key">
          <select
            onChange={e => {
              let newKey = parseInt(e.target.value);
              let newDir = newKey <= 0 ? "src" : "dst";
              this.props.onChange(
                this.props.dir,
                this.props.keyId,
                newDir,
                Math.abs(newKey),
                this.props.value
              );
            }}
            value={
              this.props.dir === "src" ? -this.props.keyId : this.props.keyId
            }
          >
            <option value="0">Select...</option>
            {this.props.attributes
              .filter(e => e.type === "node")
              .map(e => {
                return (
                  <option key={-e.id} value={-e.id}>
                    ↦ {e.name}
                  </option>
                );
              })}
            {this.props.attributes
              .filter(e => e.type === "node")
              .reverse()
              .map(e => {
                return (
                  <option key={e.id} value={e.id}>
                    ↤ {e.name}
                  </option>
                );
              })}
          </select>
          <button
            onClick={e => this.props.onRemove(this.props.dir, this.props.keyId)}
          >
            ✕
          </button>
        </div>
        <select
          onChange={e =>
            this.props.onChange(
              this.props.dir,
              this.props.keyId,
              this.props.dir,
              this.props.keyId,
              e.target.value
            )
          }
          value={this.props.value}
        >
          <option key={0} value="">
            {"<empty>"}
          </option>
          {this.state.attrValues.map((e, i) => {
            return (
              <option key={i} value={e}>
                {e}
              </option>
            );
          })}
        </select>
      </li>
    );
  }
}
