import React from "react";

import { createTheme, ThemeProvider, CssBaseline } from "@mui/material";

import Layout from "./components/Layout";

const theme = createTheme({
  typography: {
    fontFamily: "Inter, sans-serif",
  },
});

/*
purple : #6441a5
purple pink : #bf94ff
light purple : #b9a3e3
black : #262626
extraBlack : #202020
grey : "hsla(0,0%,100%,0.2)"
white : #f1f1f1
*/

const App = () => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Layout />
    </ThemeProvider>
  );
};

export default App;
