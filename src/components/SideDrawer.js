import React from "react";

import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import List from "@mui/material/List";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import ListItem from "@mui/material/ListItem";
import TextField from "@mui/material/TextField";

import { followedChannels, recommendedChannels } from "../utils/SideDrawer";

const drawerWidth = 240;

const SideDrawer = () => {
  return (
    <>
      <Drawer
        variant="permanent"
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          [`& .MuiDrawer-paper`]: {
            width: drawerWidth,
            boxSizing: "border-box",
            backgroundColor: "#262626",
            color: "#f1f1f1",
            padding: "0 12px",
          },
        }}
      >
        <Toolbar />
        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            marginTop: "-5px",
          }}
        >
          <Typography
            sx={{
              textTransform: "uppercase",
              fontSize: "13px",
              fontWeight: "600",
            }}
          >
            Followed Channels
          </Typography>
          <Button
            sx={{
              color: "#fff",
              fontSize: "15px",
              maxHeight: 35,
              minWidth: 0,
              "&:hover": {
                backgroundColor: "hsla(0,0%,100%,0.2)",
              },
              transform: "scaleX(-1)",
            }}
          >
            <i className="fi fi-br-sign-out" style={{ paddingTop: "5px" }} />
          </Button>
        </div>
        <Box sx={{ marginTop: "-15px" }}>
          <List>
            {followedChannels.map((channel) => {
              return (
                <ListItem
                  sx={{
                    color: "#f1f1f1",
                    marginLeft: "-15px",
                    marginBottom: "-10px",
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    "&:hover": {
                      cursor: "pointer",
                      backgroundColor: "#303030",
                    },
                    width: 240,
                    filter: channel.status ? null : "brightness(70%)",
                  }}
                >
                  <div
                    style={{
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <img
                      src={channel.image}
                      alt={channel.name}
                      style={{
                        borderRadius: "50px",
                        height: 32,
                        width: 32,
                      }}
                    />
                    <div style={{ paddingLeft: "8px" }}>
                      <Typography sx={{ fontSize: "14px", fontWeight: "600" }}>
                        {channel.name}
                      </Typography>
                      <Typography sx={{ fontSize: "12px", color: "grey" }}>
                        {channel.description}
                      </Typography>
                    </div>
                  </div>
                  <Typography
                    sx={{
                      color: channel.status ? "#ff0000" : "#f1f1f1",
                      fontSize: channel.status ? "14px" : "13px",
                      fontWeight: channel.status ? "800" : "400",
                      marginTop: "-15px",
                      paddingLeft: "5px",
                    }}
                  >
                    {channel.watch}
                  </Typography>
                </ListItem>
              );
            })}
          </List>
        </Box>
        <Divider />
        <Typography
          sx={{
            textTransform: "uppercase",
            fontSize: "13px",
            fontWeight: "600",
            paddingTop: "10px",
          }}
        >
          Recommended Channels
        </Typography>
        <Box sx={{ marginTop: "-10px" }}>
          <List>
            {recommendedChannels.map((channel) => {
              return (
                <ListItem
                  sx={{
                    color: "#f1f1f1",
                    marginLeft: "-15px",
                    marginBottom: "-10px",
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    "&:hover": {
                      cursor: "pointer",
                      backgroundColor: "#303030",
                    },
                    width: 240,
                    filter: channel.status ? null : "brightness(70%)",
                  }}
                >
                  <div
                    style={{
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <img
                      src={channel.image}
                      alt={channel.name}
                      style={{
                        borderRadius: "50px",
                        height: 32,
                        width: 32,
                      }}
                    />
                    <div style={{ paddingLeft: "8px" }}>
                      <Typography sx={{ fontSize: "14px", fontWeight: "600" }}>
                        {channel.name}
                      </Typography>
                      <Typography sx={{ fontSize: "12px", color: "grey" }}>
                        {channel.description}
                      </Typography>
                    </div>
                  </div>
                  <Typography
                    sx={{
                      color: channel.status ? "red" : "#f1f1f1",
                      fontSize: channel.status ? "14px" : "13px",
                      fontWeight: channel.status ? "800" : "400",
                      marginTop: "-15px",
                      paddingLeft: "5px",
                    }}
                  >
                    {channel.watch}
                  </Typography>
                </ListItem>
              );
            })}
          </List>
        </Box>
        <Divider />
        <TextField
          placeholder="Search to Add Friends"
          autoComplete="off"
          variant="outlined"
          size="small"
          sx={{
            backgroundColor: "hsla(0,0%,100%,0.2)",
            borderRadius: "5px",
            width: "100%",
            marginTop: "80px",
            color: "#fff",
          }}
          InputProps={{
            style: {
              height: 35,
              color: "#fff",
              fontSize: "13.5px",
            },
            startAdornment: (
              <i
                className="fi fi-br-search"
                style={{ margin: "5px 10px 0px 0px" }}
              />
            ),
          }}
        />
      </Drawer>
    </>
  );
};

export default SideDrawer;
