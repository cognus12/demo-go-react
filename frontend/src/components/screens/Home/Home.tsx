import { FC, useEffect } from "react";
import { Link } from "react-router-dom";
import { Counter } from "../../features/Counter/Counter";
import { HelloFromApi } from "../../features/HelloFromApi/HelloFromApi";
import { Logo } from "../../features/Logo/Logo";
import { Wrapper } from "../../layouts/Wrapper/Wrapper";

const getDynamic = async () => {
  // @ts-ignore
  const m = await import(`./dynamic/dynamic.js`)
  const { dynamic } = m

  return dynamic
}

export const Home: FC = () => {
  const onClickDynamic = () => getDynamic().then(fn => fn())

  return (
    <Wrapper>
      <Logo />
      <p>
        Demo app
        <br /> Frontend: React <br /> Backend: Go <br /> Build tools: Vite
      </p>
      <Counter />
      <button onClick={onClickDynamic}>Dynamic</button>
      <HelloFromApi />
      <Link to="/about">About page</Link>
    </Wrapper>
  );
};
