import React from "react";
import Loader from "react-loader-spinner";
import "./TopBar.css";

type Props = {
  loading: boolean;
};

function TopBar(props: Props) {
  return (
    <div className="Top-bar">
      Skydive Visualizer
      <Loader
        className="Loader"
        type="Audio"
        color="#55dbcb"
        height={30}
        width={30}
        visible={props.loading}
      />
    </div>
  );
}

export default TopBar;
