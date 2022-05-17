import { FC } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { About } from "./components/screens/About/About";

import { Home } from "./components/screens/Home/Home";

const App: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
