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
  imageSrc: string;
  index: number;
}

const Question: React.FC<QuestionProps> = ({
  question,
  options,
  imageSrc,
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

    <img
      src={imageSrc}
      alt={`Question ${index + 1}`}
      style={{ margin: "5px", marginBottom: "20px" }}
      width={500}
      height={500}
    />
  </div>
);

export default Question;
