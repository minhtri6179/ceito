import React, { useState, useEffect } from "react";
import QuestionsPart2 from "./Question";
import { get } from "http";

interface ImageGalleryProps {
  index: number;
  numofQuestion: number;
}

interface QuestionData {
  question: string;
  options: string[];
}
interface AnswerData {
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
      }));
      return questions;
    });
}
function getAnsers(from : number, to : number) {
  return fetch("http://localhost:8080/answers", {
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      const answers: AnswerData[] = [];
      const slicedData = data.slice(from*4, to*4);
      console.log(slicedData);

      for (let i = 0; i < slicedData.length; i+=4) {
        answers.push({
          options: [`A. ${slicedData[i].answer_text}`, `B. ${slicedData[i+1].answer_text}`, `C. ${slicedData[i+2].answer_text}`, `D. ${slicedData[i+3].answer_text}`],
        });
      }
      return answers;
    });
  }
    


const ListeningQuestionTest34: React.FC<ImageGalleryProps> = ({ index, numofQuestion }) => {
  
  const [questions, setQuestions] = useState<QuestionData[]>([]);
  const [answers, setAnsers] = useState<AnswerData[]>([]);
  useEffect(() => {
    getQuestions(index, index+numofQuestion)
      .then((data) => {
        setQuestions(data);
      })
      .catch((error) => {
        console.error(error);
      });
    getAnsers(index, index+numofQuestion)
      .then((data) => {
        setAnsers(data);
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
            options={answers[questionIndex]?.options || []}
            index={index + questionIndex}
          />
        ))
      }
    </div>
  );
};

export default ListeningQuestionTest34;
