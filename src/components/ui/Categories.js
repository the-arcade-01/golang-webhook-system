import React from "react";

import Typography from "@mui/material/Typography";
import Card from "@mui/material/Card";
import CardMedia from "@mui/material/CardMedia";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";

import { categories } from "../../utils/Categories";

const Categories = () => {
  return (
    <div
      style={{
        backgroundColor: "#151515",
        color: "#f1f1f1",
        padding: "15px 30px",
      }}
    >
      <div style={{ display: "flex", gap: "5px", marginBottom: "10px" }}>
        <Typography sx={{ fontSize: "18px", fontWeight: "600" }}>
          Popular
        </Typography>
        <Typography
          sx={{
            fontSize: "18px",
            fontWeight: "600",
            color: "#bf94ff",
            cursor: "pointer",
            "&:hover": {
              textDecoration: "underline",
            },
          }}
        >
          categories
        </Typography>
      </div>
      <div
        style={{
          display: "flex",
          width: "100%",
          gap: "10px",
        }}
      >
        {categories.map((category) => {
          return (
            <Card elevation={0} sx={{ background: "#151515" }}>
              <CardMedia
                component="img"
                src={category.image}
                alt={category.name}
                sx={{ cursor: "pointer" }}
              />
              <CardContent sx={{ padding: "0" }}>
                <div
                  style={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    cursor: "pointer",
                  }}
                >
                  <Typography
                    sx={{
                      color: "#fff",
                      fontSize: "14px",
                      fontWeight: "600",
                      "&:hover": {
                        color: "#6441a5",
                      },
                    }}
                  >
                    {category.name}
                  </Typography>
                  <Button
                    sx={{
                      color: "#fff",
                      fontSize: "13px",
                      maxHeight: 30,
                      minWidth: 0,
                      "&:hover": {
                        backgroundColor: "hsla(0,0%,100%,0.2)",
                      },
                      marginTop: "5px",
                    }}
                  >
                    <i className="fi fi-br-menu-dots-vertical" />
                  </Button>
                </div>
                <Typography
                  sx={{
                    fontSize: "13px",
                    color: "hsla(0,0%,100%,0.7)",
                    "&:hover": {
                      color: "#6441a5",
                    },
                    cursor: "pointer",
                    marginTop: "-3px",
                  }}
                >
                  {category.viewers}
                </Typography>
                <div
                  style={{
                    display: "flex",
                    gap: "5px",
                    marginTop: "5px",
                  }}
                >
                  {category.tags.map((tag) => {
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
              </CardContent>
            </Card>
          );
        })}
      </div>
    </div>
  );
};

export default Categories;
