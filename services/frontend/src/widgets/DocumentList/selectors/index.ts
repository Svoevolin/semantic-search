import { RootState } from "../../../store";

export const selectDocumentList = (state: RootState) =>
  state.documentList.documentList;

export const selectHasMore = (state: RootState) => state.documentList.hasMore;
