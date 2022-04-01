import React from "react";

import SideDrawer from "./ui/SideDrawer";
import Carousal from "./ui/Carousal";
import Categories from "./ui/Categories";

const Showcase = () => {
  return (
    <div style={{ display: "flex" }}>
      <SideDrawer />
      <div style={{ display: "flex", flexDirection: "column", width: "100%" }}>
        <Carousal />
        <Categories />
      </div>
    </div>
  );
};

export default Showcase;
