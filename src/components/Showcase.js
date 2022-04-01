import React from "react";

import SideDrawer from "./ui/SideDrawer";
import Carousal from "./ui/Carousal";
import Categories from "./ui/Categories";
import Suggestion from "./ui/Suggestion";

const Showcase = () => {
  return (
    <div style={{ display: "flex" }}>
      <SideDrawer />
      <div style={{ display: "flex", flexDirection: "column", width: "100%" }}>
        <Carousal />
        <Suggestion />
        <Categories />
      </div>
    </div>
  );
};

export default Showcase;
