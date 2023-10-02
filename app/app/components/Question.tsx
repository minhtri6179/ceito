// components/Question.tsx
import React, { useState, useEffect } from "react";
import {
  FormControl,
  FormControlLabel,
  FormLabel,
  RadioGroup,
  Radio,
} from "@mui/material";

interface QuestionProps {
  question: string;
  options: string[];
  index: number;
}

const Question: React.FC<QuestionProps> = ({
  question,
  options,
  index,
}) => { 
  const localStorageKey = `selectedAnswers-${index}`;
  const [selectedAnswer, setSelectedAnswer] = useState(() => {
    const storedAnswer = localStorage.getItem(localStorageKey);
    return storedAnswer || "";
  });

  useEffect(() => {
    localStorage.setItem(localStorageKey, selectedAnswer);
  }, [selectedAnswer, localStorageKey]);

  const handleAnswerChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSelectedAnswer(event.target.value);
  };
return (
  
  <div>
    <FormControl>
      <FormLabel id={`qus-ans-ansersheet-${index}`}>
        Question {index + 1}: {question}
      </FormLabel>
      <RadioGroup
          aria-labelledby={`qus-ans-ansersheet-${index}`}
          name={`ansersheet-${index}`}
          value={selectedAnswer}
          onChange={handleAnswerChange}
        >
        {options.map((option, optionIndex) => (
          <FormControlLabel
            key={optionIndex}
            value={option.charAt(0)}
            control={<Radio />}
            label={option}
          />
        ))}
      </RadioGroup>
    </FormControl>
  </div>
)};

export default Question;
