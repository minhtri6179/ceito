import React, { useState, useEffect, use } from "react";
import QuestionsPart2 from "./Question";
import getImage from "../../util/Imagehander";


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
const ids = [ '62', '65','68', '131', '135', '139', '143', '147', '149', '151', '154', '156', '158', '162', '166', '169', '172']
const part7 = [2, 2, 3, 2, 2, 4, 4, 3, 3, 4, 5, 5, 5, 5, 5] 

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
function mappingidxtofilename(idx: number, test_part: string) {
  if (idx >= 146) {
    return `/test1/part7/${idx}.png`
  }
  else if (idx >= 130) {
    return `/test1/part6/${idx}.png`
  }
  else if (idx >= 70) {
    return `/test1/part4/${idx}.png`
  }
  else if (idx >= 31) {
    return `/test1/part3/${idx}.png`
  }
}
function mappingidxtossize(idx: number) {
  if (idx >= 146) {
    var cur = 146
    for (let i = 0; i < part7.length; i++) {
      cur += part7[i]
      if (cur >= idx) {
        return part7[i]
      }
    }
  }
  else if (idx >= 130) {
    return 4
  }
  else if (idx >= 31) {
    return 3
  }
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
  const specialQuestionIds = [1,2,3]
  const groupedQuestions = specialQuestionIds.map((specialQuestionId) =>
  questions.filter((questionData, questionIndex) => questionIndex === specialQuestionId)
  );


  return (
    <div>
      {questions.map((questionData, questionIndex) => (
          <QuestionsPart2
            key={questionIndex}
            question={questionData.question}
            options={answers[questionIndex]?.options || []}
            index={idx+questionIndex}
            imageSrc={ids.includes(`${idx+questionIndex+1}`) ? `${mappingidxtofilename(idx+questionIndex+1, "hihi")}` : ""}
            img_size={ids.includes(`${idx+questionIndex+1}`) ? `${mappingidxtossize(idx+questionIndex+1)}` : "1"}
          />
        ))
      }
    </div>
  );
};

export default ListeningQuestionTest345;
