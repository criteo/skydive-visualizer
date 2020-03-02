import Axios from "axios";

export type Attribute = {
  id: number;
  name: string;
  type: string;
};

export type Dimensions = {
  src: Array<number>;
  dst: Array<number>;
};

export type Filters = {
  src: {};
  dst: {};
};

export type GraphParams = {
  maxValues: number;
  volume: number;
  dimensions: Dimensions;
  dimension: number;
  filters: Filters;
};

export type SankeyNode = {
  name: string;
};

export type SankeyLink = {
  value: number;
  nodes: Array<number>;
};

export type SankeyData = {
  nodes: Array<SankeyNode>;
  links: Array<SankeyLink>;
};

export type GraphNode = {
  id: string;
};

export type GraphLink = {
  source: string;
  target: string;
};

export type GraphData = {
  nodes: Array<GraphNode>;
  links: Array<GraphLink>;
};

export type TableData = {
  header: Array<string>;
  rows: Array<Array<any>>;
};

export class Api {
  static async getAttributes(): Promise<Array<Attribute>> {
    let res = await Axios.get("/attributes");
    return res.data;
  }

  static async getAttributeValue(attr: number): Promise<Array<string>> {
    let res = await Axios.get(`/attributes/${attr}/values`);
    return res.data;
  }

  static async getSankey(params: GraphParams): Promise<SankeyData> {
    let res = await Axios.post(`/graph/sankey`, params);
    return res.data;
  }

  static async getTable(params: GraphParams): Promise<TableData> {
    let res = await Axios.post(`/graph/table`, params);
    return res.data;
  }

  static async getGraph(params: GraphParams): Promise<GraphData> {
    let res = await Axios.post(`/graph/graph`, params);
    return res.data;
  }
}
