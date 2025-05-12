import {
  Avatar,
  Box,
  IconButton,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  TablePagination,
  Typography,
} from "@mui/material";
import CloudDownloadIcon from "@mui/icons-material/CloudDownload";
import DescriptionIcon from "@mui/icons-material/Description";
import { useAppSelector } from "../../shared/hooks";
import { selectDocumentList, selectHasMore } from "./selectors";
import styles from "./styles.module.css";
import { useState } from "react";
import { useFindDocumentsMutation } from "../../api";

type DocumentListProps = {
  text: string;
};

const DEFAULT_ROWS_PER_PAGE = 10;

export const DocumentList = ({ text }: DocumentListProps) => {
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(DEFAULT_ROWS_PER_PAGE);
  const documentList = useAppSelector(selectDocumentList);
  const hasMore = useAppSelector(selectHasMore);
  const [findDocuments] = useFindDocumentsMutation();

  const handleChangePage = (
    event: React.MouseEvent<HTMLButtonElement> | null,
    newPage: number
  ) => {
    setPage(newPage);
    findDocuments({ page: newPage + 1, pageSize: rowsPerPage, query: text });
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const rowsPerPage = parseInt(event.target.value, 10);
    setRowsPerPage(rowsPerPage);
    setPage(0);
    findDocuments({ page: 1, pageSize: rowsPerPage, query: text });
  };

  const getContent = () => {
    if (!documentList) {
      return null;
    }
    if (!documentList.length) {
      return (
        <Typography variant="h3" align="center">
          Подходящие файлы не найдены
        </Typography>
      );
    }

    return (
      <>
        <List>
          {[
            ...documentList,
            ...documentList,
            ...documentList,
            ...documentList,
            ...documentList,
            ...documentList,
          ].map((document) => (
            <ListItem
              key={document.document_id + Math.random()}
              secondaryAction={
                <IconButton edge="end" aria-label="download">
                  <CloudDownloadIcon />
                </IconButton>
              }
              className={styles.listItem}
            >
              <ListItemAvatar>
                <Avatar>
                  <DescriptionIcon />
                </Avatar>
              </ListItemAvatar>
              <ListItemText
                primary={document.file_name}
                secondary={
                  <>
                    <p>{document.snippet}</p>
                    <p>
                      {new Date(document.uploaded_at).toLocaleDateString(
                        "ru-RU"
                      )}
                    </p>
                  </>
                }
              />
            </ListItem>
          ))}
        </List>
        <TablePagination
          component="div"
          count={-1}
          page={page}
          onPageChange={handleChangePage}
          rowsPerPage={rowsPerPage}
          onRowsPerPageChange={handleChangeRowsPerPage}
          labelDisplayedRows={({ from, to }) => `${from} - ${to}`}
          labelRowsPerPage="Строк на страницу"
          slotProps={{
            actions: {
              nextButton: {
                disabled: !hasMore,
              },
            },
          }}
        />
      </>
    );
  };

  return <Box className={styles.root}>{getContent()}</Box>;
};
