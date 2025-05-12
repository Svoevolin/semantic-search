import { createRoot } from "react-dom/client";
import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
import { createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import { Provider } from "react-redux";

import "./index.css";
import App from "./App.tsx";
import { store } from "./store/index.ts";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
  components: {
    // Name of the component
    MuiListItemSecondaryAction: {
      styleOverrides: {
        // Name of the slot
        root: {
          // Some CSS
          top: "5%",
        },
      },
    },
    MuiListItem: {
      styleOverrides: {
        // Name of the slot
        root: {
          // Some CSS
          alignItems: "flex-start",
        },
      },
    },
  },
});

createRoot(document.getElementById("root")!).render(
  <ThemeProvider theme={darkTheme}>
    <Provider store={store}>
      <CssBaseline />
      <App />
    </Provider>
  </ThemeProvider>
);
