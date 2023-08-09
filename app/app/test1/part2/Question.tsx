// pages/ImageGallery.tsx
import React from "react";
import QuestionsPart2 from "../../components/QuestionsPart2"

const ImageGallery: React.FC = () => {

  const questions = Array.from({ length: 25 }, (_, index) => ({
    question: `Listening`,
    options: ["A. ", "B. ", "C. "],
  }));
  

  return (
    <div>
      {questions.map((questionData, index) => (
        <QuestionsPart2
          key={index+6}
          question={questionData.question}
          options={questionData.options}
          index={index+6}
        />
      ))}
    </div>
  );
};

export default ImageGallery;
