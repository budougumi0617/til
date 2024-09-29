import "./index.css";
import { useState } from "react";

function Square(): JSX.Element {
  const [value, setValue] = useState(null as string | null);
  function handleClick() {
    setValue("X");
  }
  return (
    <button className="square" onClick={handleClick}>
      {value}
    </button>
  );
}

export default function Board() {
  return (
    <>
      <div className="board-row">
        <Square />
        <Square />
        <Square />
      </div>
      <div className="board-row">
        <Square />
        <Square />
        <Square />
      </div>
      <div className="board-row">
        <Square />
        <Square />
        <Square />
      </div>
    </>
  );
}
