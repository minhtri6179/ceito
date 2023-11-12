// components/Question.tsx
import React, { useState, useEffect } from "react";
import {
  FormControl,
  FormControlLabel,
  FormLabel,
  RadioGroup,
  Radio,
} from "@mui/material";
import Image from 'next/image';

interface QuestionProps {
  question: string;
  options: string[];
  index: number;
  imageSrc?: string;

}

const Question: React.FC<QuestionProps> = ({
  question,
  options,
  index,
  imageSrc,
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
      <div style={{ display: 'flex', alignItems: 'center' }}>
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
        {imageSrc && (
          <div style={{marginLeft: '10px', maxWidth: '100px', maxHeight: '100px' }}>
            <Image
              src={imageSrc}
              alt={`Question ${index + 1} Image`}
              width={100}
              height={100}
            />
          </div>
        )}</div>
    </div>
  )};
  

export default Question;
