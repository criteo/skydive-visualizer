import React, { Component } from "react";
import { Attribute, Filters } from "../../services/api";
import { default as FilterCpnt } from "./Filter";
import "./FilterSection.css";

type FiltersSectionProps = {
  attributes: Array<Attribute>;
  filters: Filters;
  onChange: (f: Filters) => void;
};

export default class FiltersSection extends Component<FiltersSectionProps, {}> {
  constructor(props: FiltersSectionProps) {
    super(props);

    this.onFilterChanged = this.onFilterChanged.bind(this);
    this.onFilterRemoved = this.onFilterRemoved.bind(this);
    this.add = this.add.bind(this);
  }

  add() {
    let filters = this.props.filters;
    filters.src[0] = "";
    this.props.onChange(filters);
  }

  onFilterChanged(
    oldDir: string,
    oldKey: number,
    newDir: string,
    newKey: number,
    value: string
  ) {
    let filters = this.props.filters;
    delete filters[oldDir][oldKey];
    filters[newDir][newKey] = value;

    this.props.onChange(filters);
  }

  onFilterRemoved(dir: string, key: number) {
    let filters = this.props.filters;
    delete filters[dir][key];

    this.props.onChange(filters);
  }

  render() {
    let filters: Array<JSX.Element> = [];
    Object.keys(this.props.filters.src).forEach(k => {
      let key = parseInt(k);
      let value: string = this.props.filters.src[key];
      filters.push(
        <FilterCpnt
          key={-key}
          attributes={this.props.attributes}
          keyId={key}
          dir="src"
          value={value}
          onChange={this.onFilterChanged}
          onRemove={this.onFilterRemoved}
        />
      );
    });
    Object.keys(this.props.filters.dst).forEach(k => {
      let key = parseInt(k);
      let value: string = this.props.filters.dst[key];
      filters.push(
        <FilterCpnt
          key={key}
          attributes={this.props.attributes}
          keyId={key}
          dir="dst"
          value={value}
          onChange={this.onFilterChanged}
          onRemove={this.onFilterRemoved}
        />
      );
    });

    return (
      <div className="Section">
        <div className="Title">
          <button className="Add-button" onClick={this.add}>
            Add
          </button>
          Filter
        </div>
        <ul>{filters}</ul>
      </div>
    );
  }
}
