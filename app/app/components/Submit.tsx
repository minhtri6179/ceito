import React, { useState } from "react";
import submit from "../../util/Submit";

const SubmitButton = () => {
  const [result, setResult] = useState(""); // Initialize result with an empty string

  const handleClick = () => {
    submit()
      .then((response) => { // Capture the response from submit()
        setResult(response); // Set the result with the response data
      })
      .catch((error) => {
        console.error("Error submitting data:", error);
      });
  };

  return (
    <div>
      <button onClick={handleClick}><strong>Nộp bài thôi </strong></button>
      <div>
        <strong>Kết quả:</strong> {result}
      </div>
    </div>
  );
};

export default SubmitButton;
