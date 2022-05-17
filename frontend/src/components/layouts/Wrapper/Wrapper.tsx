import { FC, PropsWithChildren, ReactNode } from "react";
import { FCWithOnlyChildren } from "../../../interfaces/FCWithOnlyChildren";

import css from "./wrapper.module.css";

export const Wrapper: FCWithOnlyChildren = ({ children }) => (
  <div className={css.wrapper}>{children}</div>
);
