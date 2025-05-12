import type { PayloadAction } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import { DocumentList } from "../types";

export interface DocumentListState {
  documentList: DocumentList | null;
  hasMore: boolean | null;
}

const initialState: DocumentListState = {
  documentList: null,
  hasMore: null,
};

export const documentListSlice = createSlice({
  name: "documentList",
  initialState,
  reducers: {
    setDocumentList: (state, action: PayloadAction<DocumentList>) => {
      state.documentList = action.payload;
    },
    setHasMore: (state, action: PayloadAction<boolean>) => {
      state.hasMore = action.payload;
    },
  },
});

export const { actions: documentListActions, reducer: documentListReducer } =
  documentListSlice;
