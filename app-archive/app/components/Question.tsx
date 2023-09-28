// components/Question.tsx
import React from "react";
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
}) => (
  <div>
    <FormControl>
      <FormLabel id={`qus-ans-ansersheet-${index}`}>
        Question {index + 1}: {question}
      </FormLabel>
      <RadioGroup
        aria-labelledby={`qus-ans-ansersheet-${index}`}
        name={`ansersheet-${index}`}
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
);

export default Question;
