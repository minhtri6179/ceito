// pages/ImageGallery.tsx
import React, { useState, useEffect } from "react";
import QuestionsPart2 from "./Question"
interface ImageGalleryProps {
  index: number; // Receive the index prop
  numofQuestion: number;
}
const ListeningQuestion: React.FC<ImageGalleryProps> = ({index, numofQuestion}) => {

  const questions = Array.from({ length: numofQuestion }, (_, questionIndex) => ({
    question: `Listening`,
    options: ["A. ", "B. ", "C. "],
  }));
  

  return (
    <div>
      {questions.map((questionData, questionIndex) => (
        <QuestionsPart2
          key={index}
          question={questionData.question}
          options={questionData.options}
          index={index + questionIndex}
        />
      ))}
    </div>
  );
};

export default ListeningQuestion;
