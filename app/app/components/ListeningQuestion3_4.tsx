import React, { useState, useEffect, use } from "react";
import QuestionsPart2 from "./Question";
import { get } from "http";

interface ImageGalleryProps {
  name_part: string;
  
}

interface QuestionData {
  question: string;
  options: string[];
}
interface AnswerData {
  options: string[];
}

function getQuestions(name: string) {
  const url = `http://localhost:8080/questions/${name}`;
  return fetch(url, {
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      // Extract the 'question_name' property from the API response and create an array of questions
      const questions: QuestionData[] = data.map((item: any) => ({ // Use the appropriate type for 'item'
        question: item.question_text,
      }));
      return questions;
    });
}
function getAnsers(name_part:string) {
  const url = `http://localhost:8080/answers-part/${name_part}`;
  return fetch(url, {
    method: "GET",
  })
    .then((response) => response.json())
    .then((data) => {
      const answers: AnswerData[] = [];
      //const slicedData = data.slice(from*4, to*4);
      if (name_part === "ETS-23-Test1-Part5" || name_part === "ETS-23-Test1-Part6" || name_part === "ETS-23-Test1-Part7" ) {
        for (let i = 0; i < data.length; i+=4) {
          answers.push({
            options: [`${data[i].answer_text}`, `${data[i+1].answer_text}`, `${data[i+2].answer_text}`, `${data[i+3].answer_text}`],
          });
        }
      }
      else {
      for (let i = 0; i < data.length; i+=4) {
        answers.push({
          options: [`A. ${data[i].answer_text}`, `B. ${data[i+1].answer_text}`, `C. ${data[i+2].answer_text}`, `D. ${data[i+3].answer_text}`],
        });
      }}
      return answers;
    });
  }
    


const ListeningQuestionTest345: React.FC<ImageGalleryProps> = ({ name_part }) => {
  
  const [questions, setQuestions] = useState<QuestionData[]>([]);
  const [answers, setAnsers] = useState<AnswerData[]>([]);
  const lastCharacter = name_part.slice(-1);
  let idx: number;
  if (lastCharacter === "3") {
    idx = 31

  }
  else if (lastCharacter === "4") {
    idx = 70
  }
  else if (lastCharacter === "5") {
    idx = 100
  }
  else if (lastCharacter === "6") {
    idx = 130
  }
  else {
    idx = 146
  }
  useEffect(() => {
    getQuestions(name_part)
      .then((data) => {
        setQuestions(data);
      })
      .catch((error) => {
        console.error(error);
      });
    getAnsers(name_part)
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
            index={idx+questionIndex}
          />
        ))
      }
    </div>
  );
};

export default ListeningQuestionTest345;
