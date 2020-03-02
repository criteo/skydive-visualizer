import React, { useRef } from "react";
import * as d3 from "d3";
import * as d3Sankey from "d3-sankey";

const Rect = ({ index, x0, x1, y0, y1, name, value, length, colors, size }) => {
  return (
    <>
      <rect
        x={x0}
        y={y0}
        width={x1 - x0}
        height={y1 - y0}
        fill={colors(index / length)}
        stroke={"#222"}
        strokeOpacity={1}
        data-index={index}
      />
      <text
        x={x0 < size.width / 2 ? x1 + 6 : x0 - 6}
        y={(y1 + y0) / 2}
        style={{
          fill: "black",
          alignmentBaseline: "middle",
          fontSize: 12,
          textAnchor: x0 < size.width / 2 ? "start" : "end",
          pointerEvents: "none",
          userSelect: "none"
        }}
      >
        {name}
      </text>
    </>
  );
};

const Link = ({ data, width, onMouseEnter, onMouseLeave }) => {
  const link = d3Sankey.sankeyLinkHorizontal();

  return (
    <path
      id={`link-${data.index}`}
      d={link(data)}
      fill={"none"}
      stroke={"#555"}
      strokeOpacity={0.3}
      strokeWidth={width}
      onMouseEnter={e => onMouseEnter(data.index)}
      onMouseLeave={e => onMouseLeave(data.index)}
    >
      <title>
        {data.source.name} → {data.target.name} : {data.value}{" "}
      </title>
    </path>
  );
};

const Sankey = props => {
  const graph = useRef(null);

  const size = {
    width: props.width,
    height: props.height
  };

  const colors = d3.scaleOrdinal(d3.schemeCategory10);

  const sankey = d3Sankey
    .sankey()
    .nodeWidth(10)
    .nodePadding(15)
    .extent([
      [1, 10],
      [size.width - 2, size.height - 20]
    ]);

  // custom links mapping
  let mappedLinks = [];
  let linkMap = {};
  for (let i in props.data.links) {
    let lastNode = false;
    let l = [];
    for (let j in props.data.links[i].nodes) {
      if (lastNode !== false) {
        l.push({
          source: lastNode,
          target: props.data.links[i].nodes[j],
          value: props.data.links[i].value
        });
      }
      lastNode = props.data.links[i].nodes[j];
    }
    let m = [];
    for (let i in l) {
      mappedLinks.push(l[i]);
      m.push(mappedLinks.length - 1);
    }
    for (let i in m) {
      linkMap[m[i]] = m;
    }
  }

  function highlightLink(i) {
    for (let j in linkMap[i]) {
      d3.select("#link-" + linkMap[i][j]).style("stroke-opacity", 0.6);
    }
  }

  function unHighlightLink(i) {
    for (let j in linkMap[i]) {
      d3.select("#link-" + linkMap[i][j]).style("stroke-opacity", 0.3);
    }
  }

  if (props.data) {
    graph.current = sankey({
      links: mappedLinks,
      nodes: props.data.nodes
    });
    const { links, nodes } = graph.current;

    return (
      <svg width={size.width} height={size.height}>
        <g>
          {links.map((d, i) => (
            <Link
              key={d.index}
              data={d}
              width={d.width}
              length={nodes.length}
              colors={colors}
              onMouseEnter={highlightLink}
              onMouseLeave={unHighlightLink}
            />
          ))}
        </g>
        <g>
          {nodes.map((d, i) => (
            <Rect
              key={d.index}
              index={d.index}
              x0={d.x0}
              x1={d.x1}
              y0={d.y0}
              y1={d.y1}
              name={d.name}
              value={d.value}
              length={nodes.length}
              colors={colors}
              size={size}
            />
          ))}
        </g>
      </svg>
    );
  }

  return <div>Loading</div>;
};

export default Sankey;
