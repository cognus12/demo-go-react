import { FC, useState } from "react";

export const HelloFromApi: FC = () => {
  const [hello, setHello] = useState("");
  const [loading, setLoading] = useState(false);

  const fetchHello = async () => {
    setLoading(true);

    const response = await fetch(`${location.origin}/api/hello`);

    const data = await response.text();

    setHello(data);

    setLoading(false);
  };

  return (
    <>
      <p>
        <button type="button" onClick={fetchHello} disabled={loading}>
          Get hello from api
        </button>
      </p>
      <p style={{ height: "39px" }}>{hello ? hello : null}</p>
    </>
  );
};
