import React, { useState } from "react";

import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Avatar from "@mui/material/Avatar";
import Typography from "@mui/material/Typography";

import YoutubeEmbed from "./YoutubeEmbed";

import { carousal } from "../../utils/Carousal";

const Carousal = () => {
  const [current, setCurrent] = useState(0);

  const length = carousal.length;

  const prevSlide = () => {
    setCurrent(current === 0 ? length - 1 : current - 1);
  };
  const nextSlide = () => {
    setCurrent(current === length - 1 ? 0 : current + 1);
  };
  return (
    <div
      style={{
        marginTop: "50px",
        padding: "30px 20px",
        backgroundColor: "#151515",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Button
        sx={{
          position: "absolute",
          color: "#fff",
          fontSize: "18px",
          maxHeight: 35,
          minWidth: 0,
          "&:hover": {
            backgroundColor: "hsla(0,0%,100%,0.2)",
          },
          zIndex: 10,
          userSelect: "none",
          top: "25%",
          left: "270px",
        }}
        onClick={prevSlide}
      >
        <i className="fi fi-sr-angle-left" style={{ paddingTop: "5px" }} />
      </Button>
      <Button
        sx={{
          position: "absolute",
          color: "#fff",
          fontSize: "18px",
          maxHeight: 35,
          minWidth: 0,
          "&:hover": {
            backgroundColor: "hsla(0,0%,100%,0.2)",
          },
          zIndex: 10,
          userSelect: "none",
          top: "25%",
          right: "30px",
        }}
        onClick={nextSlide}
      >
        <i className="fi fi-sr-angle-right" style={{ paddingTop: "5px" }} />
      </Button>
      {carousal.map((data, index) => {
        return (
          <Card sx={{ backgroundColor: "#262626", height: "300px" }}>
            {index === current ? (
              <div style={{ display: "flex" }}>
                <YoutubeEmbed embedId={data.youtube} />
                <CardContent sx={{ padding: "15px", width: 230 }}>
                  <div style={{ display: "flex", gap: "10px" }}>
                    <Avatar
                      src={data.image}
                      sx={{ cursor: "pointer", width: 55, height: 55 }}
                    />
                    <div>
                      <Typography
                        sx={{
                          fontSize: "14px",
                          fontWeight: "600",
                          cursor: "pointer",
                          color: "rgb(191, 148, 255)",
                        }}
                      >
                        {data.name}
                      </Typography>
                      <Typography
                        sx={{
                          fontSize: "13px",
                          color: "rgb(191, 148, 255)",
                          cursor: "pointer",
                        }}
                      >
                        {data.title.length > 15
                          ? `${data.title.substr(0, 15)}...`
                          : data.title}
                      </Typography>
                    </div>
                  </div>
                  <div
                    style={{
                      display: "flex",
                      gap: "5px",
                      marginTop: "10px",
                    }}
                  >
                    {data.tags.map((tag) => {
                      return (
                        <div
                          style={{
                            backgroundColor: "hsla(0,0%,100%,0.2)",
                            padding: "2px 10px",
                            borderTopLeftRadius: "15px",
                            borderBottomLeftRadius: "15px",
                            borderTopRightRadius: "15px",
                            borderBottomRightRadius: "15px",
                            cursor: "pointer",
                          }}
                        >
                          <Typography
                            sx={{
                              fontSize: "12px",
                              color: "hsla(0,0%,100%,0.8)",
                              fontWeight: "600",
                            }}
                          >
                            {tag}
                          </Typography>
                        </div>
                      );
                    })}
                  </div>
                  <div style={{ paddingTop: "10px" }}>
                    <Typography sx={{ fontSize: "13px", color: "#f1f1f1" }}>
                      {data.description.length > 210
                        ? `${data.description.substr(0, 210)} ...`
                        : data.description}
                    </Typography>
                  </div>
                </CardContent>
              </div>
            ) : null}
          </Card>
        );
      })}
    </div>
  );
};

export default Carousal;
