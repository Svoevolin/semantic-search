import { Box, TextField, Typography, IconButton } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";

import styles from "./styles.module.css";
import { useState } from "react";
import { useFindDocumentsMutation } from "../../api";
import { DocumentList } from "../DocumentList";

const MAX_SEARCH_LENGTH = 384;

export const Search = () => {
  const [text, setText] = useState("");
  const [error, setError] = useState(false);
  const [findDocuments] = useFindDocumentsMutation();

  const onChange = (value: string) => {
    setText(value);
    if (value.length > MAX_SEARCH_LENGTH) {
      setError(true);
      return;
    }

    if (error) {
      setError(false);
    }
  };

  const onSearch = () => {
    findDocuments({ query: text, page: 1, pageSize: 10 });
  };

  return (
    <Box className={styles.root}>
      <Typography variant="h4" align="center">
        Поиcк по базе
      </Typography>
      <Box className={styles.inputWrapper}>
        <TextField
          className={styles.textField}
          label="Поиск"
          multiline
          maxRows={4}
          fullWidth
          onChange={(e) => onChange(e.target.value)}
          value={text}
          error={error}
          helperText={`${text.length} / ${MAX_SEARCH_LENGTH}`}
        />
        <IconButton
          className={styles.searchIcon}
          disabled={error}
          onClick={onSearch}
        >
          <SearchIcon />
        </IconButton>
      </Box>
      <DocumentList text={text} />
    </Box>
  );
};
