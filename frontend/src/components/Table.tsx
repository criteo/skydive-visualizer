import React, { Component } from "react";
import "./Table.css";
import { TableData } from "../services/api";
import SmartDataTable from "react-smart-data-table";

type TableProps = {
  data: TableData;
};

export default class Table extends Component<TableProps, {}> {
  render() {
    if (!this.props.data.header) {
      return null;
    }

    let data: Array<{}> = [];
    for (let i in this.props.data.rows) {
      let row = {};
      for (let j in this.props.data.rows[i]) {
        row[this.props.data.header[j]] = this.props.data.rows[i][j];
      }
      data.push(row);
    }

    return <SmartDataTable data={data} className="Table" sortable withFooter />;
  }
}
