import { FC } from "react";
import { Counter } from "../../features/Counter/Counter";
import { HelloFromApi } from "../../features/HelloFromApi/HelloFromApi";
import { Logo } from "../../features/Logo/Logo";
import { Wrapper } from "../../layouts/Wrapper/Wrapper";

export const Home: FC = () => (
  <Wrapper>
    <Logo />
    <p>Hello Vite + React!</p>
    <Counter />
    <HelloFromApi />
  </Wrapper>
);
