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
import Grid from '@mui/material/Grid';

interface QuestionProps {
  question: string;
  options: string[];
  index: number;
  imageSrc?: string;
  img_size?: string;

}

const Question: React.FC<QuestionProps> = ({
  question,
  options,
  index,
  imageSrc,
  img_size,
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
        <Grid container rowSpacing={1} columnSpacing={{ xs: 1, sm: 1, md: 1 }}>
        <Grid item xs={4}>
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
        </Grid>
        <Grid item xs={6}>
        {imageSrc && (
          <div style={{marginLeft: '20px', maxWidth: '250px'}}>
            <Image
              src={imageSrc}
              alt={`Question ${index + 1} Image`}
              width={200*Number(img_size)}
              height={200*Number(img_size)}
            />
          </div>
        )}
        </Grid>
        </Grid>
        </div>
    </div>
  )};
  

export default Question;
