// Need to use the React-specific entry point to allow generating React hooks
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { DocumentList } from "../widgets/DocumentList/types";
import { documentListActions } from "../widgets/DocumentList/slice";

// Define a service using a base URL and expected endpoints
export const api = createApi({
  reducerPath: "api",
  baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:9999/api/v1/" }),
  endpoints: (build) => ({
    uploadFile: build.mutation<
      { document_id: string; file_name: string; uploaded_at: string },
      FormData
    >({
      query: (data) => ({
        url: "documents/upload",
        method: "POST",
        body: data,
      }),
    }),
    findDocuments: build.mutation<
      DocumentList,
      { page: number; pageSize: number; query: string }
    >({
      query: ({ page, pageSize, query }) => ({
        url: `documents?_page=${page}&_pagesize=${pageSize}`,
        method: "POST",
        body: { query },
      }),
      async onQueryStarted(_, { dispatch, queryFulfilled }) {
        try {
          // Получаем ответ от бекенда – он должен вернуть { task_id: string }
          const { data, meta } = await queryFulfilled;
          dispatch(
            documentListActions.setHasMore(
              Boolean(meta?.response?.headers.get("X-Has-More"))
            )
          );
          dispatch(documentListActions.setDocumentList(data));
        } catch (e) {
          console.error("Ошибка загрузки файлов", e);
        }
      },
    }),
  }),
});

// Export hooks for usage in function components, which are
// auto-generated based on the defined endpoints
export const { useUploadFileMutation, useFindDocumentsMutation } = api;
