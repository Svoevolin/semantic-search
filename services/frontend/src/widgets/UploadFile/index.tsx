import { Alert, Box, Button, Snackbar } from "@mui/material";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import { styled } from "@mui/material/styles";

import styles from "./styles.module.css";
import { useUploadFileMutation } from "../../api";
import { useState } from "react";

const VisuallyHiddenInput = styled("input")({
  clip: "rect(0 0 0 0)",
  clipPath: "inset(50%)",
  height: 1,
  overflow: "hidden",
  position: "absolute",
  bottom: 0,
  left: 0,
  whiteSpace: "nowrap",
  width: 1,
});

export const UploadFile = () => {
  const [snackbarOpen, setSnackbarOpen] = useState<{
    show: boolean;
    type?: "success" | "error";
    message?: string;
  }>({
    show: false,
  });
  const [uploadFile] = useUploadFileMutation();

  const onUploadFile = (value: File | null | undefined) => {
    if (!value) {
      return;
    }

    const formData = new FormData();

    formData.append("file", value);

    uploadFile(formData).then((value) => {
      if (value.data) {
        setSnackbarOpen({
          show: true,
          type: "success",
          message: "Документ успешно отправлен, обработка займет до 24ч",
        });
      } else {
        setSnackbarOpen({
          show: true,
          type: "error",
          message: "Ошибка при загрузке файла",
        });
      }
    });
  };

  return (
    <Box className={styles.root}>
      <Button
        component="label"
        role={undefined}
        variant="contained"
        tabIndex={-1}
        startIcon={<CloudUploadIcon />}
      >
        Загрузить файл
        <VisuallyHiddenInput
          type="file"
          onChange={(event) => onUploadFile(event.target.files?.item(0))}
          accept=".pdf"
        />
      </Button>
      <Snackbar
        open={snackbarOpen.show}
        autoHideDuration={3000}
        onClose={() => setSnackbarOpen({ ...snackbarOpen, show: false })}
        anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
      >
        <Alert
          onClose={() => setSnackbarOpen({ ...snackbarOpen, show: false })}
          severity={snackbarOpen.type}
          sx={{ width: "100%" }}
        >
          {snackbarOpen.message}
        </Alert>
      </Snackbar>
    </Box>
  );
};
