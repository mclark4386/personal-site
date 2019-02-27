var ReactDOM = require("react-dom");
import React, { useState } from 'react';
const e = React.createElement;
const root = document.getElementById('root');

function Counter() {
  // Declare a new state variable, which we'll call "count"
  const [count, setCount] = useState(0);
  const [name, setName] = useState("");

  function handleNameChange(e) {
    setName(e.target.value);
  }
  return (
    <div>
      <h3>{name}</h3>
      <input value={name} onChange={handleNameChange}></input>
      <p>You clicked {count} times</p>
      <button onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  );
}

ReactDOM.unstable_createRoot(root).render(e(Counter));

// ReactDOM.render(
//     e(Counter),
//     root
// );
