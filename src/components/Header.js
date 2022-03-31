import React from "react";

import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Avatar from "@mui/material/Avatar";

const Header = () => {
  return (
    <>
      <AppBar
        sx={{
          height: 50,
          position: "static",
          justifyContent: "center",
          backgroundColor: "#262626",
        }}
      >
        <Toolbar
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            margin: "0 -15px",
          }}
        >
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              gap: 20,
            }}
          >
            <img
              src="images/twitch_logo.svg"
              alt="twitch_logo"
              style={{ height: 30, width: 30 }}
            />
            <Typography
              variant="subtitle2"
              sx={{
                fontWeight: "600",
                "&:hover": {
                  color: "#b9a3e3",
                  cursor: "pointer",
                },
              }}
            >
              Following
            </Typography>
            <Typography
              variant="subtitle2"
              sx={{
                fontWeight: "600",
                "&:hover": {
                  color: "#b9a3e3",
                  cursor: "pointer",
                },
              }}
            >
              Browse
            </Typography>
            <Button
              sx={{
                color: "#fff",
                fontSize: "18px",
                maxHeight: 35,
                minWidth: 0,
                "&:hover": {
                  backgroundColor: "hsla(0,0%,100%,0.2)",
                },
              }}
            >
              <i
                className="fi fi-br-menu-dots-vertical"
                style={{ paddingTop: "5px" }}
              />
            </Button>
          </div>
          <div>
            <TextField
              placeholder="Search"
              autoComplete="off"
              variant="outlined"
              size="small"
              sx={{
                backgroundColor: "hsla(0,0%,100%,0.2)",
                borderRadius: "5px",
                width: "340px",
              }}
              inputProps={{
                style: {
                  height: 20,
                  color: "#f1f1f1",
                },
              }}
            />
            <Button
              sx={{
                fontSize: "18px",
                backgroundColor: "hsla(0,0%,100%,0.2)",
                maxHeight: 36,
                minWidth: 0,
                color: "#fff",
                "&:hover": {
                  backgroundColor: "hsla(0,0%,100%,0.3)",
                },
              }}
            >
              <i className="fi fi-br-search" style={{ paddingTop: "5px" }} />
            </Button>
          </div>
          <div
            style={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              gap: 10,
            }}
          >
            <Button
              sx={{
                color: "#fff",
                fontSize: "16px",
                maxHeight: 35,
                minWidth: 0,
                "&:hover": {
                  backgroundColor: "hsla(0,0%,100%,0.2)",
                },
              }}
            >
              <i className="fi fi-br-box" style={{ paddingTop: "5px" }} />
            </Button>
            <Button
              sx={{
                color: "#fff",
                fontSize: "16px",
                maxHeight: 35,
                minWidth: 0,
                "&:hover": {
                  backgroundColor: "hsla(0,0%,100%,0.2)",
                },
              }}
            >
              <i class="fi fi-br-envelope" style={{ paddingTop: "5px" }} />
            </Button>
            <Button
              sx={{
                color: "#fff",
                backgroundColor: "hsla(0,0%,100%,0.2)",
                textTransform: "none",
                "&:hover": {
                  backgroundColor: "hsla(0,0%,100%,0.3)",
                },
                padding: "1px 12px",
                fontSize: "13px",
                fontWeight: "600",
              }}
              startIcon={
                <i
                  className="fi fi-brands-ethereum"
                  style={{ fontSize: "13px", paddingTop: "5px" }}
                />
              }
            >
              Get Bits
            </Button>
            <Avatar
              sx={{
                height: 30,
                width: 30,
                backgroundColor: "#4169E1",
                border: "2px solid #6441a5",
              }}
            />
          </div>
        </Toolbar>
      </AppBar>
    </>
  );
};

export default Header;
