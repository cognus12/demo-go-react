import { useState } from "react";

export const Counter = () => {
  const [count, setCount] = useState(0);
  return (
    <p>
      <button type="button" onClick={() => setCount((count) => count + 1)}>
        count is: {count}
      </button>
    </p>
  );
};
