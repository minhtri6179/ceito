import React, { useState, useEffect } from "react";
import QuestionsPart2 from "./Question";

interface ImageGalleryProps {
  index: number;
  numofQuestion: number;
}

interface QuestionData {
  question: string;
  options: string[];
}

function getQuestions(from : number, to : number) {
  return fetch("http://localhost:8080/questions", {
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      // Extract the 'question_name' property from the API response and create an array of questions
      const questions: QuestionData[] = data.slice(from, to).map((item: any) => ({ // Use the appropriate type for 'item'
        question: item.question_text,
        options: ["A. ", "B. ", "C. ", "D. "],
      }));
      return questions;
    });
}


const ListeningQuestionTest34: React.FC<ImageGalleryProps> = ({ index, numofQuestion }) => {
  
  const [questions, setQuestions] = useState<QuestionData[]>([]);

  useEffect(() => {
    getQuestions(index, index+numofQuestion)
      .then((data) => {
        setQuestions(data);
      })
      .catch((error) => {
        console.error(error);
      });
  }, []);


  return (
    <div>
      {questions.map((questionData, questionIndex) => (
          <QuestionsPart2
            key={questionIndex}
            question={questionData.question}
            options={questionData.options}
            index={index + questionIndex}
          />
        ))
      }
    </div>
  );
};

export default ListeningQuestionTest34;
