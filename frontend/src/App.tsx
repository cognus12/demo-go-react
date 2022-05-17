import { FC } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { About } from "./components/screens/About/About";

import { Home } from "./components/screens/Home/Home";
import { NotFound } from "./components/screens/NotFound/NotFound";

const App: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
