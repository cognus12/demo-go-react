import { FC } from "react";
import logo from "./assets/logo.svg";
import css from "./logo.module.css";

export const Logo: FC = () => (
  <img src={logo} className={css.logo} alt="logo" />
);
