// pages/ImageGallery.tsx
import React from "react";
import QuestionsPart2 from "../../components/QuestionsPart2"

const ImageGallery: React.FC = () => {

  const questions = Array.from({ length: 30 }, (_, index) => ({
    question: `Listening`,
    options: ["A. ", "B. ", "C. ", "D."],
  }));
  

  return (
    <div>
      {questions.map((questionData, index) => (
        <QuestionsPart2
          key={index+70}
          question={questionData.question}
          options={questionData.options}
          index={index+70}
        />
      ))}
    </div>
  );
};

export default ImageGallery;
