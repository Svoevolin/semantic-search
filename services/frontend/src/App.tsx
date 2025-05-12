import { Container } from "@mui/material";
import styles from "./App.module.css";
import { Search } from "./widgets/Search";
import { UploadFile } from "./widgets/UploadFile";

function App() {
  return (
    <Container className={styles.root}>
      <Search />
      <UploadFile />
    </Container>
  );
}

export default App;
