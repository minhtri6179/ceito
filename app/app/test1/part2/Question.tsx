// pages/ImageGallery.tsx
import React from "react";
import QuestionPart2 from "../../components/QuestionsPart2"

const ImageGallery: React.FC = () => {
  const imagesFolder = "test1/part1";

  const questions = [
    {
      question: "What is 2 + 2?",
      options: ["A) 3", "B) 4", "C) 5", "D) 6"],
    },
    {
      question: "Which planet is known as the Red Planet?",
      options: ["A) Venus", "B) Mars", "C) Jupiter", "D) Saturn"],
    },
    // Add more questions...
  ];

  return (
    <div>
      {questions.map((questionData, index) => (
        <QuestionPart2
          key={index}
          question={questionData.question}
          options={questionData.options}
          imageSrc={`/${imagesFolder}/${index + 1}.png`} // Assuming image filenames are 1.png, 2.png, etc.
          index={index}
        />
      ))}
    </div>
  );
};

export default ImageGallery;
