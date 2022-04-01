import React from "react";

import Typography from "@mui/material/Typography";
import Card from "@mui/material/Card";
import CardMedia from "@mui/material/CardMedia";
import CardContent from "@mui/material/CardContent";
import Avatar from "@mui/material/Avatar";

import { suggestChannels } from "../../utils/Suggestion";

const Suggestion = () => {
  return (
    <div
      style={{
        backgroundColor: "#151515",
        color: "#f1f1f1",
        padding: "15px 30px",
        marginBottom: "-15px",
      }}
    >
      <Typography
        sx={{ fontSize: "18px", fontWeight: "600", marginBottom: "10px" }}
      >
        Live channels we think you’ll like
      </Typography>
      <div
        style={{
          display: "flex",
          width: "100%",
          gap: "10px",
        }}
      >
        {suggestChannels.map((channel) => {
          return (
            <Card elevation={0} sx={{ background: "#151515" }}>
              <CardMedia
                component="img"
                src={channel.thumbnail}
                alt={channel.name}
                sx={{ cursor: "pointer" }}
              />
              <CardContent
                sx={{
                  display: "flex",
                  alignItems: "center",
                  gap: "10px",
                  padding: "0px 10px",
                }}
              >
                <Avatar
                  src={channel.image}
                  sx={{ cursor: "pointer", marginTop: "-15px" }}
                />
                <div>
                  <Typography
                    sx={{
                      fontWeight: "600",
                      color: "#f1f1f1",
                      fontSize: "15px",
                      paddingTop: "10px",
                      "&:hover": {
                        color: "#6441a5",
                      },
                      cursor: "pointer",
                    }}
                  >
                    {channel.title}
                  </Typography>
                  <Typography
                    sx={{
                      cursor: "pointer",
                      color: "#f1f1f1",
                      fontSize: "14px",
                      fontWeight: "500",
                      "&:hover": {
                        color: "#6441a5",
                      },
                    }}
                  >
                    {channel.name}
                  </Typography>
                  <div
                    style={{
                      display: "flex",
                      gap: "5px",
                      marginTop: "5px",
                    }}
                  >
                    {channel.tags.map((tag) => {
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
                </div>
              </CardContent>
            </Card>
          );
        })}
      </div>
    </div>
  );
};

export default Suggestion;
