// pages/ImageGallery.tsx
import QuestionsPart2 from "./Question"
import React, { useState, useEffect } from "react";

interface ImageGalleryProps {
  index: number; // Receive the index prop
  numofQuestion: number;
}

const ListeningQuestionTest34: React.FC<ImageGalleryProps> = ({index, numofQuestion}) => {
    const [questions, setQuestions] = useState([]); 
    useEffect(() => {
        // Function to fetch questions from the API
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Access-Control-Allow-Origin', '*');
        headers.append('Access-Control-Allow-Credentials', 'true');



        const fetchQuestions = async () => {
          try {
            const response = await fetch("http://localhost:8080/questions", {
          headers: headers,   
          mode: 'cors',
            credentials: 'include',

        });
            console.log(response);
            if (!response.ok) {
              throw new Error("Failed to fetch questions");
            }
    
            const data = await response.json();
            //setQuestions(data); // Update the state with the fetched questions
          } catch (error) {
            console.error("Error fetching questions:", error);
          }
        };
    
        // Call the fetchQuestions function when the component mounts
        fetchQuestions();
      }, []); 
    // question: `Listening`,
    // options: ["A. ", "B. ", "C. ", "D. "],
  return (
    <div>
      {questions.map((questionData, questionIndex) => (
        <QuestionsPart2
          key={index}
        //   question={questionData.question}
        //   options={questionData.options}
            question="Listening"
            options={["A. ", "B. ", "C. ", "D. "]}
          index={index + questionIndex}
        />
      ))}
    </div>
  );
};

export default ListeningQuestionTest34;
