import { FC } from "react";
import { Link } from "react-router-dom";
import { Counter } from "../../features/Counter/Counter";
import { HelloFromApi } from "../../features/HelloFromApi/HelloFromApi";
import { Logo } from "../../features/Logo/Logo";
import { Wrapper } from "../../layouts/Wrapper/Wrapper";

export const Home: FC = () => (
  <Wrapper>
    <Logo />
    <p>Demo app<br /> Frontend: React <br /> Backend: Go <br /> Build tools: Vite, Deno</p>
    <Counter />
    <HelloFromApi />
    <Link to="/about">About page</Link>
  </Wrapper>
);
