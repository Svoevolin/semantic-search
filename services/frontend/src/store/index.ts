import { configureStore } from "@reduxjs/toolkit";
import { api } from "../api";
import { documentListReducer } from "../widgets/DocumentList/slice";

export const store = configureStore({
  reducer: {
    documentList: documentListReducer,

    [api.reducerPath]: api.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(api.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
